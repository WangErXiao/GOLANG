package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	file, err := os.Open("file.go")
	if err != nil {
		return
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return
	}

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)
	fmt.Println("------------")
	readFile()
}

func readFile() {
	bs, err := ioutil.ReadFile("file.go")
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)
}
