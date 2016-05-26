package ragtime


type Message struct {
	From string `json:"from"`
	To   string `json:"to"`
	Command string `json:"command"`
	Arguments []string `json:"arguments"`
}

// func (self *Message) String() string {
// 	return self.Author + " says " + self.Body
// }
