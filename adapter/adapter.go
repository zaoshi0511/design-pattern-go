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
