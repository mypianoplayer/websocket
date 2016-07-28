package game

type Component interface {
	SetObject(o Object)
	UpdateOrder() int
	Start()
	Update()
}

type ComponentBase struct {
	object Object
}

func NewComponent() *ComponentBase {
	return &ComponentBase{nil}
}

func (c *ComponentBase) SetObject(o Object) {
	c.object = o
}
