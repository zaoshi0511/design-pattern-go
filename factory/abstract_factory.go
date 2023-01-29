package factory

type (
	AbstractFactoryFactory interface {
	}

	AbstractFactoryProduct interface {
		AbstractFactoryProductDoSomethingA()
		AbstractFactoryProductDoSomethingB()
	}

	AbstractFactoryFactoryX struct{}
	AbstractFactoryFactoryY struct{}

	AbstractFactoryProductA struct{}
	AbstractFactoryProductB struct{}
)

func (p *AbstractFactoryProductA) AbstractFactoryProductDoSomethingA() {}
func (p *AbstractFactoryProductA) AbstractFactoryProductDoSomethingB() {}

func (p *AbstractFactoryProductB) AbstractFactoryProductDoSomethingA() {}
func (p *AbstractFactoryProductB) AbstractFactoryProductDoSomethingB() {}
