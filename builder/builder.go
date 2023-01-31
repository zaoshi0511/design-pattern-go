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
