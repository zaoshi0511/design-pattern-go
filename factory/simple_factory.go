package factory

type (
	SimpleFactoryFactory struct{}

	SimpleFactoryProduct interface {
		SimpleFactoryProductMethod()
	}

	SimpleFactoryProductA struct{}

	SimpleFactoryProductB struct{}
)

func (p *SimpleFactoryProductA) SimpleFactoryProductMethod() {}
func (p *SimpleFactoryProductB) SimpleFactoryProductMethod() {}

func (f *SimpleFactoryFactory) NewProduct(param string) SimpleFactoryProduct {

	switch param {
	case "A":
		return new(SimpleFactoryProductA)
	case "B":
		return new(SimpleFactoryProductB)
	}
	return nil
}
