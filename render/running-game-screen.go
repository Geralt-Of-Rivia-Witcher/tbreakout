package render

import (
	"breakout/constants"
	"breakout/entities"
	"fmt"

	"github.com/gdamore/tcell/v3"
)

type Renderer struct {
	screen tcell.Screen
}

func NewRenderer(screen tcell.Screen) *Renderer {
	return &Renderer{
		screen: screen,
	}
}

func (renderer *Renderer) Clear() {
	renderer.screen.Clear()
}

func (renderer *Renderer) DrawPaddle(paddle *entities.Paddle) {
	style := paddleStyle()
	paddleLengthOnEachSideOfCenter := paddle.Width / 2
	for i := paddle.X - paddleLengthOnEachSideOfCenter; i <= paddle.X+paddleLengthOnEachSideOfCenter; i++ {
		renderer.screen.SetContent(i, paddle.Y, '█', nil, style)
	}
}

func (renderer *Renderer) DrawHUD(lives int, score int, screenWidth int, screenHeight int) {
	border := borderStyle()
	label := hudLabelStyle()

	drawSpacer(screenWidth, constants.BorderWidth, renderer.screen)

	drawText(renderer.screen, screenWidth/3, 3, "LIVES: ", label)
	drawLives(renderer.screen, screenWidth/3+5, 3, lives)
	drawText(renderer.screen, screenWidth/3*2, 3, fmt.Sprintf("SCORE: %d", score), scoreStyle())

	drawSpacer(screenWidth, constants.TopHUDElementHeight, renderer.screen)
	drawBroders(renderer.screen, screenWidth, screenHeight, border)

	renderer.DrawInputHints(screenWidth, screenHeight)
}

func drawBroders(screen tcell.Screen, screenWidth int, screenHeight int, style tcell.Style) {
	for i := 1; i <= screenWidth; i++ {
		screen.SetContent(i, screenHeight-constants.BottomBorderHeight, '━', nil, style)
	}
	for i := 1; i < screenHeight-constants.BottomBorderHeight; i++ {
		screen.SetContent(constants.BorderWidth, i, '|', nil, style)
		screen.SetContent(screenWidth-constants.BorderWidth, i, '|', nil, style)
	}
	screen.SetContent(constants.BorderWidth, screenHeight-constants.BottomBorderHeight, '└', nil, style)
	screen.SetContent(screenWidth-constants.BorderWidth, screenHeight-constants.BottomBorderHeight, '┘', nil, style)
	screen.SetContent(constants.BorderWidth, constants.BorderWidth, '┌', nil, style)
	screen.SetContent(screenWidth-constants.BorderWidth, constants.BorderWidth, '┐', nil, style)
	screen.SetContent(screenWidth-constants.BorderWidth, constants.BorderWidth, '┐', nil, style)
	screen.SetContent(constants.BorderWidth, constants.TopHUDElementHeight, '├', nil, style)
	screen.SetContent(screenWidth-constants.BorderWidth, constants.TopHUDElementHeight, '┤', nil, style)
}

func drawSpacer(screenWidth int, y int, screen tcell.Screen) {
	style := borderStyle()
	for i := 1; i <= screenWidth-1; i++ {
		screen.SetContent(i, y, '=', nil, style)
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

func (renderer *Renderer) DrawBall(ball *entities.Ball) {
	style := ballStyle()
	renderer.screen.SetContent(ball.X-1, ball.Y, '▓', nil, style)
	renderer.screen.SetContent(ball.X, ball.Y, '▓', nil, style)
	renderer.screen.SetContent(ball.X+1, ball.Y, '▓', nil, style)
}

func (renderer *Renderer) DrawBricks(bricks []*entities.Brick) {
	for _, brick := range bricks {
		brickLengthOnEachSideOfCenter := brick.Width / 2
		style := baseStyle().Foreground(brickColorAtY(brick.Y)).Bold(true)
		for i := brick.X - brickLengthOnEachSideOfCenter; i <= brick.X+brickLengthOnEachSideOfCenter; i++ {
			if brick.Alive {
				renderer.screen.SetContent(i, brick.Y, '▓', nil, style)
			}
		}
	}
}

func (renderer *Renderer) DrawInputHints(screenWidth, screenHeight int) {
	style := inputHintStyle()

	hint := "< > MOVE    ESC QUIT"

	y := screenHeight - constants.BorderWidth - 1
	x := (screenWidth - len(hint)) / 2

	for i, ch := range hint {
		renderer.screen.SetContent(x+i, y, ch, nil, style)
	}
}
