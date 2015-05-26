package main

import (
	"fmt"
)

func main() {
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("21.2-21.1=", 21.2-21.1)

	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!false)

}
