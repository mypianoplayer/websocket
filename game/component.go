package game

type Component interface {
	SetObject(o Object)
	Object() Object
	UpdateOrder() int
	Start()
	Update()
}

type ComponentBase struct {
	object Object
	updateOrder int
}

func NewComponentBase(updateOrder int) *ComponentBase {
	return &ComponentBase{
		object:nil,
		updateOrder:updateOrder,
	}
}

func (c *ComponentBase) SetObject(o Object) {
	c.object = o
}

func (c *ComponentBase) Object() Object {
	return c.object
}

func (c *ComponentBase) UpdateOrder() int {
	return c.updateOrder
}