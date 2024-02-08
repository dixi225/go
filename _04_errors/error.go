package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0.0, errors.New("No Divison by 0")
	}
	return a / b, nil
}

func main() {
	res, err := divide(3.0, 0.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
