package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

const msg = "Hello, OTUS!"

func main() {
	fmt.Println(stringutil.Reverse(msg))
}
