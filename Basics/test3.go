package main

import "fmt"

type newEmpAssignment interface {
	generateMail() string
}

type person struct {
	name string
	age  int
	mail string
}

// a var can have a type of interface
// use it to perform operation of every interface
func printMail(temp newEmpAssignment) {
	fmt.Println(temp)
}

// a fn-> func area(arg) int{ }
// method has a special receiver arguemnt
// this reveicer can be pointer or value both

func (someone person) generateMail() string {
	return (someone.name + "@onenote.com")
}

func main() {

	// Structures ############
	fmt.Println(person{"Bob", 20, ""})

	first := person{"suyash", 21, ""}
	fmt.Println(first.name, first.age)

	ptr := &first
	// no need for (*ptr).age, pointers to struct are auto dereferenced
	fmt.Println(ptr.age)

	// Method #############

	// to define fn for structs
	newEmp := person{"Danny", 25, ""}
	fmt.Println(newEmp.generateMail())

	// Interfaces ############
	// collection of method signatures /prototype
	printMail(newEmp)

}
