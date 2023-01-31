# Adapter Pattern 适配器模式     

---

```go
package adapter

type (
	Slot interface {
		Charge()
	}

	Adapter struct {
		slot Slot
	}

	USB struct {
	}

	TypeC struct {
	}
)

func (u *USB) Charge() {}

func (t *TypeC) Charge() {}

func (a *Adapter) SetSlot(slot Slot) {
	a.slot = slot
}

func (a *Adapter) Charge() {
	a.slot.Charge()
}
```

在适配器模式中，核心思想是，对外暴露统一接口 `Charge` 再在内部分发到不同的处理器处理 `USB.Charge()` 和 `TypeC.Charge()`           

像在 go 中通常会对 `interface` 进行断言，本质上也是一种适配，为了符合当前函数中的变量类型，要通过断言将 interface 转化成对应的数据类型

对外部仅暴露 `Adapter.Charge()` 接口，客户端根据自己的需要去选择具体的 `Slot`，而实际的 `Charge()` 执行也是由具体的 `Slot` 完成