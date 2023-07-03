package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to maps")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages :", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	delete(languages, "RB")

	fmt.Println("List of all languages: ", languages)

	for _, value := range languages {
		fmt.Printf("For key v, value is %v\n", value)
	}

}
