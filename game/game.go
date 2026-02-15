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

type GameState int

const (
	StateTitle GameState = iota
	StatePlaying
	StateGameOver
)

type Game struct {
	screen    tcell.Screen
	renderer  *render.Renderer
	paddle    *entities.Paddle
	ball      *entities.Ball
	bricks    []*entities.Brick
	score     int
	lives     int
	gameState GameState
	running   bool
}

func NewGame(screen tcell.Screen) *Game {
	width, height := screen.Size()
	renderer := render.NewRenderer(screen)
	paddle := entities.NewPaddle(width, height, 23, 6)
	ball := entities.NewBall(width, height)
	bricks := entities.GenerateBricks(5, 2, width, constants.TopHUDElementHeight)

	return &Game{
		screen:    screen,
		renderer:  renderer,
		paddle:    paddle,
		ball:      ball,
		bricks:    bricks,
		score:     0,
		lives:     3,
		gameState: StateTitle,
		running:   true,
	}
}

func (game *Game) Run() {
	userInputChannel := make(chan input.InputAction, 16)
	go input.GetInput(game.screen, userInputChannel)

	showSubtitle := false
	for game.running {
		width, height := game.screen.Size()
		game.handleInput(width, userInputChannel, game.gameState)
		switch game.gameState {
		case StateTitle:
			showSubtitle = !showSubtitle
			render.DrawTitleScreen(game.screen, width, height, showSubtitle)
			time.Sleep(400 * time.Millisecond)
		case StatePlaying:
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
}

func (game *Game) handleInput(screenWidth int, userInputChannel chan input.InputAction, gameState GameState) {
	for {
		select {
		case userAction := <-userInputChannel:
			switch userAction {
			case input.ActionEnterKeyPressed:
				if gameState == StateTitle {
					game.gameState = StatePlaying
				}

			case input.ActionLeftKeyPressed:
				if gameState == StatePlaying {
					game.paddle.Move(-1, screenWidth)
				}

			case input.ActionRightKeyPressed:
				if gameState == StatePlaying {
					game.paddle.Move(1, screenWidth)
				}

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
