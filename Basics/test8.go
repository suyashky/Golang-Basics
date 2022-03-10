package main

import "fmt"

func main() {

	jobs := make(chan string, 5) // buffered

	jobs <- "A"
	jobs <- "B"
	jobs <- "C"
	jobs <- "D"

	close(jobs)

	// can range over the values stored in channels
	// only when channel is closed
	for j := range jobs {
		fmt.Println(j)
	}

}
