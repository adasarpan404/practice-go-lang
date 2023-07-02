package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
	// "time"
)

func main() {
	fmt.Println("Welcome to mymaths")

	// rand.Seed(time.Now().Unix())
	// fmt.Println(rand.Intn(5) + 1)
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
}
