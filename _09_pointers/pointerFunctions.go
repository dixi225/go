package main

import "fmt"

type car struct {
	color string
}

func (c *car) changeColor(color string) {
	c.color = color
}

func main() {
	honda := car{
		color: "black",
	}

	fmt.Println(honda.color)

	honda.changeColor("red")

	fmt.Println(honda.color)
}
