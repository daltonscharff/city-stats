package main

import (
	"fmt"
	"strings"
)

func handler() {
	location := "Austin, TX"
	loc := strings.ToLower(location)
	fmt.Println(location)

	numbeo := Numbeo{}

	fmt.Println(numbeo.Find((loc)))
}

func main() {
	handler()
}
