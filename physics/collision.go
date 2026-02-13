package physics

import "breakout/entities"

func DetectWallCollision(screenWidth int, ball *entities.Ball) {
	if ball.X <= 4 || ball.X >= screenWidth-3 {
		ball.Dx = -ball.Dx
	}
	if ball.Y <= 6 {
		ball.Dy = -ball.Dy
	}
}

func DetectPaddleCollisionAndCheckIfAlive(screenHeight int, ball *entities.Ball, paddle *entities.Paddle) bool {
	ballNextY := ball.Y + ball.Dy
	if ballNextY == paddle.Y {
		paddleStart := paddle.X - paddle.Width/2 - 1
		paddleEnd := paddle.X + paddle.Width/2 + 1
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

func DetectBrickCollisionAndGetScoreGained(ball *entities.Ball, bricks []*entities.Brick) int {
	ballX := ball.X
	ballY := ball.Y

	score := 0
	hit := false
	for _, brick := range bricks {
		if !brick.Alive {
			continue
		}
		brickStartX := brick.X - (brick.Width / 2)
		brickEndX := brick.X + (brick.Width / 2)
		if ballY+ball.Dy == brick.Y {
			if brickStartX-1 <= ballX && ballX <= brickEndX+1 {
				brick.Alive = false
				score += 100
				hit = true
			}
		}
	}
	if hit {
		ball.Dy = -ball.Dy
	}
	return score
}
