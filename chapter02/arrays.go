package main

import (
	"fmt"
)

func main() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	var y [5]float64
	y[0] = 98
	y[1] = 93
	y[2] = 77

	var total float64 = 0
	for i := 0; i < 3; i++ {
		total += y[i]
	}
	fmt.Println(total / float64(len(y)))

	total = 0
	for _, value := range x {
		total += float64(value)
	}
	fmt.Println(total / float64(len(y)))

	z := [5]float64{98, 93, 77, 82, 83}
	fmt.Println(z)

	m := make([]float64, 5)
	fmt.Println(m)

	n := make([]float64, 5, 10)
	fmt.Println(n)
}
