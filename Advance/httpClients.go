package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	response, err := http.Get("https://mail.google.com/mail/u/2/#inbox")
	check(err)
	defer response.Body.Close()

	// return status code
	// 200- success
	// 400- error
	fmt.Println("Response status: ", response.Status)

	// buffered scanner
	scanner := bufio.NewScanner(response.Body)
	// scanner.scan returns nil when EOF
	for i := 0; scanner.Scan() && i < 10; i++ {
		fmt.Println(scanner.Text())
	}
	check(scanner.Err())

}
