package player

import (
	"image/color"

	"codegame.com/codegame/tank"
)

var (
	red    = color.RGBA{0xff, 0, 0, 0xff}
	blue   = color.RGBA{0, 0, 0xff, 0xff}
	green  = color.RGBA{0, 0xff, 0, 0xff}
	yellow = color.RGBA{0xff, 0xff, 0, 0xff}
	orange = color.RGBA{0xff, 0xA5, 0, 0xff}
)

type Player struct{}

type PlayerInterface interface {
	Update(*tank.Tank, []tank.TankPosition)
	Create() *tank.Tank
}
