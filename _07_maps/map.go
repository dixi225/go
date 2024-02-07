package main

import "fmt"

func main() {
	fruitPrices := make(map[string]int)
	fruitPrices["Apple"] = 80
	fruitPrices["Banana"] = 80
	fruitPrices["Orange"] = 80
	fruitPrices["Kivi"] = 80

	//

	newFruitPrices := map[string]int{
		"Apple":  60,
		"Banana": 80,
		"Orange": 10,
		"Kivi":   20,
	}

	fmt.Println(len(newFruitPrices))

	//Inserting

	newFruitPrices["Mango"] = 100

	fmt.Println(len(newFruitPrices))

	//Deleting

	delete(newFruitPrices, "Kivi")

	fmt.Println(len(newFruitPrices))

	//checking if an element exist or not

	_, ok := newFruitPrices["Mango"]

	if ok {
		fmt.Println("Exist")
	} else {
		fmt.Println("Doesn't Exist")
	}
}
