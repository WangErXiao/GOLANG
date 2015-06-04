package main

import (
	"fmt"
)

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func main() {
	sq1 := new(Square)
	sq1.side = 5

	areaIntf := sq1
	fmt.Println("The square has area:%f\n", areaIntf.Area())

	areaInf1 := Shaper(sq1)
	fmt.Println("The square has area:%f\n", areaInf1.Area())

}
