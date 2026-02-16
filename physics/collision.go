package physics

import (
	"github.com/Geralt-Of-Rivia-Witcher/tbreakout/constants"
	"github.com/Geralt-Of-Rivia-Witcher/tbreakout/entities"
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
		ballNextX := ball.X + ball.Dx

		if paddleStart <= ballNextX && ballNextX <= paddleEnd {
			ball.Dy = -ball.Dy
			midOfMidOfPaddle := ((paddle.Width / 2) / 2) + 1

			if ballNextX == paddle.X {
				ball.Dx = 0
			} else if ballNextX < paddle.X {
				leftMidOfPaddle := paddle.X - midOfMidOfPaddle - 1
				if ballNextX <= leftMidOfPaddle {
					ball.Dx = -(ball.BallSpeed + 1)
				} else {
					ball.Dx = -ball.BallSpeed
				}
			} else {
				rightMidOfMiddle := paddle.X + midOfMidOfPaddle + 1
				if ballNextX >= rightMidOfMiddle {
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
		nextBallY := ball.Y + ball.Dy
		if nextBallY == brick.Y || ball.Y == brick.Y { // need to test this part more.
			nextBallX := ball.X + ball.Dx
			if brickStartX-2 <= nextBallX && nextBallX <= brickEndX+2 {
				brick.Alive = false
				score += 100
				hit = true
				// break
			}
		}
	}
	if hit {
		ball.Dy = -ball.Dy
	}
	return score, remainingBricks
}
