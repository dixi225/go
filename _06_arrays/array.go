package main

import "fmt"

func main() {
	var myOldArr [10]int
	fmt.Println(myOldArr)
	newArr := [3]string{"hi", "harsh", "here"}
	fmt.Println(newArr)
	// arrays are of static size , but we have slices, that's notthing but the address of starting index to ending index(not inlusive)

	//slices are reference of arrays

	newArrSlice := newArr[0:2]
	fmt.Println(newArrSlice)
	newArrSlice[0] = "hello"
	fmt.Println(newArr)

	//Slices without explicitly defining an array - Make keyword

	// mySlice := make([]int, 5, 10)

	mySlice := make([]int, 5)
	fmt.Println(cap(mySlice))

	Brands := []string{"Apple", "Samsung", "Motrala", "Redmi", "Xiomi"}
	chineseBrands := []string{"Xiomi", "Redmi", "Realme"}
	count := 0
	for _, brand := range Brands {
		for _, cBrand := range chineseBrands {
			if brand == cBrand {
				count++
			}
		}
	}
	fmt.Println(count)
}
