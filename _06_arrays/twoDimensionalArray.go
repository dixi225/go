package main

import (
	"fmt"
)

func main() {
	matrix := make([][]int, 0)
	for i := 0; i < 10; i++ {
		row := make([]int, 0)
		for j := 0; j < 10; j++ {
			row = append(row, j)
		}
		matrix = append(matrix, row)
	}
	fmt.Println(matrix)
}
