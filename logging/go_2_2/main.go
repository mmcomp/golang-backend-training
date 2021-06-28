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

func (receiver Logger) Logf(format string, a ...interface{}) {
	if receiver.OutputText {
		fmt.Fprintf(receiver.Output, format, a...)
	}
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
