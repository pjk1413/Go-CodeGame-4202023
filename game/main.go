package main

import (
	"log"

	"codegame.com/codegame/player"
	"codegame.com/codegame/players"
	"codegame.com/codegame/tank"
	"github.com/hajimehoshi/ebiten/v2"
)

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
	players []player.PlayerInterface
	tanks   []tank.Tank
	x       float64
	y       float64
}

func (g *Game) Update() error {
	// // Basic bot movements
	for i := 0; i < len(g.tanks); i++ {
		var hit, move, score bool = false, true, false

		// get tank position
		pos := g.tanks[i].GetPosition()

		// Detect for collisions with edges
		if OffMap(pos, float64(width-200), float64(height)) {
			hit, move = true, false
		}

		// check for collisions
		for t := 0; t < len(g.tanks); t++ {
			if t != i {
				enemy := g.tanks[t].GetPosition()

				// Check for collision with another tank
				if Collision(pos, enemy) {
					hit, move = true, false
				}

				// check for collision with bullet
				if Hit(pos, enemy) {
					hit, move = true, false
				}

				// check if your bullet hit an enemy
				if Score(pos, enemy) {
					score = true
				}
			}
		}

		if move {
			// create an array of enemy positions to pass into Update function
			tp := []tank.TankPosition{}

			// get all the enemy tanks positions
			for t := 0; t < len(g.tanks); t++ {
				if t != i {
					tp = append(tp, g.tanks[t].GetPosition())
				}
			}

			// pass positions to Update as well as tank
			g.players[i].Update(&g.tanks[i], tp)
		}

		if hit {
			g.tanks[i].Hit()
		}

		if score {
			g.tanks[i].Score()
		}

	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Order matters here

	// draw the world onto the screen
	g.world.drawScreen(screen, g.tanks)

	// draw the tanks
	for i := 0; i < len(g.players); i++ {
		g.tanks[i].Draw(screen)
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

	// init tanks and players, for now all players will need to be hard coded into the game
	g.players = []player.PlayerInterface{
		&players.SampleTank{},
		&players.Destroyer{},
		&players.Killer{},
		&players.BigBoy{},
	}

	// Create all players
	for _, player := range g.players {
		g.tanks = append(g.tanks, *player.Create())
	}

	// Set starting position for all players/tanks
	for i, position := range StartingPositions(4) {
		g.tanks[i].InitTank(position.x, position.y, position.facing)
	}

	// Set window size and title
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle(title)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
