package tank

import (
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

	facing float64

	hit         bool
	hitCooldown float64

	radarx float64
	radary float64
	band   float64

	health int
	Color  color.Color
}

func Create(x, y float64, color color.Color) Tank {
	t := Tank{}
	t.x = x
	t.y = y

	t.height = 50
	t.width = 50
	t.speed = 0
	t.maxSpeed = 3

	t.bulletSize = 5
	t.bulletSpeed = 8
	t.Color = color
	t.facing = 0
	return t
}

func (t *Tank) FireBullet() {
	t.fire = true

	if t.bulletx == 0 && t.bullety == 0 {
		t.bulletx = t.x
		t.bullety = t.y
	}
}

func (t *Tank) moveBullet() {
	t.bulletx += t.bulletSpeed * math.Cos(t.facing)
	t.bullety += t.bulletSpeed * math.Sin(t.facing)

	if t.bulletx > 640 || t.bulletx < 0 {
		t.fire = false
		t.bulletx = 0
		t.bullety = 0
	}
}

func (t *Tank) isHit(x, y float64) {
	t.hit = true
	t.hitCooldown = 0
}

func (t *Tank) detect(players []Tank) [][]int {
	// returns a map of players locations that are within radar range
	return [][]int{}
}

func (t *Tank) Accelerate() {
	// fmt.Println("accelerate")
	// if t.speed < float64(t.maxSpeed) {
	// 	t.speed += 1
	// }
	t.speed += .01
	// t.move()
}

func (t *Tank) Decelerate() {
	if t.speed > float64(-t.maxSpeed+1) {
		t.speed -= .1
	}
}

func (t *Tank) breaking() {
	if t.speed > 1 {
		t.speed -= 0.5
	} else if t.speed > 0.7 {
		t.speed -= 0.3
	} else if t.speed > 0.3 {
		t.speed = 0
	} else if t.speed < -1 {
		t.speed += 0.5
	} else if t.speed < -0.7 {
		t.speed += 0.3
	} else if t.speed < -0.3 {
		t.speed = 0
	}
}

func (t *Tank) rotateLeft() {
	t.facing -= .1
}

func (t *Tank) rotateRight() {
	t.facing += .1
}

func (t *Tank) rotateRadar() {
	t.band += .05
}

func (t *Tank) move() {
	t.x = t.x + t.speed*math.Cos(t.facing)
	t.y = t.y + t.speed*math.Sin(t.facing)
	t.rotateRadar()
}

func (t *Tank) Update() {
	t.move()
	t.moveBullet()
}

func (t *Tank) Draw(screen *ebiten.Image) {

	// draw tank
	rect := ebiten.NewImage(40, 40)
	s := rect.Bounds().Size()
	rect.Fill(t.Color)

	op := &ebiten.DrawImageOptions{}

	// tank movement
	// Move the image's center to the screen's upper-left corner.
	// This is a preparation for rotating. When geometry matrices are applied,
	// the origin point is the upper-left corner.
	op.GeoM.Translate(-float64(s.X)/2, -float64(s.Y)/2)

	// Rotate the image. As a result, the anchor point of this rotate is
	// the center of the image.
	// op.GeoM.Rotate(float64(t.facing%360) * 2 * math.Pi / 360)
	op.GeoM.Rotate(t.facing)
	op.GeoM.Translate(t.x, t.y)
	screen.DrawImage(rect, op)

	// draw radar
	vector.StrokeLine(screen, float32(t.x), float32(t.y), float32(25*math.Cos(float64(t.band))+t.x), float32(25*math.Sin(float64(t.band))+t.y), 2, color.Black, false)

	// draw bullet
	if t.fire {
		// t.moveBullet()

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

// creates a tank and returns it with default values attached
// func createTank(x, y float64) Tank {
// 	t := Tank{}
// 	t.x = x
// 	t.y = y

// 	t.height = 50
// 	t.width = 50
// 	t.speed = 0
// 	t.maxSpeed = 3

// 	t.bulletSize = 5
// 	t.bulletSpeed = 8

// 	t.facing = 0
// 	return t
// }
