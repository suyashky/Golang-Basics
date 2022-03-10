package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	// give any number of argumnets, it forms path
	// automatically removes extra //

	// can also get directory and file name seperated
	p := filepath.Join("c//", "//temp//", "res", "pics")
	fmt.Println(p, filepath.Dir(p), filepath.Base(p))

}
