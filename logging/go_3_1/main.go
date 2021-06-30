package main

import (
	"fmt"
	"io"
	"os"
)

type Logger struct {
	Output io.Writer
}

func (receiver Logger) Log(a ...interface{}) {
	if receiver.Output == nil {
		return
	}
	fmt.Fprintln(receiver.Output, a...)
}

func (receiver Logger) Begin(a ...interface{}) {
	receiver.Log("BEGIN")
}

func (receiver Logger) End(a ...interface{}) {
	receiver.Log("END")
}

func main() {
	log := Logger{
		Output: os.Stdout,
	}

	log.Begin()

	log.End()
}
