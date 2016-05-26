package ragtime

type SceneManager struct {
    scene *Scene
    SetSceneCh chan *Scene
}

func NewSceneManager() *SceneManager {
    return &SceneManager{nil, make(chan *Scene)}
}

func (sm *SceneManager) Start() {
    go sm.listen()
}

func (sm *SceneManager) listen() {
	for {
		select {

		case s := <-sm.SetSceneCh:
		    sm.scene = s
        }
	}
}