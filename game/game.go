package game

import (
	"breakout/entities"
	"breakout/input"
	"breakout/render"

	"github.com/gdamore/tcell/v3"
)

type Game struct {
	screen   tcell.Screen
	renderer *render.Renderer
	paddle   *entities.Paddle
	running  bool
}

func NewGame(screen tcell.Screen) *Game {
	width, _ := screen.Size()
	renderer := render.NewRenderer(screen)
	paddle := entities.NewPaddle(width, 19)

	return &Game{
		screen:   screen,
		renderer: renderer,
		paddle:   paddle,
		running:  true,
	}
}

func (game *Game) Run() {
	for game.running {
		game.handleInput()
		game.render()
	}
}

func (game *Game) handleInput() {
	dx := 2

	UserAction := input.GetInput(game.screen)
	if UserAction == input.ActionExit {
		game.running = false
	}
	switch UserAction {
	case input.ActionLeftKeyPressed:
		game.paddle.Move(-dx)
	case input.ActionRightKeyPressed:
		game.paddle.Move(dx)
	}
}

func (game *Game) render() {
	game.renderer.Clear()
	game.renderer.DrawPaddle(game.paddle)
	game.screen.Show()
}
