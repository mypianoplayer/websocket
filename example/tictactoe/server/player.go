package main

import (
    "log"
    // "reflect"
    "ragtime/game"
    "ragtime/example/tictactoe/server/component"
    )

type Player struct {
    game.Object
    component.Position
    component.View
}

func NewPlayer() *Player {
    return &Player {
        game.Object{},
        component.Position{},
        *component.NewView("{-]"),
    }
}


func (p *Player) Test() {
    log.Println("TEST")
}


func (o *Player) EachComponent() chan game.Component {
    ch := make(chan game.Component)

    go func() {
        ch <- &o.Position
        close(ch)
    }()
    // v := reflect.ValueOf(*o)
    // n := v.NumField()
    // log.Println("numf", n)
    // i := 0
    // go func() {
    //     for {
    //         if i >= n {
    //             close(ch)
    //             break;
    //         }

    //         ch <- v.Field(i).Interface().(game.Component)
    //         i++
    //     }
    // }()

    return ch
}