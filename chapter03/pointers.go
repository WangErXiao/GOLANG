package main

import (
	"fmt"
)

func zero(x int) {
	x = 0
}
func one(xPtr *int) {
	*xPtr = 1
}

func square(x *float64) {
	*x = *x * *x
}
func main() {
	x := 5
	zero(x)
	fmt.Println(x)
	one(&x)
	fmt.Println(x)

	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr)
	z := 2.0
	square(&z)
	fmt.Println(z)
}
