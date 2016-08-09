package server


type Message struct {
    ID int
	Params []string `json:"params"`
}

// type Message struct {
// 	ObjectID int `json:"objectID"`
// 	MethodID int `json:"methodID"`
// 	Arguments []string `json:"arguments"`
// }

// type Event struct {
// 	ObjectID int `json:"objectID"`
// 	EventID int `json:"eventID"`
// 	Params []string `json:"params"`
// }

// func (self *Message) String() string {
// 	return self.Author + " says " + self.Body
// }
