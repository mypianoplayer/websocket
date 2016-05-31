package ragtime

type GameObject struct {
    objectName string
    gameComponents map[int]*GameComponent
}

func (o *GameObject) Start() {

}

func (o *GameObject) AddComponent(c *GameComponent) {
    o.gameComponents[c.GetType()] = c
    c.SetGameObject(o)
}

func (o *GameObject) GetComponent(componentType int) *GameComponent {
    return o.gameComponents[componentType]
}