package entities

type Paddle struct {
	X          int
	leftLimit  int
	rightLimit int
	Width      int
}

func NewPaddle(screenWidth int, paddleWidth int) *Paddle {
	if paddleWidth%2 == 0 {
		panic("Paddle width cannot be an even number")
	}
	return &Paddle{
		X:          screenWidth / 2,
		leftLimit:  0,
		rightLimit: screenWidth,
		Width:      paddleWidth,
	}
}

func (paddle *Paddle) Move(dx int) {
	paddle.X += dx
	if dx < 0 {
		leftEdge := paddle.X - (paddle.Width / 2)
		if leftEdge <= 0 {
			paddle.X = paddle.Width / 2
		}
	} else if dx > 0 {
		rightEdge := paddle.X + (paddle.Width / 2)
		if rightEdge >= paddle.rightLimit {
			paddle.X = paddle.rightLimit - (paddle.Width / 2) - 1
		}
	}

}
