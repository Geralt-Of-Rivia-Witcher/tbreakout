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
				if brickStartX <= ballX && ballX <= brickEndX {
					brick.Alive = false
					ball.Dy = -ball.Dy
					break
				}
			}
		}
	}
}
