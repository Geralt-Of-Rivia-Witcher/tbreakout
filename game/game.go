package game

import (
	"breakout/entities"
	"breakout/input"
	"breakout/render"
	"time"

	"github.com/gdamore/tcell/v3"
)

type Game struct {
	screen   tcell.Screen
	renderer *render.Renderer
	paddle   *entities.Paddle
	ball     *entities.Ball
	running  bool
}

func NewGame(screen tcell.Screen) *Game {
	width, height := screen.Size()
	renderer := render.NewRenderer(screen)
	paddle := entities.NewPaddle(width, 19)
	ball := entities.NewBall(width, height)

	return &Game{
		screen:   screen,
		renderer: renderer,
		paddle:   paddle,
		ball:     ball,
		running:  true,
	}
}

func (game *Game) Run() {
	for game.running {
		game.handleInput()
		game.detectWallCollision()
		game.detectPaddleCollision()
		game.ball.Move()
		game.render()
		time.Sleep(16 * time.Millisecond)
	}
}

func (game *Game) handleInput() {
	dx := 6

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
	game.renderer.DrawBall(game.ball)
	game.screen.Show()
}

func (game *Game) detectWallCollision() {
	width, _ := game.screen.Size()
	if game.ball.X <= 2 || game.ball.X >= width-2 {
		game.ball.Dx = -game.ball.Dx
	}
	if game.ball.Y <= 0 {
		game.ball.Dy = -game.ball.Dy
	}
}

func (game *Game) detectPaddleCollision() {
	_, height := game.screen.Size()
	if game.ball.Y == height-1 {
		paddleStart := game.paddle.X - game.paddle.Width/2
		paddleEnd := game.paddle.X + game.paddle.Width/2
		if paddleStart <= game.ball.X && game.ball.X <= paddleEnd {
			game.ball.Dy = -game.ball.Dy
		}
	}
	if game.ball.Y >= height {
		game.running = false
	}
}
