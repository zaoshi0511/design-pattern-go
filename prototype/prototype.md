# Prototype Pattern 原型模式

---

```go
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

```

在原型模式中，核心思想是，能够依照一个实例复制出一个新的实例，而无需知晓这个实例的具体类型是什么        

要复制一个实例，传统方式是，知道这个实例的具体类型，创建一个该类型的新的实例，再将要复制实例的属性拷贝进这个新实例中      

在传统方式中，很重要的一步是 **知道这个实例的类型**，在某些情况下，有些实例的类型我们是无法访问的，但我们可通过实例类型暴露的某些方法来修改实例，此时我们就可以使用原型模式在无法操作类型的情况下创建一个新的实例

在上述代码中，我将 `prototypeObject` 首字母设置为小写，则在跨包时，不能直接创建 `prototypeObject` 实例，只能通过 `Clone()` 创建新的实例