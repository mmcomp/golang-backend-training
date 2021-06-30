package main

import (
	"flag"
	"io/ioutil"
	"os"

	l "github.com/mmcomp/go-log"
)

func main() {
	var verboseMode bool
	flag.BoolVar(&verboseMode, "v", false, "Verbose Mode")
	flag.Parse()

	var log l.Logger
	if verboseMode {
		log = l.Logger{
			Output: os.Stdout,
		}
	} else {
		log = l.Logger{
			Output: ioutil.Discard,
		}
	}

	log.Begin()

	log.End()
}
