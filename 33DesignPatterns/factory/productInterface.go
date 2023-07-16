package factory

import "fmt"

type CheesePizza struct{}

func (c *CheesePizza) Prepare() {
	fmt.Println("Preparing Cheese Pizza......")
}

func (c *CheesePizza) Bake() {
	fmt.Println("Baking Cheese Pizza......")
}

func (c *CheesePizza) Cut() {
	fmt.Println("Cutting Cheese Pizza......")
}

func (c *CheesePizza) Box() {
	fmt.Println("Boxing Cheese Pizza......")
}

type PepperoniPizza struct{}

func (p *PepperoniPizza) Prepare() {
	fmt.Println("Preparing Pepperoni Pizza......")
}

func (p *PepperoniPizza) Bake() {
	fmt.Println("Baking Pepperoni Pizza......")
}

func (p *PepperoniPizza) Cut() {
	fmt.Println("Cutting Pepperoni Pizza......")
}

func (p *PepperoniPizza) Box() {
	fmt.Println("Boxing Pepperoni Pizza......")
}
