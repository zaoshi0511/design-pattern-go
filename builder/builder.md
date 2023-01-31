# Builder Pattern 建造者模式         

---
            
```go
package builder

type (
	Director struct {
		ConcentrateBuilder CityBuilder
	}

	City interface {
		Road()
		House()
	}

	CityBuilder interface {
		BuildRoad()
		BuildHouse()
		Complete() City
	}

	MedievalCityBuilder struct{}

	ModernCityBuilder struct{}
)

func (b *MedievalCityBuilder) BuildRoad()  {}
func (b *MedievalCityBuilder) BuildHouse() {}
func (b *MedievalCityBuilder) Complete() City {
	return nil
}

func (b *ModernCityBuilder) BuildRoad()  {}
func (b *ModernCityBuilder) BuildHouse() {}
func (b *ModernCityBuilder) Complete() City {
	return nil
}

func NewBuilder(builder CityBuilder) *Director {
	return &Director{
		ConcentrateBuilder: builder,
	}
}

func (b *Director) Build() City {
	b.ConcentrateBuilder.BuildRoad()

	b.ConcentrateBuilder.BuildHouse()

	return b.ConcentrateBuilder.Complete()
}
```

在建造者模式中，通过代码定义了最终的建造产物 `City`           

对于一个 `City` 可能有不同的风格，因此实例化两个具体的建造者，中世纪建造者 `MedievalCityBuilder` 和 现代建造者 `ModernCityBuilder`

对于 `Director` 可以理解为一个建造公司，每建造一个 `City` 都需要具体的 `ConcentrateBuilder` 去统领全局     

但无论有没有 `Director`，我们都可以直接创建某个具体的 `Builder` 去执行建造过程，去得到一个 `City`
```go
medievalCityBuilder := new(MedievalCityBuilder)
medievalCityBuilder.BuildRoad()
medievalCityBuilder.BuildHouse()
medievalCityBuilder.Complete()
```     

但对于建造出的最终产物 `City` 不同的 `XXXBuilder` 建造出的风格各不相同，也就是实际过程中的不同的 Builder 建造某一部分的具体细节有较大差异

且建造某一部分的过程是**对外部公开**的，创建 builder 实例的人可以手动修改建造流程，比如先执行 `BuildHouse()` 再执行 `BuildRoad()`      

在这种情况下，虽然都能建造出一个 `City` 但是在流程上是十分混乱的

而我们使用一个 `Director` 来管理具体的建造者，将其建造的过程封装在 `Director` 内，使建造的流程的规范统一

**建造者模式的优点**
1. 对建造的过程作了有效的封装，调用者无需关心内部的建造过程(或者说通过 `Director` 暴露的建造过程可以清晰的了解统一的规范过程)

2. 每个具体的建造者之间相互独立，不同的建造者按需实现自己需要的接口，且自己实现的接口内部不受其他建造者影响

**建造者模式的缺点**
1. 对于 `Builder` 接口，需要该接口的实现有大量共同点，否则建造出的最终产物会有很大限制
   > 举个例子，`MedievalCityBuilder` 在 `City` 中可以建造农场 `Farm`，这就要求 `MedievalCityBuilder` 需要有 `BuildFarm()` 方法             
   > 
   > 但对于 `ModernCityBuilder` 没法建造 `Farm` 但是可以建造 `Bank` 或者 `Hospital`，这就要求 `ModernCityBuilder` 需要有 `BuildBank()` 和 `BuildHospital()` 方法
   >        
   > 当具体的建造者某些特殊需求需要体现在最终产物时，只能将建造过程添加到公共接口上，但对于不需要该功能的具体建造者，只能实现该空接口