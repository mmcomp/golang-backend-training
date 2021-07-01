package main

import (
	"fmt"
	f "go_2_4/flags"
)



func main() {
	f.Handle()
	if !f.Shhh {
		fmt.Printf("Hello %s\n", f.Name)
	}
}
