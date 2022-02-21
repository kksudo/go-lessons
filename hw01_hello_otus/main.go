package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

var hi = "Hello, OTUS!"

func main() {
	fmt.Printf("%s", stringutil.Reverse(hi))
}
