package main

import (
	"fmt"
	"strings"

	"github.com/daltonscharff/city-stats/sources"
)

func handler() {
	location := "Austin, TX"
	loc := strings.ToLower(location)
	fmt.Println(location)

	numbeo := sources.Numbeo{}

	fmt.Println(numbeo.Find((loc)))
}

func main() {
	handler()
}
