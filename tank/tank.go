package tank

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	red    = color.RGBA{0xff, 0, 0, 0xff}
	blue   = color.RGBA{0, 0, 0xff, 0xff}
	green  = color.RGBA{0, 0xff, 0, 0xff}
	yellow = color.RGBA{0xff, 0xff, 0, 0xff}
	orange = color.RGBA{0xff, 0xA5, 0, 0xff}
)

type TankPosition struct {
	Width      int
	Height     int
	X          float64
	Y          float64
	BulletX    float64
	BulletY    float64
	BulletSize int
	Facing     float64
}

type Tank struct {
	width    int
	height   int
	speed    float64
	maxSpeed int

	x float64
	y float64

	fire        bool
	bulletSize  int
	bulletx     float64
	bullety     float64
	bulletSpeed float64

	facing    float64
	setFacing float64

	hit         bool
	hitCooldown float64

	radarx float64
	radary float64
	band   float64

	respawn         bool
	respawnCooldown float64
	startingPosX    float64
	startingPosY    float64
	health          int
	score           int
	Name            string
	Color           color.Color
}

func (t *Tank) GetPosition() TankPosition {
	return TankPosition{
		X:          t.x - (float64(t.width) / 2),
		Y:          t.y - (float64(t.height) / 2),
		Width:      t.width,
		Height:     t.height,
		BulletX:    t.bulletx - (float64(t.bulletSize) / 2),
		BulletY:    t.bulletx - (float64(t.bulletSize) / 2),
		BulletSize: t.bulletSize,
		Facing:     t.facing,
	}
}

func InitTank(x, y, facing float64, c color.Color, name string) *Tank {
	t := &Tank{}
	t.x = x
	t.y = y
	t.Name = name
	t.Color = red

	t.startingPosX = x
	t.startingPosY = y

	t.height = 50
	t.width = 50
	t.speed = 0
	t.maxSpeed = 3

	t.fire = false
	t.bulletSize = 5
	t.bulletSpeed = 5
	t.facing = 0
	return t
}

// TODO: remove this from public
func (t *Tank) Position(x, y float64) {
	t.x = x
	t.y = y

	t.startingPosX = x
	t.startingPosY = y

	t.height = 50
	t.width = 50
	t.speed = 0
	t.maxSpeed = 3

	t.fire = false
	t.bulletx = 0
	t.bullety = 0
	t.bulletSize = 5
	t.bulletSpeed = 5
	t.Color = red
	t.facing = 0
	t.setFacing = 0
}

func (t *Tank) SetFacing(face float64) {
	t.setFacing = face
}

func (t *Tank) Fire() {
	if t.bulletx == 0 && t.bullety == 0 {
		t.bulletx = t.x
		t.bullety = t.y
	}
	t.fire = true
}

func (t *Tank) Accelerate() {
	// increase tank speed
	if t.speed < float64(t.maxSpeed) {
		t.speed += .01
	}
}

func (t *Tank) Decelerate() {
	// decrease tank speed
	if t.speed > float64(-t.maxSpeed+1) {
		t.speed -= .01
	}
}

func (t *Tank) RotateLeft() {
	if t.setFacing != 0 {
		if t.facing > t.setFacing && t.facing-.1 < t.setFacing {
			t.facing = t.setFacing
		} else {
			t.facing -= .1
		}
	} else {
		t.facing -= .1
	}

}

func (t *Tank) RotateRight() {
	if t.setFacing != 0 {
		if t.facing < t.setFacing && t.facing+.1 > t.setFacing {
			t.facing = t.setFacing
		} else {
			t.facing += .1
		}
	} else {
		t.facing += .1
	}
}

func (t *Tank) Update() {
	// updates movements for the tank
	t.x = t.x + t.speed*math.Cos(t.facing)
	t.y = t.y + t.speed*math.Sin(t.facing)

	// update movements for bullet
	if t.fire {
		t.bulletx += t.bulletSpeed * math.Cos(t.facing)
		t.bullety += t.bulletSpeed * math.Sin(t.facing)

		if t.bulletx > 800 || t.bulletx < 0 || t.bullety < 0 || t.bullety > 600 {
			t.fire = false
			t.bulletx = 0
			t.bullety = 0
		}
	}

	// rotate radar
	t.band += .05
}

func (t *Tank) Draw(screen *ebiten.Image) {
	// Draw all movements related to the tank

	if t.respawn && t.respawnCooldown >= 100 {
		t.respawnCooldown += 1

		if t.respawnCooldown == 150 {
			t.respawn = false
			t.respawnCooldown = 0
			t.speed = 0
			t.x = t.startingPosX
			t.y = t.startingPosY
		}
	} else {
		// TODO: check if a single cooldown timer can be used
		if t.respawn {
			t.respawnCooldown += 1
		}
		// draw tank
		rect := ebiten.NewImage(t.width, t.height)
		s := rect.Bounds().Size()
		rect.Fill(t.Color)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(-float64(s.X)/2, -float64(s.Y)/2) // helps with accurate rotation
		op.GeoM.Rotate(t.facing)                            // rotate the tank
		op.GeoM.Translate(t.x, t.y)                         // place the tank on the screen
		screen.DrawImage(rect, op)                          // draw

		// draw radar
		vector.StrokeLine(screen, float32(t.x), float32(t.y), float32(25*math.Cos(float64(t.band))+t.x), float32(25*math.Sin(float64(t.band))+t.y), 2, color.Black, false)

		// draw bullet
		if t.fire {
			fmt.Println("FIRE")

			bullet := ebiten.NewImage(5, 5)

			bullet.Fill(color.Black)

			ops := &ebiten.DrawImageOptions{}
			ops.GeoM.Translate(t.bulletx, t.bullety)
			screen.DrawImage(bullet, ops)
		}

		// draw hit to tank
		if t.hit {
			// equation to make explosion effect
			d := -0.01*(t.hitCooldown*t.hitCooldown) + (1.1 * t.hitCooldown) + 1

			if d > 0 {
				// make explosion grow and fall
				explosion := ebiten.NewImage(int(d), int(d))
				se := explosion.Bounds().Size()
				explosion.Fill(color.RGBA{0xcd, 0xb3, 0x5d, 0xff})

				oph := &ebiten.DrawImageOptions{}

				oph.GeoM.Translate(-float64(se.X)/2, -float64(se.Y)/2)
				oph.GeoM.Translate(t.x, t.y)
				screen.DrawImage(explosion, oph)
				t.hitCooldown += 1
			} else {
				t.hit = false
				t.hitCooldown = 0
			}
		}

	}

}

// PRIVATE METHODS BELOW THIS LINE

func (t *Tank) Hit() {
	if !t.respawn {
		t.hit = true // triggers the drawing only
		t.speed = 0
		t.health = t.health - 1
		t.hitCooldown = 0
		t.respawn = true
		t.respawnCooldown = 0
		t.facing = 0
	}
}

func (t *Tank) Score() {
	t.score += 1
	t.fire = false
	t.bulletx = 0
	t.bullety = 0
}

func (t *Tank) Stop() {
	t.speed = 0
}

func (t *Tank) reset() {

	t.x = t.startingPosX
	t.y = t.startingPosY
	t.respawn = false

}

func (t *Tank) detect(players []Tank) [][]int {
	// returns a map of players locations that are within radar range
	return [][]int{}
}

func (t *Tank) collision() {
	// determine if the tank has collided with a wall or another player
}
