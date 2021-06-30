package flags

import (
	"flag"
	"fmt"
)

var name string
var shhh bool

func init() {
	flag.StringVar(&name, "name", "NoName", "It is your name!")
	flag.BoolVar(&shhh, "shhh", false, "It is used to SHHH the output!")
	flag.Parse()
}

func Action() {
	if !shhh {
		fmt.Printf("Hello %s\n", name)
	}
}
