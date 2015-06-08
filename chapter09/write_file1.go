package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	inputFile := "products.txt"
	outputFile := "products_copy.txt"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return
	}
	fmt.Printf("%s\n", string(buf))
	ioutil.WriteFile(outputFile, buf, 0x644)
	if err != nil {
		panic(err.Error())
	}

}
