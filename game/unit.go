package main

import "github.com/hajimehoshi/ebiten/v2"

type UnitDefault struct {
	
}

// represents the character that moves around the screen
type Unit interface {
	setPosition(x float64, y float64) // set the initial position
	fire()                            // will fire projectile
	draw(screen *ebiten.Image)        // will draw all actions of the character
	accelerate()                      // will add one acceleration to the character
	deaccelerate()                    // will minus one acceleration from the character
	rotateRight()                     // will rotate the character right
	rotateLeft()                      // will rotate the character left
	breaking()                        // will apply breaks (not an immediate stop)
	move()                            // moves all elements associated with the character
	hit()                             // returns true if the character is hit
	detect()                          // returns list of all other character locations that are in range
	drawUnit() *ebiten.Image          // returns a new character image
}
