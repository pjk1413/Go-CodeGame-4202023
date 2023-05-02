package players

import (
	"image/color"

	"codegame.com/codegame/tank"
)

type TemplateTanks struct {
	Name  string
	Color color.Color
}

// Will need the tilemap and the player locations
func (tt *TemplateTanks) Update(t *tank.Tank, pos []tank.TankPosition) {
	p := t.GetPosition()

	// circle := 2 * math.Pi

	if p.Facing > 89 && p.Facing < 91 {
		t.Fire()
	} else {
		t.RotateRight()
	}

	t.Update()
}

func (tt *TemplateTanks) Create() *tank.Tank {
	// Create a new tank
	return &tank.Tank{
		Name:  "TemplateTank",
		Color: color.RGBA{0, 0xff, 0xff, 0xff},
	}
}
