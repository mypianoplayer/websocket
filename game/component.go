package game

type ComponentType int

type Component interface {
	ComponentType() ComponentType
	SetObject(o Object)
	Object() Object
	UpdateOrder() int
	Start()
	Update()
}

type ComponentBase struct {
	object Object
	componentType ComponentType
	updateOrder int
}

func NewComponentBase(componentType ComponentType,updateOrder int) *ComponentBase {
	return &ComponentBase{
		object:nil,
		componentType:componentType,
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

func (c *ComponentBase) ComponentType() ComponentType {
	return c.componentType
}