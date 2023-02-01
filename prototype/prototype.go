package prototype

type (
	prototypeObject struct {
		Field1 int
		Field2 string
	}
)

func (o *prototypeObject) Clone() *prototypeObject {
	newObj := *o
	return &newObj
}

func GetPrototypeObject() *prototypeObject {
	return new(prototypeObject)
}
