package factory

type (
	AbstractFactoryFactory interface {
		NewProductA() AbstractFactoryProductA
		NewProductB() AbstractFactoryProductB
	}

	AbstractFactoryProductA interface {
		AbstractFactoryProductAFeature()
	}

	AbstractFactoryProductB interface {
		AbstractFactoryProductBFeature()
	}

	AbstractFactoryFactoryTypeX struct{}
	AbstractFactoryFactoryTypeY struct{}

	AbstractFactoryProductTypeXA struct{}
	AbstractFactoryProductTypeXB struct{}
	AbstractFactoryProductTypeYA struct{}
	AbstractFactoryProductTypeYB struct{}
)

func (p *AbstractFactoryProductTypeXA) AbstractFactoryProductAFeature() {}
func (p *AbstractFactoryProductTypeXB) AbstractFactoryProductBFeature() {}
func (p *AbstractFactoryProductTypeYA) AbstractFactoryProductAFeature() {}
func (p *AbstractFactoryProductTypeYB) AbstractFactoryProductBFeature() {}

func (f *AbstractFactoryFactoryTypeX) NewProductA() AbstractFactoryProductA {
	return new(AbstractFactoryProductTypeXA)
}

func (f *AbstractFactoryFactoryTypeX) NewProductB() AbstractFactoryProductB {
	return new(AbstractFactoryProductTypeXB)
}

func (f *AbstractFactoryFactoryTypeY) NewProductA() AbstractFactoryProductA {
	return new(AbstractFactoryProductTypeYA)
}

func (f *AbstractFactoryFactoryTypeY) NewProductB() AbstractFactoryProductB {
	return new(AbstractFactoryProductTypeYB)
}

func NewAbstractFactoryFactory(factoryType string) AbstractFactoryFactory {
	switch factoryType {
	case "typeA":
		return new(AbstractFactoryFactoryTypeX)
	case "typeB":
		return new(AbstractFactoryFactoryTypeY)
	}

	return nil
}
