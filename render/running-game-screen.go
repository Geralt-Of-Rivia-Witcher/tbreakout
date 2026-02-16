package render

import (
	"breakout/constants"
	"breakout/entities"
	"fmt"

	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
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
	paddleLengthOnEachSideOfCenter := paddle.Width / 2
	for i := paddle.X - paddleLengthOnEachSideOfCenter; i <= paddle.X+paddleLengthOnEachSideOfCenter; i++ {
		renderer.screen.SetContent(i, paddle.Y, '█', nil, tcell.StyleDefault)
	}
}

func (renderer *Renderer) DrawHUD(lives int, score int, screenWidth int, screenHeight int) {
	style := tcell.StyleDefault.Foreground(color.White).Background(color.Black)

	drawSpacer(screenWidth, constants.BorderWidth, renderer.screen)

	drawText(renderer.screen, screenWidth/3, 3, "LIVES: ", style)
	drawLives(renderer.screen, screenWidth/3+5, 3, lives, style)
	drawText(renderer.screen, screenWidth/3*2, 3, fmt.Sprintf("SCORE: %d", score), style)

	drawSpacer(screenWidth, constants.TopHUDElementHeight, renderer.screen)
	drawBroders(renderer.screen, screenWidth, screenHeight, style)

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
	style := tcell.StyleDefault.Foreground(color.White).Background(color.Black)
	for i := 1; i <= screenWidth-1; i++ {
		screen.SetContent(i, y, '=', nil, style)
	}
}

func drawLives(screen tcell.Screen, x int, y int, lives int, style tcell.Style) {
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
	renderer.screen.SetContent(ball.X-1, ball.Y, '▓', nil, tcell.StyleDefault)
	renderer.screen.SetContent(ball.X, ball.Y, '▓', nil, tcell.StyleDefault)
	renderer.screen.SetContent(ball.X+1, ball.Y, '▓', nil, tcell.StyleDefault)
}

func (renderer *Renderer) DrawBricks(bricks []*entities.Brick) {
	for _, brick := range bricks {
		brickLengthOnEachSideOfCenter := brick.Width / 2
		for i := brick.X - brickLengthOnEachSideOfCenter; i <= brick.X+brickLengthOnEachSideOfCenter; i++ {
			if brick.Alive {
				renderer.screen.SetContent(i, brick.Y, '▓', nil, tcell.StyleDefault)
			}
		}
	}
}

func (renderer *Renderer) DrawInputHints(screenWidth, screenHeight int) {
	style := tcell.StyleDefault.
		Foreground(color.Gainsboro).
		Background(color.Black)

	hint := "< > MOVE    ESC QUIT"

	y := screenHeight - constants.BorderWidth - 1
	x := (screenWidth - len(hint)) / 2

	for i, ch := range hint {
		renderer.screen.SetContent(x+i, y, ch, nil, style)
	}
}
