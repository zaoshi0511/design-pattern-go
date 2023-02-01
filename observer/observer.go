package observer

type (
	Observer struct {
	}

	ObservedSubject struct {
		Observers []*Observer
		Value     interface{}
	}
)

func (s *ObservedSubject) AddObserver(obs *Observer) {
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
