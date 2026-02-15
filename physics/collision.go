package physics

import (
	"breakout/constants"
	"breakout/entities"
)

func DetectWallCollision(screenWidth int, ball *entities.Ball) {
	nextBallX := ball.X + ball.Dx
	if nextBallX <= constants.BorderWidth+2 || nextBallX >= screenWidth-constants.BorderWidth-2 {
		ball.Dx = -ball.Dx
	}
	if ball.Y <= constants.TopHUDElementHeight+1 {
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

func DetectBrickCollisionAndGetScoreGainedAndRemainingBricks(ball *entities.Ball, bricks []*entities.Brick) (int, int) {
	ballX := ball.X
	ballY := ball.Y

	score := 0
	hit := false
	remainingBricks := 0
	for _, brick := range bricks {
		if !brick.Alive {
			continue
		}
		remainingBricks++
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
	return score, remainingBricks
}
