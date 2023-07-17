package factory

func FactoryUsage() {
	cheesePizzaFactory := &CheesePizzaFactory{}
	cheesePizza := cheesePizzaFactory.CreatePizza()
	cheesePizza.Prepare()
	cheesePizza.Bake()
	cheesePizza.Cut()
	cheesePizza.Box()

	pepperoniPizzaFactory := &PepperoniPizzaFactory{}
	pepperoniPizza := pepperoniPizzaFactory.CreatePizza()
	pepperoniPizza.Prepare()
	pepperoniPizza.Bake()
	pepperoniPizza.Cut()
	pepperoniPizza.Box()
}
