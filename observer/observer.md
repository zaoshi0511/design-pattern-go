# Observer Pattern 观察者模式

---

```go
package observer

type (
	Observer struct {
	}

	ObservedSubject struct {
		Observers []*Observer
		Value     interface{}
	}
)

func (s *ObservedSubject) RegisterObserver(obs *Observer) {
	s.Observers = append(s.Observers, obs)
}

func (s *ObservedSubject) RemoveObserver(obs *Observer) {
	var idx int
	for k, v := range s.Observers {
		if v == obs {
			idx = k
			break
		}
	}

	s.Observers = append(s.Observers[:idx], s.Observers[idx+1])
}

func (s *ObservedSubject) Modify(newValue interface{}) {
	s.Value = newValue
	s.notify()
}

func (s *ObservedSubject) notify() {
	for _, v := range s.Observers {
		v.DoSomethingAfterNotified()
	}
}

func (s *Observer) DoSomethingAfterNotified() {}
```

在观察者模式中，核心思想是，当被观察的对象执行某些特殊操作时，要通知与其关联的观察者          

在上述的代码示例中，都使用了具体的 `struct` 来定义观察者和被观察对象，也可将其抽象成接口

```go
package observer

type (
	Observer interface {
		DoSomethingAfterNotified()
    }
	
	ObservedSubject interface {
		RegisterObserver(obs Observer)
		RemoveObserver(obs Observer)
		Notify()
    }
)
```

对于被观察对象，要能够添加或者移除与自身关联的观察者，同时也要能通知与自身关联的所有观察者   

对于观察者，在接收到被观察对象的通知后，可以不做反应，也可以根据通知的具体细节做出不同的反应