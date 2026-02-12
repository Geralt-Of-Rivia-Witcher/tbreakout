package entities

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
		Y:         screenHeight - 2,
		Dx:        0,
		Dy:        -BallSpeed,
		BallSpeed: 1,
	}
}

func (ball *Ball) Move() {
	ball.X += ball.Dx
	ball.Y += ball.Dy
}
