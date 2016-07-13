package component

import (
    "log"
    )


type Position struct {
    pos [2]float32
}

func (p *Position) IsDeleted() bool {
    return false
}

func (p *Position) Update() {
    log.Println("pos update")

    // server.SendAll pos
}

func (p *Position) UpdateOrder() int {
    return 0
}