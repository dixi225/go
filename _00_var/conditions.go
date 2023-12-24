package main

import "fmt"

func main() {
	response := ""
	fmt.Println("Are you missing Janvi?")
	fmt.Scan(&response)
	if response == "Yes" {
		fmt.Println("I know , she's cute")
	} else {
		fmt.Println("Stop Lying")
	}

	//Initial if-else block

	if length := 3; length == 2 {
		fmt.Println("Length is 2")
	} else {
		fmt.Println("Length is not 2")
	}

	// fmt.Println(length) to limit the scope of the vairable

}
