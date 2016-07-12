package component

import (
    "log"
    )


type Position struct {
    position [2]float32
}

func (p *Position) IsDeleted() bool {
    return false
}

func (p *Position) Update() {
    log.Println("pos update")
}

func (p *Position) UpdateOrder() int {
    return 0
}