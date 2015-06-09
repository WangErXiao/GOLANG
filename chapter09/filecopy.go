package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	CopyFile("filecopy.txt", "filecopy.go")
	fmt.Println("Copy done!")
}
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}
