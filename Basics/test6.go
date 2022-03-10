// select
// say we have multiple channels
// we want to pick one at a time for sending/ receiving
// use select

// Here we have a basic example
package main

import "fmt"

func greet() chan string {
	pipe := make(chan string, 4)

	// we have capacity = 4
	// need loop to select 4 items
	// it selects different item in differnet itearion "mostly"

	for i := 1; i <= 4; i++ {
		select {
		// they contain all possible task
		// from which one has to be picked everytime
		case pipe <- "Hello":
		case pipe <- "There":
		case pipe <- "Bye":
		}
	}
	close(pipe)
	return pipe
}

func main() {
	pipe := greet()
	for i := range pipe {
		fmt.Println(i)
	}
	// so, 4 items out of all those were in selct
	// on random basis are selected and returned
}
