package main

import (
	"fmt"
)

func main() {
	const (
		x string = "Hello world"
	)
	var (
		a = 5
		b = 10
		c = 15
	)
	fmt.Println(x)

	fmt.Println(a + b + c)
}
