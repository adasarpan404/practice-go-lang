package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files")

	content := "This need to go in a file - LCO"

	file, err := os.Create("./arpan.arp")

	checkNilError(err)
	length, err := io.WriteString(file, content)

	checkNilError(err)
	fmt.Println("Length is: ", length)

	defer file.Close()

	readFile(file.Name())
}

func readFile(filename string) {
	dataBytes, err := ioutil.ReadFile(filename)
	checkNilError(err)

	fmt.Println("Text data inside the file is \n", string(dataBytes))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
