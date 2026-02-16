package entities

import "breakout/constants"

const BallSpeed = 1

type Ball struct {
	X         int
	Y         int
	Dx        int
	Dy        int
	BallSpeed int
}

func NewBall(screenWidth int, screenHeight int) *Ball {
	return &Ball{
		X:         screenWidth / 2,
		Y:         screenHeight - constants.BottomBorderHeight - 1,
		Dx:        0,
		Dy:        -BallSpeed,
		BallSpeed: 1,
	}
}

func (ball *Ball) Move() {
	ball.X += ball.Dx
	ball.Y += ball.Dy
}

func (ball *Ball) ResetBall(screenWidth int, screenHeight int) {
	ball.X = screenWidth / 2
	ball.Y = screenHeight - constants.BottomBorderHeight - 1
	ball.Dx = 0
	ball.Dy = -BallSpeed
	ball.BallSpeed = 1
}
