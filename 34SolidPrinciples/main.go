package main

import (
	"fmt"

	"github.com/adasarpan404/solidprinciples/openclosedprinciple"
	"github.com/adasarpan404/solidprinciples/singleresponsiblity"
)

func main() {

	fmt.Println("Single Responsibility \n Without")
	singleresponsiblity.NoSingleResponsiblityGo()
	fmt.Println("After implementing single responsiblity")
	singleresponsiblity.SingleResponsiblityGo()
	fmt.Println("Open Closed Responsibility")
	openclosedprinciple.OpenClosedPrinciple()
}
