# Options Pattern 选项模式

---

```go
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
```

选项模式并非 23 中通用设计模式中的一种，而是函数式编程中的一种设计模式，全称为 **Functional Options Pattern**

对于 Go 而言，不能够方法重载，因此无法写出可变参数的构造函数，但是由于 Go 支持函数式编程，因此可以利用函数式编程中的 **选项模式** 来实现可变参的构造函数

为什么对于每个选项函数 `OptionFunc()` 要返回实际操作的函数？

如果对于选项函数 `OptionFunc()` 将实际的操作放在内部，则就必须要求外部额外传入一组属性值

```go
func ObjectWithField(obj *OptionsObject, field interface{}){
	value := field.(int)
	obj.IntField = value
}
```

这样对于构造函数而言就无法专注于传入的参数了，构造函数本身仅依赖对于外部的参数，返回给一个实例

```go
package options

type OptionFunc func(obj *OptionsObject, value interface{})

func ObjectWithField(obj *OptionsObject, field interface{}) {
	value := field.(int)
	obj.Field1 = value
}

func NewOptionsObject(params []interface{}, opts ...OptionFunc2) {
	obj := new(OptionsObject)
	for k, f := range opts {
		f(obj, params[k])
	}
}

```