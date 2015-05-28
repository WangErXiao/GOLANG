package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x float64
	y float64
	r float64
}
type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("hi my name is ", p.Name)
}

func circleArea(c Circle) float64 {
	return math.Pi * math.Pow(c.r, 2)
}

func main() {
	c := Circle{0, 0, 5}
	fmt.Println(c.r)
	fmt.Println(c.x)

	fmt.Println(math.MaxFloat64)
	fmt.Println(circleArea(c))
	p := Person{"yao"}
	p.Talk()
}
