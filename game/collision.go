package main

import (
	"codegame.com/codegame/tank"
)

func OffMap(pos tank.TankPosition, worldX, worldY float64) bool {

	// x axis
	if pos.X < 0 || pos.X+float64(pos.Width) > worldX {
		return true
	}

	if pos.Y < 0 || pos.Y+float64(pos.Height) > worldY {
		return true
	}

	return false
}

func valueInRange(value, min, max float64) bool {
	return (value >= min) && (value <= max)
}

func rectOverlap(A, B tank.TankPosition) bool {
	xOverlap := valueInRange(A.X, B.X, B.X+float64(B.Width)) ||
		valueInRange(B.X, A.X, A.X+float64(A.Width))

	yOverlap := valueInRange(A.Y, B.Y, B.Y+float64(B.Height)) ||
		valueInRange(B.Y, A.Y, A.Y+float64(A.Height))

	return xOverlap && yOverlap
}

func Collision(pos, enemy tank.TankPosition) bool {
	return rectOverlap(pos, enemy)
}

func bulletOverlap(A, B tank.TankPosition) bool {

	xOverlap := valueInRange(A.BulletX, B.X, B.X+float64(B.Width)) ||
		valueInRange(B.X, A.BulletX, A.BulletX+float64(A.BulletSize))

	yOverlap := valueInRange(A.BulletY, B.Y, B.Y+float64(B.Height)) ||
		valueInRange(B.Y, A.BulletY, A.BulletY+float64(A.BulletSize))

	return xOverlap && yOverlap
}

func Score(pos, enemy tank.TankPosition) bool {
	return bulletOverlap(pos, enemy)
}

func Hit(pos, enemy tank.TankPosition) bool {
	return bulletOverlap(enemy, pos)
}
