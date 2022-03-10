// Topics covered
// Basic of GO program
// Variables
// Loop
// If else
// Switch
// Arrays
// Slices
// Map
// Range

package main // compulsary

import "fmt" // for in/op contains fn like print

// for importing multiple packages
/*
import {
	"fmt"
	"math"
}
*/

func main() {
	fmt.Print("Hello, World!")

	// Variables ###############
	var a, b string = "john", "trello"
	fmt.Println(a, b)

	// uninitialized vars have zero value of that type
	var c int
	fmt.Println(c)

	name := "Boris"
	fmt.Println(name)

	// Loop ###############

	// curly braces are compulsary
	n := 5
	for i := 0; i < n; i++ {
		fmt.Print("Here ")
	}

	fmt.Printf("\n")

	n = 3
	for n > 0 {
		fmt.Println(n)
		n--
	}

	// something similar to while(1)
	for {
		fmt.Println("heyo!")
		break
	}

	// If else ############

	// curly braces are compulsary
	n = 5
	if n == 5 {
		fmt.Println("yes it is 5!")
	}

	// can have a preceding statement(*new)
	// here, a is local var, available in else section too
	// NOTE! cant start else on next line, start after if's ending curly braces
	if a := 2; a == 2 {
		fmt.Println("a==2 TRUE")
	} else {
		fmt.Println("a!=2 False")
	}

	// Switch cases ###########
	i := 2
	fmt.Println("Write", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	}

	// switch as if else
	num := 20
	switch {
	case num < 10:
		fmt.Println("num < 10")
	case num > 10:
		fmt.Println("num > 10")
	}

	// Arrays ###################

	// cant initialize to some values({1, 2,3}) using var
	var arr [5]int
	fmt.Println("Curr- ", arr)
	arr[0] = 1
	arr[1] = 10
	arr[3] = 20
	fmt.Println("Updated- ", arr)

	arr2 := [4]int{1, 2, 3, 4}
	fmt.Println("Second- ", arr2)

	var dp [4][3]int
	fmt.Println(dp)

	fmt.Println(len(arr))
	fmt.Println(len(dp), len(dp[0]))

	// Slices ###################
	// advance version of arrays
	// if you give size, then its array if not, then slice
	// can use append, copy on slice not arrays

	d := []string{} //OR var d []string
	d = append(d, "one")
	fmt.Println(d)

	// another way
	sl := make([]string, 2)
	sl[0] = "one"
	sl[1] = "two"
	fmt.Println(sl)

	sl2 := make([]string, 0)
	sl2 = append(sl, "three", "four")
	fmt.Println(sl2)

	f := make([]string, len(sl2))
	copy(f, sl2) // copy from to, sl2 -> f
	fmt.Println(f)

	// can slice the slices OR can slice arrays
	values := [5]int{2, 3, 4, 5, 6}
	fmt.Println(values)

	// values[low:high]extract value from index low to high-1
	subValue := values[2:3] // cant do [2:2] OR [2:1], invalid
	fmt.Println(subValue)

	subValue = values[:5]
	fmt.Println(subValue)

	subValue = values[0:]
	fmt.Println(subValue)

	subValue = values[3:]
	fmt.Println(subValue)

	// Maps #################

	// using make to create a slice
	demoSlice := make([]string, 3)
	fmt.Println(demoSlice)

	// for map[key]value
	mp := make(map[int]int)
	mp[1]++
	mp[1]++
	mp[3]++
	mp[4] = 5
	fmt.Println(mp)

	// to declare & init in single line
	mp2 := map[string]int{"one": 1, "two": 2}

	// deleting a key,value
	delete(mp, 4)
	fmt.Println(mp)

	// getting a value, it returns value, bool, here bool is optional
	val := mp2["one"]
	fmt.Println(val)

	val2, found := mp2["three"]
	fmt.Println(val2, found)

	// check this cool thing, how to check if present and print
	// here (found) refers to (found==true)
	if val, found = mp2["two"]; found {
		fmt.Println("two is there")
	}

	// Range ###############
	// for iterating over various data structures
	// returns index, value,
	// if anyone is not needed simply put _(blank identifier) there

	arr3 := [5]int{3, 4, 5, 6, 7}
	for i, v := range arr3 {
		fmt.Println(i, v)
	}
	for i, _ := range arr3 { // can also be for  _, v :=range arr3
		fmt.Println(i)
	}

	mp4 := map[int]int{2: 3, 4: 5, 6: 1}
	for k, v := range mp4 {
		fmt.Println(k, v)
	}
}
