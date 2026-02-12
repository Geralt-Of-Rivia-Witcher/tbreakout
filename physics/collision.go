package physics

import "breakout/entities"

func DetectWallCollision(screenWidth int, ball *entities.Ball) {
	if ball.X <= 1 || ball.X >= screenWidth-2 {
		ball.Dx = -ball.Dx
	}
	if ball.Y <= 0 {
		ball.Dy = -ball.Dy
	}
}

func DetectPaddleCollisionAndCheckIfAlive(screenHeight int, ball *entities.Ball, paddle *entities.Paddle) bool {
	if ball.Y == screenHeight-1 {
		paddleStart := paddle.X - paddle.Width/2
		paddleEnd := paddle.X + paddle.Width/2
		if paddleStart <= ball.X && ball.X <= paddleEnd {
			ball.Dy = -ball.Dy
			midOfMidOfPaddle := ((paddle.Width / 2) / 2) + 1

			if ball.X == paddle.X {
				ball.Dx = 0
			} else if ball.X < paddle.X {
				leftMidOfPaddle := paddle.X - midOfMidOfPaddle - 1
				if ball.X <= leftMidOfPaddle {
					ball.Dx = -(ball.BallSpeed + 1)
				} else {
					ball.Dx = -ball.BallSpeed
				}
			} else {
				rightMidOfMiddle := paddle.X + midOfMidOfPaddle + 1
				if ball.X >= rightMidOfMiddle {
					ball.Dx = ball.BallSpeed + 1
				} else {
					ball.Dx = ball.BallSpeed
				}
			}
		}
	}
	if ball.Y >= screenHeight {
		return false
	}
	return true
}

func DetectBrickCollision(ball *entities.Ball, bricks []*entities.Brick) {
	ballX := ball.X
	ballY := ball.Y
	for _, brick := range bricks {
		if brick.Alive {
			brickStartX := brick.X - (brick.Width / 2)
			brickEndX := brick.X + (brick.Width / 2)
			if ballY+ball.Dy == brick.Y {
				if brickStartX-1 <= ballX && ballX <= brickEndX+1 {
					brick.Alive = false
					ball.Dy = -ball.Dy
					break
				}
			}
		}
	}
}
