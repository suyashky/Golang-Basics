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

	// dump all data to a file
	d1 := []byte("Hehe ayee boy!")
	err := ioutil.WriteFile("data.txt", d1, 0644)
	check(err)

	// checking by reading
	data, err := ioutil.ReadFile("data.txt")
	check(err)
	fmt.Println(string(data))

	// for custom writes

	// open for writing(different than open for read)
	file1, err := os.Create("data2.txt")

	// we can also have a slce of byte
	d2 := []byte("basic data")
	_, err = file1.Write(d2)
	check(err)

	file1.Close()

	//check if printed
	file2, err := os.Open("data2.txt")
	check(err)

	b2 := make([]byte, 6)
	_, err = file2.Read(b2)
	check(err)
	fmt.Println(string(b2))

	file2.Close()
}
