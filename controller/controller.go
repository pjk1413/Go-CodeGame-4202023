package controller

import "fmt"

type Controller interface {
	// Init()
	// Detect()
	// Update()
	Accelerate()
	Deaccelerate()
	RotateRight()
	RotateLeft()
	// Break()
	// GetPosition()
	// Fire()
	// GetFacing()
	// GetSpeed()
}

type TankController struct{}

func (t *TankController) Accelerate()   { fmt.Println("Accelerate") }
func (t *TankController) Deaccelerate() { fmt.Println("Deaccelerate") }
func (t *TankController) RotateRight()  {}
func (t *TankController) RotateLeft()   {}
