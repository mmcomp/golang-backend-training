package flags

import (
	"flag"
	"fmt"
)

func Handle() {
	var name string
	var shhh bool
	flag.StringVar(&name, "name", "NoName", "It is your name!")
	flag.BoolVar(&shhh, "shhh", false, "It is used to SHHH the output!")
	flag.Parse()
	if !shhh {
		fmt.Printf("Hello %s\n", name)
	}
}
