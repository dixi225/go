package main

import (
	"fmt"
)

func getMessege(fig shape) {
	_, ok := fig.(square) //checks if fig interface got implemented on square struct or not
	if ok {
		fmt.Println("square")
	}
	switch v := fig.(type) {
	case square:
		fmt.Println(v.side)

	case rectangle:
		fmt.Println(v.length)

	default:
		fmt.Println("other")
	}
	fmt.Println(fig.area())
	fmt.Println(fig.perimeter())

}

type shape interface {
	area() int
	perimeter() int
}

type rectangle struct {
	length  int
	breadth int
}

type square struct {
	side int
}

func (r rectangle) area() int {
	return r.breadth * r.length
}

func (r rectangle) perimeter() int {
	return 2 * (r.breadth + r.length)
}

func (s square) area() int {
	return s.side * s.side
}

func (s square) perimeter() int {
	return 4 * s.side
}
func main() {
	rect := rectangle{
		length:  2,
		breadth: 6,
	}
	getMessege(rect)

}
