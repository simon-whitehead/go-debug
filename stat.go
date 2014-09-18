package godebug

import (
	"fmt"
	"runtime"
	"time"
)

// Ripped straight from Effective Go
type ByteSize float64

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2fYB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%.2fZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%.2fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}

type stat struct {
	Name  string
	Value string
}

func getRuntimeStats() []stat {
	abs := float64(time.Millisecond)

	m := runtime.MemStats{}

	runtime.ReadMemStats(&m)

	// I don't know if these are 100% accurate in terms of their meanings
	return []stat{
		stat{Name: "# of goroutines", Value: fmt.Sprintf("%d", runtime.NumGoroutine())},
		stat{Name: "# of allocated objects", Value: fmt.Sprintf("%d", m.HeapObjects)},
		stat{Name: "# of garbage collections", Value: fmt.Sprintf("%d", m.NumGC)},
		stat{Name: "Total GC pause time", Value: fmt.Sprintf("%.2fms", float64(m.PauseTotalNs)/abs)},
		stat{Name: "OS allocated memory", Value: fmt.Sprintf("%s", ByteSize(m.Sys))},
		stat{Name: "Runtime allocated memory", Value: fmt.Sprintf("%s", ByteSize(m.TotalAlloc))},
		stat{Name: "Runtime used memory", Value: fmt.Sprintf("%s", ByteSize(m.Alloc))},
		stat{Name: "Stack size", Value: fmt.Sprintf("%s", ByteSize(m.StackInuse))},
		stat{Name: "Available stack size", Value: fmt.Sprintf("%s", ByteSize(m.StackSys))},
	}
}
