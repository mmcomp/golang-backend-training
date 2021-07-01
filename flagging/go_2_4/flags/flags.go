package flags

import (
	"flag"
)

var Name string
var Shhh bool

func Handle() {
	flag.StringVar(&Name, "name", "NoName", "It is your name!")
	flag.BoolVar(&Shhh, "shhh", false, "It is used to SHHH the output!")
	flag.Parse()
}
