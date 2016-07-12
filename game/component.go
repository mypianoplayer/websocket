package game

type Component interface {
	UpdateOrder() int
	Update()
}

