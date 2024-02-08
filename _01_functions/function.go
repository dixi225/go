package main

import (
	"fmt"
)

func add(x, y int) (a, b int) {
	a = x
	b = x + y
	return a, b
	// return
}

func main() {
	x, _ := add(2, 3)
	fmt.Println(x)
}
