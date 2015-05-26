package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter a number:")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)
	fmt.Println(`1
		2
		3
		4
		5`)
}
