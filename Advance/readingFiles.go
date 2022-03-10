package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// use readfile to get whole content of file in byte array
	// no need to open/close file
	dat, err := ioutil.ReadFile("sample.txt")
	check(err)
	fmt.Println(string(dat)) // printed that []byte

	// for more control, or read specific parts only
	// opening a file
	file1, err := os.Open("sample.txt")
	check(err)

	// reading only 4 bytes from start
	b1 := make([]byte, 4)
	n1, err := file1.Read(b1)
	check(err)
	fmt.Println(n1, string(b1))

	// seek can used to move file pointer
	// say we want to read from 6rd char, use seek
	file1.Seek(6, 0) // 5 char skipped, pointer moced to 6th
	b2 := make([]byte, 4)
	_, err = file1.Read(b2)
	check(err)
	fmt.Println(string(b2))

	file1.Close() // have to do it manually
}
