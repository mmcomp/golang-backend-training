package main

import (
	"flag"
)

func Init() {
	flag.StringVar(&name, "name", "NoName", "It is your name!")
	flag.BoolVar(&shhh, "shhh", false, "It is used to SHHH the output!")
	flag.Parse()
}
