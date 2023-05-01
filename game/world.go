package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	mplusNormalFont font.Face
	tilesImage      *ebiten.Image
)

// type WorldInterface interface {
// 	startingPositions() [][]int
// 	drawScreen(screen *ebiten.Image)
// 	init()
// }

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
	x float64
	y float64
}

func isStartingPosition(c int, r int) bool {
	if 36 > c && 4 < c && r > 4 && r < 26 {
		return true
	} else {
		return false
	}
}

func (w *World) init() {
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

	rand.Seed(time.Now().UnixNano()) // deprecated
	w.frequency = 1000
	w.tileSize = 20
	tilesImage = ebiten.NewImage(w.tileSize, w.tileSize)
	w.bgColor = color.White
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

func (w *World) startingPositions() [][]int {
	// will return coordinates for the starting positions

	return [][]int{}
}

func (w *World) scoreboard(screen *ebiten.Image, pos int, player PlayerInterface) {
	// Draw player name
	vector.DrawFilledRect(screen, 800, float32((600/4)*pos-(600/4-5)), 200, 32, player.GetColor(), false)
	text.Draw(screen, player.GetName(), mplusNormalFont, 800+10, int((600/4)*pos-((600/4)-30)), color.Black)

	// Draw lines to create scoreboard
	if pos == 1 {
		vector.DrawFilledRect(screen, 800, 0, 200, 5, color.Black, false)
	}
	vector.DrawFilledRect(screen, 800, float32((600/4)*pos), 200, 5, color.Black, false)

	// Draw Score
	text.Draw(screen, fmt.Sprintf("Score: %d", pos), mplusNormalFont, 800+10, int((600/4)*pos-((600/4)-60)), color.Black)

	// Draw Health Bar
	text.Draw(screen, fmt.Sprintf("Health: %d/3", pos), mplusNormalFont, 800+10, int((600/4)*pos-((600/4)-90)), color.Black)
	vector.DrawFilledRect(screen, 800+10, float32((600/4)*pos-((600/4)-110)), 180, 30, color.RGBA{0, 0xff, 0, 0xff}, false)
	vector.DrawFilledRect(screen, 800+10, float32((600/4)*pos-((600/4)-110)), 90, 30, color.RGBA{0xff, 0, 0, 0xff}, false)
}

func (w *World) drawScreen(screen *ebiten.Image, players []PlayerInterface) {
	screen.Fill(w.bgColor)

	for ci, _ := range w.tilemap {
		for _, r := range w.tilemap[ci] {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(r.x, r.y)
			vector.DrawFilledRect(screen, float32(r.x), float32(r.y), float32(w.tileSize), float32(w.tileSize), r.color, false)
		}
	}

	for i, player := range players {
		fmt.Println(player.GetName())
		w.scoreboard(screen, i+1, player)
	}
}
