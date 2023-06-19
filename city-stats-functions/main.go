package main

import (
	"fmt"

	"github.com/daltonscharff/city-stats/internal/sources"
)

func handler() {
	location := "Austin, TX"
	fmt.Println(location)

	numbeo := sources.Numbeo{}

	fmt.Println(numbeo.Find((location)))
}

func main() {
	handler()
}
