package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "NoName", "It is your name!")
	flag.Parse()
	fmt.Printf("Hello %s\n", name)
}
