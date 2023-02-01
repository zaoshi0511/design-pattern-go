package options

type OptionFunc2 func(obj *OptionsObject, value interface{})

func ObjectWithField(obj *OptionsObject, field interface{}) {
	value := field.(int)
	obj.Field1 = value
}

func NewOptionsObject2(params []interface{}, opts ...OptionFunc2) {
	obj := new(OptionsObject)
	for k, f := range opts {
		f(obj, params[k])
	}
}
