package main

import "fmt"

func main() {
	fmt.Println("Welcome to my array")

	var fruitList [4]string

	fruitList[0] = "apple"
	fruitList[1] = "banana"
	fruitList[2] = "coconut"
	fruitList[3] = "mango"
	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list is: ", len(fruitList))

	var vegList = [5]string{"potato", "beans", "mushroom"}
	fmt.Println("Vegy list is: ", len(vegList))
}
