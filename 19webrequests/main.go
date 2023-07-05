package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://github.com/pocketbase"

func main() {
	fmt.Println("Welcome to web requests")

	response, err := http.Get(url)

	checkNilError(err)

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close()

	dataBytes, err := ioutil.ReadAll(response.Body)
	checkNilError(err)

	content := string(dataBytes)

	fmt.Println("Content is:", content)

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
