package render

import (
	"breakout/entities"

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
	_, height := renderer.screen.Size()
	for i := paddle.X - (paddle.Width / 2); i <= paddle.X+(paddle.Width/2); i++ {
		renderer.screen.SetContent(i, height-1, '█', nil, tcell.StyleDefault)
	}
}
