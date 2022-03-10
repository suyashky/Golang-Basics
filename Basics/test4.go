package main

import (
	"fmt"
	"time"
)

func f(N int) {
	for i := 1; i < N; i++ {
		// sleep for 1 sec, before next step of execution
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

func main() {

	// GO routines
	// way to achieve concurrency
	// they are light weight threads

	// sequential calls, will run synchronously
	// synchronous-> later fn will wait for return from first call
	// then start its own call

	// right now nothing will print
	// bcz control will return from main before their execution
	go f(4)
	go f(4)

	// using it to stop main from ending-> enter a char to end it
	fmt.Scanln()

	// now concurrency is acheived, both fc runns simultaneously/
	// asynchronously->
}
