package main

import "fmt"

func modify(number *int) {
	*number *= 2
}

func main() {
	x := 2
	y := x
	z := &x //address of x
	fmt.Println(x, y)
	y = 3
	fmt.Println(x, y)

	*z = 69
	fmt.Println(x)
	/*
		USECASES:-
			1) for allowing functions to mutate the value of variable passes as an argumnet
			2) to improve efficieny of program
	*/

	myList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(myList)

	for i := 0; i < len(myList); i++ {
		modify(&myList[i])
	}

	fmt.Println(myList)
}
