package main

import (
	"fmt"
	"io"
	"os"
)

type Logger struct {
	OutputText bool
	Output     io.Writer
}

func (receiver Logger) Log(a ...interface{}) {
	if receiver.OutputText {
		fmt.Fprintln(receiver.Output, a...)
	}
}

func (receiver Logger) Logf(format string, a ...interface{}) {
	if receiver.OutputText {
		fmt.Fprintf(receiver.Output, format, a...)
	}
}

func (receiver Logger) Begin(a ...interface{}) {
	receiver.Log("BEGIN")
}

func (receiver Logger) End(a ...interface{}) {
	receiver.Log("END")
}

func main() {
	log := Logger{
		OutputText: true,
		Output:     os.Stdout,
	}

	name := "Mehrdad"
	fmt.Printf("Hello %s!\n", name)
	log.Logf("The name was %q\n", name)

}
