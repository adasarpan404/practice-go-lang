package factory

type CheesePizzaFactory struct{}

func (cpf *CheesePizzaFactory) CreatePizza() Pizza {
	return &CheesePizza{}
}

type PepperoniPizzaFactory struct{}

func (ppf *PepperoniPizzaFactory) CreatePizza() Pizza {
	return &PepperoniPizza{}
}
