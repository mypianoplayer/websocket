package ragtime

// import (
//     "log"
//     "time"
//     )

// type SceneManager struct {
//     scene Scene
//     SetSceneCh chan Scene
// }

// func NewSceneManager() *SceneManager {
//     return &SceneManager{nil, make(chan Scene)}
// }

// func (sm *SceneManager) Start() {
//     log.Println("Scene Manager Start")
//     go sm.listen()
// }

// func (sm *SceneManager) listen() {
// 	ticker := time.NewTicker(time.Second)///60.0)

// 	for {
// 		select {

//         case <-ticker.C:
//             if(sm.scene != nil) {
//                 sm.scene.Update()
//             }
// 		case s := <-sm.SetSceneCh:
// 		    sm.scene = s
//         }
// 	}
// }