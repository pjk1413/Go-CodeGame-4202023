package player

import (
	"codegame.com/codegame/tank"
)

type Player struct{}

type PlayerInterface interface {
	Update(*tank.Tank, []tank.TankPosition)
	Create() *tank.Tank
}
