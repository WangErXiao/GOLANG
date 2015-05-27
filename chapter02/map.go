package main

import (
	"fmt"
)

func main() {
	x := make(map[string]int)
	x["key"] = 10
	x["yao"] = 11
	x["xxx"] = 12
	fmt.Println(x)
	fmt.Println(x["key"])
	age, ok := x["yao"]
	fmt.Println(age, ok)

	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Boron",
		"C":  "Carbo",
	}
	for k, _ := range elements {
		fmt.Println(elements[k])
	}

}
