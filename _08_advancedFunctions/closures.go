package main

import "fmt"

func parentFunction(a, b int) func(a int) int {

	//here c lies out side of scope of returning function but still,c exist in scope of function under the property of closures

	c := a + b
	return func(i int) int {
		return i + c
	}
}

func testFunction(a string, b func(string) string) {
	fmt.Println("argumented string is : " + a)
	fmt.Println("function/processed string is : " + b("harsh"))
}
func main() {
	a, b := 1, 2
	// c:=3
	temp := parentFunction(a, b)
	fmt.Println(temp(2))
	testFunction("hi", func(b string) string { //anonymous function
		return b + " here"
	})
}
