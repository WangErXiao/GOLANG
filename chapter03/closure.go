package main

import (
	"fmt"
)

func main() {
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 1))

	xx()
	fmt.Println(factorial(5))
}
func xx() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())

}
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
