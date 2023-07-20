package builder

import "fmt"

type Vehicle interface {
	Drive()
}

type Car struct{}

func (c *Car) Drive() {
	fmt.Println("Driving a car.")
}

type Bike struct{}

func (b *Bike) Drive() {
	fmt.Println("D.")
}

type Plane struct{}

func (p *Plane) Drive() {
	fmt.Println("Flying a plane.")
}

type VehicleFactory interface {
	CreateVehicle() Vehicle
}

type CarFactory struct{}

func (cf *CarFactory) CreateVehicle() Vehicle {
	return &Car{}
}

type BikeFactory struct{}

func (bf *BikeFactory) CreateVehicle() Vehicle {
	return &Bike{}
}

type PlaneFactory struct{}

func (pf *PlaneFactory) CreateVehicle() Vehicle {
	return &Plane{}
}

func Builder() {
	carFactory := &CarFactory{}
	car := carFactory.CreateVehicle()
	car.Drive()

	// Create a bike using BikeFactory
	bikeFactory := &BikeFactory{}
	bike := bikeFactory.CreateVehicle()
	bike.Drive()

	// Create a plane using PlaneFactory
	planeFactory := &PlaneFactory{}
	plane := planeFactory.CreateVehicle()
	plane.Drive()
}
