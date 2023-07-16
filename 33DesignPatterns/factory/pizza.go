package factory

type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
}
