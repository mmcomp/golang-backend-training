package main

import (
	"fmt"

	"github.com/mmcomp/go-money"
)

func main() {
	money, err := money.ParseCAD("$1.23")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("money:", money)
}
