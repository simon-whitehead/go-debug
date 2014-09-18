package main

import "github.com/simon-whitehead/go-debug"

func main() {
	debugger := godebug.NewDebugger(8080)
	i := 0

	debugger.Watch("i", &i)

	debugger.Break()

	i = 50

	debugger.Break()
}
