package render

import (
	"fmt"

	"github.com/Geralt-Of-Rivia-Witcher/tbreakout/constants"
	"github.com/Geralt-Of-Rivia-Witcher/tbreakout/entities"
	"github.com/gdamore/tcell/v3"
)

func RenderRunningGameScreen(
	screen tcell.Screen,
	screenWidth, screenHeight, lives, score int,
	paddle *entities.Paddle,
	bricks []*entities.Brick,
	ball *entities.Ball,
) {
	clear(screen)
	DrawHUD(lives, score, screenWidth, screenHeight, screen)
	DrawPaddle(paddle, screen)
	DrawBall(ball, screen)
	DrawBricks(bricks, screen)
	screen.Show()
}

func clear(screen tcell.Screen) {
	screen.SetStyle(baseStyle())
	screen.Clear()
}

func DrawPaddle(paddle *entities.Paddle, screen tcell.Screen) {
	style := paddleStyle()
	paddleLengthOnEachSideOfCenter := paddle.Width / 2
	for i := paddle.X - paddleLengthOnEachSideOfCenter; i <= paddle.X+paddleLengthOnEachSideOfCenter; i++ {
		screen.SetContent(i, paddle.Y, '▓', nil, style)
	}
}

func DrawHUD(lives int, score int, screenWidth int, screenHeight int, screen tcell.Screen) {
	border := borderStyle()
	label := hudLabelStyle()

	drawSpacer(screenWidth, constants.BorderWidth, screen)

	drawText(screen, screenWidth/3, 3, "LIVES: ", label)
	drawLives(screen, screenWidth/3+5, 3, lives)
	drawText(screen, screenWidth/3*2, 3, fmt.Sprintf("SCORE: %d", score), scoreStyle())

	drawSpacer(screenWidth, constants.TopHUDElementHeight, screen)
	drawBorders(screen, screenWidth, screenHeight, border)

	DrawInputHints(screenWidth, screenHeight, screen)
}

func drawBorders(screen tcell.Screen, screenWidth int, screenHeight int, style tcell.Style) {
	left := constants.BorderWidth
	right := screenWidth - constants.BorderWidth
	top := constants.BorderWidth
	hudBottom := constants.TopHUDElementHeight
	bottom := screenHeight - constants.BottomBorderHeight

	for x := left; x <= right; x++ {
		screen.SetContent(x, top, '═', nil, style)
		screen.SetContent(x, hudBottom, '═', nil, style)
		screen.SetContent(x, bottom, '═', nil, style)
	}
	for y := top; y <= bottom; y++ {
		screen.SetContent(left, y, '║', nil, style)
		screen.SetContent(right, y, '║', nil, style)
	}

	screen.SetContent(left, top, '╔', nil, style)
	screen.SetContent(right, top, '╗', nil, style)
	screen.SetContent(left, bottom, '╚', nil, style)
	screen.SetContent(right, bottom, '╝', nil, style)
	screen.SetContent(left, hudBottom, '╠', nil, style)
	screen.SetContent(right, hudBottom, '╣', nil, style)

	lightStyle := dividerLineStyle().Bold(true)
	for y := hudBottom + 2; y < bottom-1; y += 4 {
		screen.SetContent(left+1, y, '•', nil, lightStyle)
		screen.SetContent(right-1, y, '•', nil, lightStyle)
	}
}

func drawSpacer(screenWidth int, y int, screen tcell.Screen) {
	style := borderStyle()
	for i := 1; i <= screenWidth-1; i++ {
		screen.SetContent(i, y, '═', nil, style)
	}
}

func drawLives(screen tcell.Screen, x int, y int, lives int) {
	style := livesStyle()
	for i := 1; i <= lives; i++ {
		screen.SetContent(x+i*2, y, '♥', nil, style)
	}
}

func drawText(screen tcell.Screen, x int, y int, text string, style tcell.Style) {
	for i, r := range text {
		screen.SetContent(x+i, y, r, nil, style)
	}
}

func DrawBall(ball *entities.Ball, screen tcell.Screen) {
	style := ballStyle()
	screen.SetContent(ball.X-1, ball.Y, '█', nil, style)
	screen.SetContent(ball.X, ball.Y, '█', nil, style)
	screen.SetContent(ball.X+1, ball.Y, '█', nil, style)
}

func DrawBricks(bricks []*entities.Brick, screen tcell.Screen) {
	for _, brick := range bricks {
		brickLengthOnEachSideOfCenter := brick.Width / 2
		style := baseStyle().Foreground(brickColorAtY(brick.Y)).Bold(true)
		for i := brick.X - brickLengthOnEachSideOfCenter; i <= brick.X+brickLengthOnEachSideOfCenter; i++ {
			if brick.Alive {
				screen.SetContent(i, brick.Y, '█', nil, style)
			}
		}
	}
}

func DrawInputHints(screenWidth, screenHeight int, screen tcell.Screen) {
	style := inputHintStyle()

	hint := "[<-][->] MOVE   [R] RESTART   [ESC] QUIT"

	y := screenHeight - constants.BorderWidth - 1
	x := (screenWidth - len(hint)) / 2

	for i, ch := range hint {
		screen.SetContent(x+i, y, ch, nil, style)
	}
}
