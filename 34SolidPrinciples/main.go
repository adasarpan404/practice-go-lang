package main

import (
	"fmt"

	"github.com/adasarpan404/solidprinciples/dependencyinversionprinciple"
	"github.com/adasarpan404/solidprinciples/interfaceseggregationprinciple"
	"github.com/adasarpan404/solidprinciples/liskovsubstitutionprinciple"
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
	fmt.Println("Liskov Substitution Principle")
	liskovsubstitutionprinciple.LiskovSubstitution()
	fmt.Println("Interface Seggregation Principle")
	fmt.Println("Without")
	interfaceseggregationprinciple.NoInterfaceSeggregation()
	fmt.Println("With")
	interfaceseggregationprinciple.InterfaceSeggregation()
	fmt.Println("Dependency Inversion Principle")
	dependencyinversionprinciple.DependencyInversion()
}
