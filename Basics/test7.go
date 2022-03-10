package main

import "fmt"

func task(jobs chan string, done chan bool) {

	// here, (j, more) is a two valued return from channel
	// it sends false to bool value if that channel is closed
	for {
		j, more := <-jobs
		if more {
			fmt.Println("receiving: ", j)
		} else {
			fmt.Println("received all jobs!")
			done <- true // marking task as done, will await for it in main's end
			return
		}
	}
}

func main() {
	jobs := make(chan string, 5) // buffered
	done := make(chan bool)

	jobs <- "A"
	jobs <- "B"
	jobs <- "C"
	jobs <- "D"
	jobs <- "E"

	go task(jobs, done)

	close(jobs)

	fmt.Println("all jobs are sent!")

	// awaiting
	<-done
}
