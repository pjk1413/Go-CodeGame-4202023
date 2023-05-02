package main

import (
	"fmt"
	"image/color"
	"log"

	"codegame.com/codegame/player"
	"codegame.com/codegame/players"
	"codegame.com/codegame/tank"
	"github.com/hajimehoshi/ebiten/v2"
)

// https://blog.carlmjohnson.net/post/2021/how-to-use-go-embed/ - possible use for getting files

var (
	red    = color.RGBA{0xff, 0, 0, 0xff}
	blue   = color.RGBA{0, 0, 0xff, 0xff}
	green  = color.RGBA{0, 0xff, 0, 0xff}
	yellow = color.RGBA{0xff, 0xff, 0, 0xff}
	orange = color.RGBA{0xff, 0xA5, 0, 0xff}
)

const (
	title  string  = "Tanks"
	speed  float64 = 5
	width  int     = 1000
	height int     = 600
	rows   int     = 6
	cols   int     = 5
)

type P struct {
	t tank.Tank
	i player.PlayerInterface
}

type Game struct {
	world   World // set world here
	players []player.PlayerInterface
	tanks   []tank.Tank
	ps      map[string]map[string]P
	x       float64
	y       float64
}

func (g *Game) Update() error {
	// // Basic bot movements
	for i := 0; i < len(g.tanks); i++ {
		var hit, move, score bool = false, true, false

		pos := g.tanks[i].GetPosition()
		// fmt.Println(g.tanks[i].Name, g.tanks[i].GetPosition().X)

		// Detect for collisions with edges
		if OffMap(pos.X, pos.Y, float64(pos.Width), float64(pos.Height), float64(width-200), float64(height)) {
			hit, move = true, false
			fmt.Println("wall collision")
		}

		// check for bullet collisions and tank collisions
		for t := 0; t < len(g.tanks); t++ {
			if t != i {
				enemy := g.tanks[t].GetPosition()

				// Check for collision with another tank
				if Collision(pos, enemy) {
					hit, move = true, false
					// fmt.Println("tank collision", pos.X, enemy.X)
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
			g.players[i].Update(&g.tanks[i])
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
	// Order matters

	// draw the world onto the screen
	g.world.drawScreen(screen, g.tanks)
	// text.Draw(screen, fmt.Sprintf("POSX: %d", g.x), mplusNormalFont, int(g.x), 25, color.Black)
	// text.Draw(screen, tank.Name, mplusNormalFont, 800+10, int((600/4)*pos-((600/4)-30)), color.Black)
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

	// assign tank and players here
	// g.tanks = map[string]map[string]P

	g.tanks = []tank.Tank{tank.Tank{Name: "Sample Tank", Color: red}, tank.Tank{Name: "Destroyer", Color: blue}} // players will need to init better
	g.players = []player.PlayerInterface{&players.TemplateTanks{}, &players.Destroyer{}}

	// Need to initialize these better
	g.tanks[0].Position(100, 100)
	g.tanks[1].Position(500, 500)

	// Set window size and title
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle(title)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
