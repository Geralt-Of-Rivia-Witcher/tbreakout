package entities

import "github.com/Geralt-Of-Rivia-Witcher/tbreakout/constants"

type Paddle struct {
	X     int
	Y     int
	Width int
	Speed int
}

func NewPaddle(screenWidth int, screenHeight int, paddleWidth int, speed int) *Paddle {
	if paddleWidth%2 == 0 {
		panic("Paddle width cannot be an even number")
	}
	return &Paddle{
		X:     screenWidth / 2,
		Y:     screenHeight - constants.BottomBorderHeight - 1,
		Width: paddleWidth,
		Speed: speed,
	}
}

func (paddle *Paddle) ResetPaddle(screenWidth int) {
	paddle.X = screenWidth / 2
}

func (paddle *Paddle) Move(direction int, screenWidth int) {
	paddle.X += (paddle.Speed * direction)
	if direction < 0 {
		leftEdge := paddle.X - (paddle.Width / 2)
		if leftEdge < constants.BorderWidth+1 {
			paddle.X = constants.BorderWidth + paddle.Width/2 + 1
		}
	} else if direction > 0 {
		rightEdge := paddle.X + (paddle.Width / 2)
		if rightEdge > screenWidth-constants.BorderWidth-1 {
			paddle.X = screenWidth - constants.BorderWidth - 1 - paddle.Width/2
		}
	}

}
