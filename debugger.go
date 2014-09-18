package godebug

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type Debugger struct {
	started     bool
	breakpoints map[string]breakpoint
	watches     map[string]interface{}
	dbgChan     chan struct{}
	c           string
	fileName    string
	lines       string
	cbp         string
	stats       []stat
}

func NewDebugger(port int) *Debugger {
	d := &Debugger{
		started:     false,
		dbgChan:     make(chan struct{}, 0),
		breakpoints: make(map[string]breakpoint),
		watches:     make(map[string]interface{}),
	}

	d.listen(port)

	return d
}

// Stops execution of the code until asked to continue.
func (d *Debugger) Break() {
	if d.started {

		// snapshot the runtime stats
		d.stats = getRuntimeStats()

		var stack [4096]byte
		runtime.Stack(stack[:], false)
		s := fmt.Sprintf("%s\n", stack[:])

		file, line := getFileDetails(s)

		d.fileName = file

		id := generateBreakpointId(file, line)
		b := breakpoint{
			enabled: true,
		}

		// add the breakpoint if it doesn't already exist
		if bc, exists := d.breakpoints[id]; exists {
			if !bc.enabled {
				// just return if breakpoint is disabled
				return
			}
		} else {
			d.breakpoints[id] = b
		}

		d.takeFileSnapshot(file, line)

		select {
		case <-d.dbgChan:
			break
		}

	}
}

// Adds a variable to be watched. The key identifies the variable in the UI
func (d *Debugger) Watch(key string, i interface{}) error {
	if d.started {
		if _, exists := d.watches[key]; !exists {
			d.watches[key] = i
		} else {
			return errors.New("Watch key already exists.")
		}
	}
	return nil
}

func (d *Debugger) listen(port int) {
	go func() {
		http.HandleFunc("/", d.index)
		http.HandleFunc("/debugger.css", d.styles)
		http.HandleFunc("/debugger.js", d.scripts)
		http.HandleFunc("/code", d.code)
		http.HandleFunc("/continue", d.cont)
		http.HandleFunc("/watches", d.watch)
		http.HandleFunc("/toggle", d.toggle)

		http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	}()

	d.started = true
}

func (d *Debugger) index(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprint(w, debugger_markup)
}

func (d *Debugger) styles(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/css")
	fmt.Fprint(w, debugger_css)
}

func (d *Debugger) scripts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/javascript")
	fmt.Fprint(w, debugger_js)
}

func (d *Debugger) code(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	enc := json.NewEncoder(w)

	enc.Encode(struct {
		Code, FileName, Lines, Breakpoint string
		Stats                             []stat
	}{
		d.c,
		d.fileName,
		d.lines,
		d.cbp,
		d.stats,
	})
}

func (d *Debugger) cont(w http.ResponseWriter, r *http.Request) {
	d.dbgChan <- struct{}{}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (d *Debugger) watch(w http.ResponseWriter, r *http.Request) {
	i := 1
	w.Header().Add("Content-Type", "application/json")
	enc := json.NewEncoder(w)

	e := make([]watch, 0)

	for k, v := range d.watches {
		e = append(e, watch{Name: k, Value: v, Index: i})
		i++
	}

	enc.Encode(e)
}

func (d *Debugger) toggle(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if b, exists := d.breakpoints[id]; exists {
		b.enabled = !b.enabled
		d.breakpoints[id] = b
	}
}

func (d *Debugger) takeFileSnapshot(file string, line int) {
	var b bytes.Buffer
	var l bytes.Buffer

	b.WriteString("<pre>")

	f, err := os.Open(file)
	if err != nil {
		b.WriteString(err.Error())
		return
	}

	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	c := 1
	for s.Scan() {
		if c == line {
			var classes string = "highlight"
			id := generateBreakpointId(file, c)
			bp, exists := d.breakpoints[id]

			if exists && bp.enabled {
				d.cbp = id
			}

			if exists && !bp.enabled {
				classes += " disabled"
			}
			b.WriteString("<span class=\"" + classes + "\">" + s.Text() + "</span><br/>")
			l.WriteString(strconv.Itoa(c) + ":<br/>")
		} else if c > line-10 {
			b.WriteString(strings.Replace(strings.Replace(s.Text(), "<", "&lt;", -1), ">", "&gt;", -1) + "<br/>")
			l.WriteString(strconv.Itoa(c) + ":<br/>")
		} else if c > line+11 {
			break
		}
		c++
	}
	b.WriteString("</pre>")
	d.c = strings.Replace(b.String(), "\t", "&nbsp;&nbsp;&nbsp;&nbsp;", -1)
	d.lines = l.String()
}
