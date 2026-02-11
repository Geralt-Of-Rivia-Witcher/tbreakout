package entities

type Paddle struct {
	X     int
	Width int
	Speed int
}

func NewPaddle(screenWidth int, paddleWidth int, speed int) *Paddle {
	if paddleWidth%2 == 0 {
		panic("Paddle width cannot be an even number")
	}
	return &Paddle{
		X:     screenWidth / 2,
		Width: paddleWidth,
		Speed: speed,
	}
}

func (paddle *Paddle) Move(direction int, screenWidth int) {
	paddle.X += (paddle.Speed * direction)
	if direction < 0 {
		leftEdge := paddle.X - (paddle.Width / 2)
		if leftEdge <= 0 {
			paddle.X = paddle.Width / 2
		}
	} else if direction > 0 {
		rightEdge := paddle.X + (paddle.Width / 2)
		if rightEdge >= screenWidth {
			paddle.X = screenWidth - (paddle.Width / 2) - 1
		}
	}

}
