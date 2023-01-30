# Factory pattern 工厂模式          

---

## Simple Factory 简单工厂模式            
        
```go
package facotory

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
```

在简单工厂模式中，**工厂是具体的，产品是抽象的**，产品的具体细节，依赖于外部传入的参数       

对于抽象的产品，我们定义其应该具备的一些功能，`SimpleFactoryProductMethod()`       

举个例子，假如产品是浏览器，那么所有的抽象浏览器都应该具有 `浏览网页()` 这样一个功能           

工厂通过客户端传入的 `浏览器名称(productType)` 来选择使用什么浏览器浏览网页              

**简单工厂的优点**

1. 类型的 **创建与使用分离**，当需要某个类时，仅传入 **需要设定的参数**，获取到工厂返回值直接使用即可

**简单工厂的缺点**

1. 每当工厂要添加新的 "产品"，都需要单独去写一份额外的生产代码，违反了开闭原则
2. 如果工厂种类少，产品类型多，则容易出现 **工厂生产了与自己关系不大的产品**，有违背 **单一指责** 原则
3. 当工厂的 **生产方法** 存在安全隐患或者不可用时，所有依赖该工厂生产的类都会受到影响


## Method Factory 工厂方法模式            
        
```go
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
```

在工厂方法模式中，**工厂是抽象的，产品是抽象的**，产品的具体细节，依赖于生产该产品的工厂      

在简单工厂模式中，我们只能向单一的具体的工厂给出参数，以获得对应的产品；而在工厂方法模式中，我们可以选择工厂      

参数由原本对 `产品类型` 的指定，转变为对 `工厂类型` 的指定       

举个例子，如果说简单工厂模式是在一个电脑店内购买电脑，这个电脑店出售联想、戴尔、宏碁、华硕等品牌的电脑，我们进店只需要向店员说明我们需要什么品牌的电脑，店员就会向我们介绍对应的电脑          

而工厂方法模式就是在电脑城中，选择指定品牌的专卖店，以华硕为例，只要我们进到华硕的专卖店，那么该电脑店的任意一台电脑都是华硕电脑        

**工厂方法的优点**

1. 进一步解耦，从简单工厂的 **给出参数 -> 创建类型** 到 **给出参数 -> 调用子类的创建方法 -> 创建类型**
2. 可扩展性强，当需要新增工厂时，仅需要实现工厂的生产方法即可，不会对原有的代码做出修改
3. 由于生产的类型是在子类中决定的，因此遵守了 **单一指责** 原则

**工厂方法的缺点**

1. 每当需要添加新的产品时，都需要添加新的工厂去继承方法，代码量会增多。当某类产品使用次数较少时，也会增加同体量的代码，利用率不高      

## Abstract Factory 抽象工厂模式      
        
```go
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

```

在抽象工厂模式中，**工厂是抽象的，产品是抽象的**，产品的**特性**，依赖于生产该产品的工厂        

在简单工厂和工厂方法模式中，对于一个**具体的工厂**，都只能生产一种产品       

为了使一个工厂能够生产多种产品，在工厂方法模式中就需要定义多个抽象的产品            

这些由同一个工厂生产出的产品的集合，可以称为一个产品族(series)，这些产品自带的有着 `生产自某工厂` 的标签      

借着工厂方法模式中的例子，现在不是买电脑，而是购买鼠标、键盘、显示器、硬盘、内存这类的配件。在普通的电脑店中，什么样的配件都有，如果你需要某个品牌的具体配件，比如 **华硕的键盘**，这是给出了两个参数，品牌和配件类型           

如果是进入到专卖店购买配件，那么在进入专卖店时，就已经选择了品牌，接下来只要给出对应的配件类型；当然只要你还在这个专卖店中，无论你购买何种配件类型，这个配件类型的品牌都是不变的

**抽象工厂的优点**

1. 当工厂要增加生产的产品类型时，只用在抽象工厂接口中增加对应方法即可

**抽象工厂的缺点**

1. 如果需要对产品按照其他类型分类，抽象工厂的实现类会横向拓展

2. 如果某些工厂需要单独处理，不容易定制化；如果在抽象工厂中直接添加方法，又会导致其实现类增加无用的方法
