package ragtime


type Scene struct {
    gameObjects []*GameObject;
    gameComponents map[int][]*GameComponent
    //gameComponents  component typeをキーとしたマップ
}

func NewScene() *Scene {
    return &Scene{nil,nil};
}

func (s *Scene) Update() {

	for _, comps := range s.gameComponents {
	    for _, c := range comps {
    		c.Update()
	    }
	}

}

func (s *Scene) AddGameObject(o *GameObject) {

}