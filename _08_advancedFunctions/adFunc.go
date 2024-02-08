package main

import "fmt"

//The function that accepts another function as an argument is called high order function
//The function that is passed as an parameter is called first class function
//Currying-> taking a function as an argument and modifying or enhancing the function and returning the function

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

//High order function
func aggregate(a, b, c int, arithmetic func(a, b int) int) int {
	return arithmetic(arithmetic(a, b), c)
}

//Currying

func findSquare(a int, multiply func(a, b int) int) int {
	return multiply(a, a)
}

func main() {
	fmt.Println(aggregate(2, 3, 4, add))
	fmt.Println(findSquare(8, multiply))
}
