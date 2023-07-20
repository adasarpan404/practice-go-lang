package openclosedprinciple

import "fmt"

// Entity interface
type Entity interface {
	CalculateBMI() float64
}

// John entity
type John struct {
	height int
	weight int
}

func (j *John) CalculateBMI() float64 {
	return float64(j.height) / float64(j.weight)
}

// Jane entity
type Jane struct {
	height int
	weight int
}

func (j *Jane) CalculateBMI() float64 {
	return float64(j.height) / float64(j.weight)
}

// Dog entity
type Dog struct {
	height int
	weight int
}

func (d *Dog) CalculateBMI() float64 {
	return float64(d.height) / float64(d.weight)
}

func OpenClosedPrinciple() {
	// Creating instances of John, Jane, and Dog
	john := &John{height: 175, weight: 70}
	jane := &Jane{height: 160, weight: 55}
	dog := &Dog{height: 40, weight: 10}

	// Using the Entity interface to calculate BMI for different entities
	entities := []Entity{john, jane, dog}
	for _, entity := range entities {
		bmi := entity.CalculateBMI()
		fmt.Printf("BMI: %.2f\n", bmi)
	}
}
