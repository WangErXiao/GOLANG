package main

import (
	"fmt"
)

var y string = "Hello global"

func main() {
	var x string = "Hello World"
	fmt.Println(x)
	fmt.Println(y)
}
