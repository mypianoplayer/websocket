package component

import (
    "log"
    )

type View struct {
    text string
}

func NewView(t string) *View {
    return &View{t}
}

func (v *View) Update() {
    log.Println("view update")
}

func (v *View) UpdateOrder() int {
    return 1;
}