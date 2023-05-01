package main

import (
	"log"

	"codegame.com/codegame/players"
	"codegame.com/codegame/tank"
	"github.com/hajimehoshi/ebiten/v2"
)

// https://blog.carlmjohnson.net/post/2021/how-to-use-go-embed/ - possible use for getting files

const (
	title  string  = "Tanks"
	speed  float64 = 5
	width  int     = 1000
	height int     = 600
	rows   int     = 6
	cols   int     = 5
)

type Game struct {
	world   World // set world here
	players []PlayerInterface
	tanks   []tank.Tank
	// controllers []controller.TankController
}

func (g *Game) Update() error {
	// User changes to movements

	// // Basic bot movements
	for i := 0; i < len(g.players); i++ {
		g.players[i].Update()
	}

	// Detect collisions and hits

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	// draw the world onto the screen
	g.world.drawScreen(screen, g.players)

	// draw the scoreboard onto the screen
	// g.world.drawScoreboard(screen)

	// Draw the tank and bullet actions
	for i := 0; i < len(g.players); i++ {

		g.players[i].Draw(screen)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return width, height
}

func main() {
	// & is the memory address
	g := &Game{}
	g.world = World{}
	g.world.init()

	// initialize the players
	g.players = []PlayerInterface{&players.Sample{}}
	for i, _ := range g.players {
		g.players[i].Init(50, 50)
	}

	// Initialize the array with preset values
	// for i, player := range players {
	// 	g.controllers[i] = create(player)
	// }

	// Set window size and title
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle(title)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
