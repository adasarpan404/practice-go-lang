package liskovsubstitutionprinciple

import "fmt"

type Vehicle interface {
	Start()
}

type Car struct{}

func (c Car) Start() {
	fmt.Println("Car starting...")
}

type Bike struct{}

func (b Bike) Start() {
	fmt.Println("Bike starting...")
}

func startVehicle(vehicle Vehicle) {
	vehicle.Start()
}

func LiskovSubstitution() {
	car := Car{}
	bike := Bike{}

	startVehicle(car)
	startVehicle(bike)
}
