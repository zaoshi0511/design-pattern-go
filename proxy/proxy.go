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
