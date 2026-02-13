package render

import (
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
	for i := paddle.X - (paddle.Width / 2); i <= paddle.X+(paddle.Width/2); i++ {
		renderer.screen.SetContent(i, paddle.Y, '█', nil, tcell.StyleDefault)
	}
}

func (renderer *Renderer) DrawHUD(lives int, score int, screenWidth int, screenHeight int) {
	style := tcell.StyleDefault.Foreground(color.White)
	style = tcell.StyleDefault.Background(color.Black)

	drawSpacer(screenWidth, 1, renderer.screen)

	drawText(renderer.screen, screenWidth/3, 3, "LIVES: ", style)
	drawLives(renderer.screen, screenWidth/3+5, 3, lives, style)
	drawText(renderer.screen, screenWidth/3*2, 3, fmt.Sprintf("SCORE: %d", score), style)

	drawSpacer(screenWidth, 5, renderer.screen)
	drawBroders(renderer.screen, screenWidth, screenHeight, style)
}

func drawBroders(screen tcell.Screen, screenWidth int, screenHeight int, style tcell.Style) {
	for i := 1; i <= screenWidth; i++ {
		screen.SetContent(i, screenHeight-1, '━', nil, style)
	}
	for i := 1; i <= screenHeight; i++ {
		screen.SetContent(1, i, '|', nil, style)
		screen.SetContent(screenWidth-1, i, '|', nil, style)
	}
	screen.SetContent(1, screenHeight-1, '└', nil, style)
	screen.SetContent(screenWidth-1, screenHeight-1, '┘', nil, style)
	screen.SetContent(1, 1, '┌', nil, style)
	screen.SetContent(screenWidth-1, 1, '┐', nil, style)
	screen.SetContent(screenWidth-1, 1, '┐', nil, style)
	screen.SetContent(1, 5, '├', nil, style)
	screen.SetContent(screenWidth-1, 5, '┤', nil, style)
}

func drawSpacer(screenWidth int, y int, screen tcell.Screen) {
	style := tcell.StyleDefault.Foreground(color.White)
	style = tcell.StyleDefault.Background(color.Black)
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
		for i := brick.X - 2; i <= brick.X+2; i++ {
			if brick.Alive {
				renderer.screen.SetContent(i, brick.Y, '▓', nil, tcell.StyleDefault)
			}
		}
	}
}
