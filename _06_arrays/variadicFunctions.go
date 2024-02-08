package main

import (
	"fmt"
)

func sum(num ...int) (ans int) {
	for i := 0; i < len(num); i++ {
		ans += num[i]
	}
	return ans
}

func main() {
	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(mySlice...))

	mySlice = append(mySlice, 6)
	fmt.Println(sum(mySlice...))

	mySlice = append(mySlice, []int{7, 8, 9}...) // append is a variadic function
	fmt.Println(sum(mySlice...))

}
