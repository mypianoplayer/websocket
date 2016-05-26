package ragtime

type Scene struct {
    gameObjects []*GameObject;
    gameComponents map[int][]*GameComponent
    //gameComponents  component typeをキーとしたマップ
}

func (s *Scene) Update() {

// 	for _, c := range s.connections {
// 		c.SendMessageCh <- msg
// 	}

}