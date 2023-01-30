package factory

type (
	MethodFactoryFactory interface {
		NewProduct() MethodFactoryProduct
	}

	MethodFactoryProduct interface {
		MethodFactoryProductMethod()
	}

	MethodFactoryFactoryA struct{}
	MethodFactoryFactoryB struct{}

	MethodFactoryProductA struct{}
	MethodFactoryProductB struct{}
)

func (f *MethodFactoryFactoryA) NewProduct() MethodFactoryProduct {
	return new(MethodFactoryProductA)
}

func (f *MethodFactoryFactoryB) NewProduct() MethodFactoryProduct {
	return new(MethodFactoryProductB)
}

func (p *MethodFactoryProductA) MethodFactoryProductMethod() {}

func (p *MethodFactoryProductB) MethodFactoryProductMethod() {}

var (
	MethodFactoryFactoryAInstance = new(MethodFactoryFactoryA)
	MethodFactoryFactoryBInstance = new(MethodFactoryFactoryB)
)

func NewMethodFactoryFactory(factoryType string) MethodFactoryProduct {
	switch factoryType {
	case "typeA":
		return MethodFactoryFactoryAInstance.NewProduct()
	case "typeB":
		return MethodFactoryFactoryBInstance.NewProduct()
	}
	return nil
}
