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

func (f *SimpleFactoryFactory) NewSimpleFactoryProduct(productType string) SimpleFactoryProduct {

	switch productType {
	case "A":
		return new(SimpleFactoryProductA)
	case "B":
		return new(SimpleFactoryProductB)
	}
	return nil
}
