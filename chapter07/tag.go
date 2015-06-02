package main

import (
	"fmt"
	"reflect"
)

type TagType struct {
	field1 bool   "An import answer"
	field2 string "The name of thing"
	field3 int    "How much there are"
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixFiled := ttType.Field(ix)
	fmt.Println(ixFiled.Tag)
}

func main() {
	tt := TagType{true, "Barak Obama", 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}
