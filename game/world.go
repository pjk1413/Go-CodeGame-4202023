package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"time"

	"codegame.com/codegame/tank"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	mplusNormalFont font.Face
	strikeface      font.Face
	tilesImage      *ebiten.Image
)

const (
	obstacles bool = false
)

type World struct {
	tileSize   int
	tilemap    [40][30]Block // hard coded size
	bgColor    color.Color
	blockColor color.Color
	frequency  int
}

type Block struct {
	x     float64
	y     float64
	color color.Color
	wall  bool
}

type Position struct {
	x      float64
	y      float64
	facing float64
}

func (w *World) toMap() [40][30]map[string]float64 {
	var tm [40][30]map[string]float64

	for ci, _ := range w.tilemap {
		for ri, r := range w.tilemap[ci] {
			tm[ci][ri] = map[string]float64{"x": r.x, "y": r.y, "size": float64(w.tileSize)}
		}
	}

	return tm
}

func isStartingPosition(c int, r int) bool {
	if 36 > c && 4 < c && r > 4 && r < 26 {
		return true
	} else {
		return false
	}
}

func (w *World) init() {
	// set background color
	w.bgColor = color.White

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})

	if err != nil {
		log.Fatal(err)
	}

	if obstacles {
		rand.Seed(time.Now().UnixNano()) // deprecated
		w.frequency = 1000
		w.tileSize = 20
		tilesImage = ebiten.NewImage(w.tileSize, w.tileSize)

		w.blockColor = color.Black

		// creates a tilemap
		var tilemap [40][30]Block
		for c := 0; c < 40; c++ {
			for r := 0; r < 30; r++ {
				block := Block{}
				block.color = color.White

				// create a wall around the playing field
				if c == 0 || r == 0 || c == len(tilemap)-1 || r == len(tilemap[0])-1 {
					block.color = w.blockColor
					block.wall = true

					// creates random obstacles
				} else if rand.Intn(w.frequency) < 10 && isStartingPosition(c, r) {
					block.color = w.blockColor
					block.wall = true
				}

				block.x = float64((c) * 20)
				block.y = float64((r) * 20)

				tilemap[c][r] = block
			}
		}

		w.tilemap = tilemap
	}

}

func StartingPositions(total int) []Position {
	// will return coordinates for the starting positions
	pos := []Position{
		Position{x: 50, y: 50, facing: 1},
		Position{x: 750, y: 50, facing: 3},
		Position{x: 50, y: 550, facing: 3},
		Position{x: 750, y: 550, facing: 3},
	}

	return pos[:total]
}

func (w *World) scoreboard(screen *ebiten.Image, pos int, tank tank.Tank) {

	// Draw player name
	vector.DrawFilledRect(screen, 800, float32((600/4)*pos-(600/4-5)), 200, 32, tank.GetColor(), false)

	name := tank.Name

	if tank.GetHealth() == 0 {
		name = "OUT *" + tank.Name
	}

	text.Draw(screen, name, mplusNormalFont, 800+10, int((600/4)*pos-((600/4)-30)), color.Black)

	// Draw lines to create scoreboard
	if pos == 1 {
		vector.DrawFilledRect(screen, 800, 0, 200, 5, color.Black, false)
	}
	vector.DrawFilledRect(screen, 800, float32((600/4)*pos), 200, 5, color.Black, false)

	// Draw Score
	text.Draw(screen, fmt.Sprintf("Score: %d", tank.GetScore()), mplusNormalFont, 800+10, int((600/4)*pos-((600/4)-60)), color.Black)

	// Draw Health Bar
	text.Draw(screen, fmt.Sprintf("Health: %d/3", tank.GetHealth()), mplusNormalFont, 800+10, int((600/4)*pos-((600/4)-90)), color.Black)
	vector.DrawFilledRect(screen, 800+10, float32((600/4)*pos-((600/4)-110)), 180, 30, color.RGBA{0xff, 0, 0, 0xff}, false)
	vector.DrawFilledRect(screen, 800+10, float32((600/4)*pos-((600/4)-110)), float32(tank.GetHealth())*180/3, 30, color.RGBA{0, 0xff, 0, 0xff}, false)

	// Draw scoreboard border
	vector.DrawFilledRect(screen, 800, 0, 5, 600, color.Black, false)
}

func (w *World) drawScreen(screen *ebiten.Image, tanks []tank.Tank) {
	screen.Fill(w.bgColor)

	// Draw obstacles to map
	if obstacles {
		for ci, _ := range w.tilemap {
			for _, r := range w.tilemap[ci] {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(r.x, r.y)
				vector.DrawFilledRect(screen, float32(r.x), float32(r.y), float32(w.tileSize), float32(w.tileSize), r.color, false)
			}
		}
	}

	for i, tank := range tanks {
		w.scoreboard(screen, i+1, tank)
	}
}
