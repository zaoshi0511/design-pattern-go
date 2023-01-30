# Proxy pattern 代理模式

---

```go
package proxy

type (
	Common interface {
		DoSomething()
	}

	Trustor struct{}

	Proxy struct {
		Agent Trustor
	}
)

func (t *Trustor) DoSomething() {}

func (p *Proxy) DoSomething() {

	p.BeforeDoSomething()

	p.Agent.DoSomething()

	p.AfterDoSomething()
}

func (p *Proxy) BeforeDoSomething() {}
func (p *Proxy) AfterDoSomething()  {}

func NewProxy() *Proxy {
	return &Proxy{}
}

```

在代理模式中，核心思想是: **代理类 `Proxy` 替被代理对象(我这里命名的是委托对象) `Trustor` 完成某件事**           

从解释的语义上其实不难理解       

对我个人而言，期初我很难理解的一点是，委托人(被代理对象)为什么要作为一个参数传入去生成一个代理对象，就像下面代码
```go
package proxy

func NewProxy(trustor *Trustor) *Proxy{
	return Proxy{
		Agent: trustor,
    }
}
```

现在想来，结合生活中的例子，既然某人要代替你去完成某件事，那必须**得有一个人去完成**，这个人就是事情的执行者            

举个更为明显的例子，我找人代替我去买火车票。购买火车票需要身份证，不然不能确定票是给谁买的       

我找的这个代理人，我需要将自己的身份证交给他去窗口购买，可以理解为我将自己的一份身份信息的拷贝传给了他，他拿着我的个人信息去执行我要做的事       

从外部来看，委托人 `Trustor` 和代理人 `Proxy` 都完成了这件事 `DoSomething()`，但最后实际生效的结果是给到了委托人 `Trustor`，就好像委托人自己去完成了这件事一样      

从内部来看，代理人接收到任务之后，他可以去摸鱼，可以吃饭，甚至环游世界，只要他完成了代理任务，那他的职责就算完成，因此我们可手动在代理人 `DoSomething()` 前后去定义代理人自身的一些行为 `BeforeDoSomething()` 和 `AfterDoSomething()`