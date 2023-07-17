package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/adasarpan404/designPatterns/factory"
)

func main() {
	fmt.Println("Design Patterns in GO")

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	input = strings.TrimSpace(input) // Remove leading/trailing whitespaces

	if strings.EqualFold(input, "factory") {
		fmt.Println(true)
		factory.FactoryUsage()
	} else {
		fmt.Println("Invalid input.")
	}
}
