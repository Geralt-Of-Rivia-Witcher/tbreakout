package game

import (
	"breakout/entities"
	"breakout/input"
	"breakout/physics"
	"breakout/render"
	"time"

	"github.com/gdamore/tcell/v3"
)

type Game struct {
	screen   tcell.Screen
	renderer *render.Renderer
	paddle   *entities.Paddle
	ball     *entities.Ball
	bricks   []*entities.Brick
	score    int
	lives    int
	running  bool
}

func NewGame(screen tcell.Screen) *Game {
	width, height := screen.Size()
	renderer := render.NewRenderer(screen)
	paddle := entities.NewPaddle(width, 23, 6)
	ball := entities.NewBall(width, height)
	bricks := entities.GenerateBricks(5, 2, width)

	return &Game{
		screen:   screen,
		renderer: renderer,
		paddle:   paddle,
		ball:     ball,
		bricks:   bricks,
		score:    0,
		lives:    3,
		running:  true,
	}
}

func (game *Game) Run() {
	for game.running {
		width, height := game.screen.Size()
		game.handleInput(width)
		physics.DetectWallCollision(width, game.ball)
		isAlive := physics.DetectPaddleCollisionAndCheckIfAlive(height, game.ball, game.paddle)
		if !isAlive {
			game.lives--
			if game.lives > 0 {
				game.paddle.ResetPaddle(width)
				game.ball.ResetBall(width, height)
			} else {
				game.running = false
			}
		}
		game.score += physics.DetectBrickCollisionAndGetScoreGained(game.ball, game.bricks)
		game.ball.Move()
		game.render()
		time.Sleep(50 * time.Millisecond)
	}
}

func (game *Game) handleInput(screenWidth int) {
	UserAction := input.GetInput(game.screen)
	if UserAction == input.ActionExit {
		game.running = false
	}
	switch UserAction {
	case input.ActionLeftKeyPressed:
		game.paddle.Move(-1, screenWidth)
	case input.ActionRightKeyPressed:
		game.paddle.Move(1, screenWidth)
	case input.ActionExit:
		game.running = false
	}
}

func (game *Game) render() {
	game.renderer.Clear()
	game.renderer.DrawPaddle(game.paddle)
	game.renderer.DrawBall(game.ball)
	game.renderer.DrawBricks(game.bricks)
	game.screen.Show()
}
