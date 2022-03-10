package main

import (
	"fmt"
	"net/url"
)

func main() {
	sample := "http://suyashky:deko12k0@stagesite.com:8080/C://file/doc?dob=2-11-2000#anchor"

	u, err := url.Parse(sample)
	if err != nil {
		panic(err)
	}

	// authentication info
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	fmt.Println(u.User.Password()) // (password,bool)

	// host & port
	fmt.Println(u.Host)
	fmt.Println(u.Hostname())
	fmt.Println(u.Port())

	// path & fragment
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// query params
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery) // returna a map
	fmt.Println(m)
	fmt.Println(m["dob"])
}
