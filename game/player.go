package main

import (
	"image/color"

	"codegame.com/codegame/tank"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	name  string
	tank  tank.Tank
	color color.Color
}

type PlayerInterface interface {
	Init(float64, float64)
	Update()
	Draw(*ebiten.Image)
	GetName() string
	GetColor() color.Color
}

// represents the specific player
// type Player struct {
// 	health float64
// 	score  int
// 	unit   Unit // Anything that implements a Unit can be here
// }

// func (p *Player) init(pos Position) {
// 	// set the initial position for the unit
// 	p.unit.setPosition(pos.x, pos.y)

// 	// set the init stats window
// }
