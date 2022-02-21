package main

import (
	"fmt"
	"golang.org/x/example/stringutil"
)

var hi = "Hello, OTUS!"

func main() {
	fmt.Printf("Main direction:%s\n", hi)
	fmt.Printf("Reverse direction: %s\n", stringutil.Reverse(hi))
}
