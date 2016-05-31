package main

import (
    "ragtime"
)

type PositionComponent struct {
    ragtime.GameComponent
    position []float
}

func NewPositionComponent() *PositionComponent
{
    return &PositionComponent {
        ragtime.GameComponent{ ComponentType_Position },
        0.0,
        0.0
    }
}