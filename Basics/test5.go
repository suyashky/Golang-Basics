// Channel synchronization
// can use <-chan to block until our single go routine finishes
// if there are multiple go routines use wait, discussed later

package main

import (
	"fmt"
	"time"
)

// this will run in go routine
// will send a bool value to channel once done
// this channel can be used to tell another func that its work is done
func worker(done chan bool) {
	fmt.Println("wroking..")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
	//hell yeah
}

func main() {
	// buffered channel of capacity 1
	done := make(chan bool, 1)
	go worker(done)

	// this receive is put here
	// so wait for go routine to finish
	<-done

	// program will go to this line onl when go routine is finished
	// upper line works as blocking receiver,
	// it blocks any other opration until goroutne finishes
	fmt.Println("done")
}
