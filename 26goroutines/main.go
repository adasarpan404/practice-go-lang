package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup //pointer

func main() {
	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()

}

func getStatusCode(endPoint string) {

	defer wg.Done()
	res, err := http.Get(endPoint)

	if err != nil {
		fmt.Println("OOPS in endpoint")
	}
	fmt.Printf("%d status for %s \n", res.StatusCode, endPoint)
}
