package players

import (
	"image/color"

	"codegame.com/codegame/tank"
	"github.com/hajimehoshi/ebiten/v2"
)

// Tank name goes here
const name = "Sample"

// Tank color goes here
var tankColor = color.RGBA{0xff, 0, 0, 0xff}

type Sample struct { // create a struct with a name and tank field with a unqiue name to identify your bot
	name  string
	tank  tank.Tank
	color color.Color
}

func (s *Sample) Update() { // this is the main update function, it is called every tick, the tank within your struct will contain all methods and data you have acces to
	s.tank.Accelerate()
	s.tank.FireBullet()

	s.tank.Update() // this call makes changes to all positions of your items on the map (KEEP THIS HERE)
}

func (s *Sample) Init(x, y float64) {
	s.name = name                     // assign a name that will be displayed on the screen (this should match your struct name closely ideally)
	s.color = tankColor               // assign a color
	t := tank.Create(x, y, tankColor) // create a tank that can be referenced throughout the game
	s.tank = t
}

func (s *Sample) Draw(screen *ebiten.Image) {
	s.tank.Draw(screen)
}

func (s *Sample) GetName() string {
	return s.name
}

func (s *Sample) GetColor() color.Color {
	return s.color
}
