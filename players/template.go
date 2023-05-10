package players

import (
	"codegame.com/codegame/colors"
	"codegame.com/codegame/tank"
)

type SampleTank struct {}

// Will need the tilemap and the player locations
func (sampleTank *SampleTank) Update(t *tank.Tank, pos []tank.TankPosition) {
	// for _, position := range pos {
		// loop through the positions of all enemy tanks
	// }

	// All the actions that can be performed
	t.GetPosition()
	t.Accelerate()
	t.Decelerate()
	t.RotateLeft()
	t.RotateRight()
	t.Fire()
	t.Stop()

	t.Update() // This always needs to be called within the Update method
}

func (tt *SampleTank) Create() *tank.Tank {
	// Create a new tank
	return &tank.Tank{
		Name:  "SampleTank",
		Color: colors.RED,
	}
}
