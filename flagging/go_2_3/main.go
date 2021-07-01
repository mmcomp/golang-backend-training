package main

import "fmt"

var name string
var shhh bool

func main() {
	Init()
	if !shhh {
		fmt.Printf("Hello %s\n", name)
	}
}
