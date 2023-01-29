package factory

type (
	MethodFactoryFactory interface {
		NewProductA() MethodFactoryProductA
		NewProductB() MethodFactoryProductB
	}

	MethodFactoryProductA interface {
		MethodFactoryProductAFeature()
	}

	MethodFactoryProductB interface {
		MethodFactoryProductBFeature()
	}
)
