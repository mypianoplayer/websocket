package main

import (
    "ragtime/game"
    )

type Input struct {
    game.Object
    clicked bool
    pos [2]float32
}


