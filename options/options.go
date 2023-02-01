package options

type (
	OptionFunc    func(value interface{}) func(obj *OptionsObject)
	OptionsObject struct {
		Field1 int
		Field2 string
	}
)

func NewOptionsObject(opts ...OptionFunc) *OptionsObject {
	obj := new(OptionsObject)

	for _, f := range opts {
		f(obj)
	}

	return obj
}

func ObjectWithField1(field1 interface{}) func(obj *OptionsObject) {
	return func(obj *OptionsObject) {
		value := field1.(int)
		obj.Field1 = value
	}
}

func ObjectWithField2(field2 interface{}) func(obj *OptionsObject) {
	return func(obj *OptionsObject) {
		value := field2.(string)
		obj.Field2 = value
	}
}
