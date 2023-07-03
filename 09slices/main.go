package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on SLICES")

	var fruitList = []string{"Peach", "Papaya", "Apple"}

	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")

	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])

	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867

	highScores = append(highScores, 555, 666, 321, 111)

	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	sort.Ints(highScores)

	fmt.Println(highScores)

	var courses = []string{"reactjs", "javascript", "swift", "ruby"}

	fmt.Println(courses)

	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)
}
