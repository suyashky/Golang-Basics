// topics
// Functions
// Multiple return type
// variadic fn
// closures

package main

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func opr(a, b int) (int, int) {
	return (a + b), (a - b)
}

func sum(nums ...int) int {
	total := 0
	for _, val := range nums {
		total += val
	}
	return total
}

// this fn has a anonymous fn defines inside which return int
// this fn returns a fn, again call whatever you get as return from it

// func() int
// this is the return type- a fn with no name returning int
func check() func() int {
	i := 0
	// below fn closes over i, so form closure
	return func() int {
		i++
		return i
	}
}

func main() {

	//  Functions #############

	res := add(2, 3)
	println(res)
	println(sub(5, 3))

	// Returning multiple values ##########
	a, b := opr(20, 10)
	println(a, b)
	a, _ = opr(10, 5)
	_, b = opr(51, 51)
	println(a, b)

	// Variadic functions #########

	// i.e. a function that can have any number of arguments
	// ex. a print function can have any arguments
	println(sum(1, 2))
	println(sum(1, 2, 3, 4))

	// If you have a slice of arguments
	set := []int{2, 3, 4}
	println(sum(set...))

	// Closures ############

	// anonymous function-> when we need a function inline without naming it
	return1 := check()
	println(return1())
	println(return1())
	println(return1())

	return2 := check()
	println(return2())

	//
}
