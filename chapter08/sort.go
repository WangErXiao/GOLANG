package main

import (
	"fmt"
)

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Sort(data Sorter) {
	for pass := 1; pass < data.Len(); pass++ {
		for i := 0; i < data.Len()-pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}
func IsSorted(data Sorter) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}

type IntArray []int

func (p IntArray) Len() int {
	return len(p)
}
func (p IntArray) Less(i, j int) bool {
	return p[i] < p[j]
}
func (p IntArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]

}

func main() {
	data := []int{12, -11, 1, 2, -3, 0}
	a := IntArray(data)
	Sort(a)
	if !IsSorted(a) {
		panic("fail")
	}
	fmt.Printf("The sorted array is:%v\n", a)
}
