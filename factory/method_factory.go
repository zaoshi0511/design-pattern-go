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
