package main

import (
	"fmt"
)

type rectangle struct {
	length  int
	breadth int
}

func area(r rectangle) int {
	return r.breadth * r.length
}

// dynamic attribute of a struct
func (r rectangle) perimeter() int {
	return 2 * (r.length + r.breadth)
}

func main() {
	rect := rectangle{
		length:  8,
		breadth: 8,
	}
	fmt.Println(area(rect))
	fmt.Println(rect.perimeter())
}
