package main

import (
	"fmt"
	"net/http"
)

// handler function
// take renponseWriter(an interface) and
// httprequest of type struct

// a handler
// its for writing a response to responseWriter
func hello(w http.ResponseWriter, req *http.Request) {

	// fprint used to write to responseWriter
	fmt.Fprint(w, "hello sir!")

}

// another handler function
// use to add all HTTP request headers to response body
func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	// register our handler functions on server route
	// sets up default router, take a func as argument
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// to use the route that we have setup
	http.ListenAndServe(":8090", nil)

}
