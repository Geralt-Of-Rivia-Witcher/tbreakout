package game

import (
	"breakout/constants"
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
	paddle := entities.NewPaddle(width, height, 23, 6)
	ball := entities.NewBall(width, height)
	bricks := entities.GenerateBricks(5, 2, width, constants.TopHUDElementHeight)

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
	userInputChannel := make(chan input.InputAction, 16)
	go input.GetInput(game.screen, userInputChannel)

	for game.running {
		width, height := game.screen.Size()
		game.handleInput(width, userInputChannel)
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
		game.render(width, height)
		time.Sleep(50 * time.Millisecond)
	}
}

func (game *Game) handleInput(screenWidth int, userInputChannel chan input.InputAction) {
	for {
		select {
		case userAction := <-userInputChannel:
			switch userAction {
			case input.ActionLeftKeyPressed:
				game.paddle.Move(-1, screenWidth)

			case input.ActionRightKeyPressed:
				game.paddle.Move(1, screenWidth)

			case input.ActionExit:
				game.running = false
			}
		default:
			return
		}
	}
}

func (game *Game) render(screenWidth int, screenHeight int) {
	game.renderer.Clear()
	game.renderer.DrawHUD(game.lives, game.score, screenWidth, screenHeight)
	game.renderer.DrawPaddle(game.paddle)
	game.renderer.DrawBall(game.ball)
	game.renderer.DrawBricks(game.bricks)
	game.screen.Show()
}
