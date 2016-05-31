package ragtime

type GameComponent struct {
    gameObject *GameObject
    componentType int
}

func (c *GameComponent) Update() {

}

func (c *GameComponent) GetType() int {
    return c.componentType
}

func (c *GameComponent) SetGameObject(o *GameObject) {
    c.gameObject = o
}