package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	// A Context carries deadlines,
	// cancellation signals, and other request-scoped
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	// we just gonna wait 10 sec, to simulate some work
	// once ctxdone is called it retuns

	// ctx.done channel can make return soon
	// it just cancel the curr work and return asap
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}

}

func main() {
	//register our handler on the “/hello” route,
	// and start serving.
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
