package game

import (
	"time"

	"github.com/Geralt-Of-Rivia-Witcher/tbreakout/constants"
	"github.com/Geralt-Of-Rivia-Witcher/tbreakout/entities"
	"github.com/Geralt-Of-Rivia-Witcher/tbreakout/input"
	"github.com/Geralt-Of-Rivia-Witcher/tbreakout/physics"
	"github.com/Geralt-Of-Rivia-Witcher/tbreakout/render"
	"github.com/gdamore/tcell/v3"
)

type GameState int

const (
	StateTitle GameState = iota
	StatePlaying
	StateGameOver
)

type RunningGameEntities struct {
	paddle *entities.Paddle
	ball   *entities.Ball
	bricks []*entities.Brick
	score  int
	lives  int
}

type TitleScreenEntities struct {
	showSubtitle bool
}

type Game struct {
	screen              tcell.Screen
	gameState           GameState
	running             bool
	runningGameEntities *RunningGameEntities
	titleScreenEntities TitleScreenEntities
}

func NewGame(screen tcell.Screen) *Game {

	return &Game{
		screen:              screen,
		gameState:           StateTitle,
		running:             true,
		runningGameEntities: nil,
		titleScreenEntities: TitleScreenEntities{
			showSubtitle: false,
		},
	}
}

func (game *Game) setupRunningGameEntities(width int, height int) {
	game.runningGameEntities = &RunningGameEntities{
		paddle: entities.NewPaddle(width, height, 23, 6),
		ball:   entities.NewBall(width, height),
		bricks: entities.GenerateRandomLayout(width, constants.TopHUDElementHeight+3),
	}
	game.runningGameEntities.score = 0
	game.runningGameEntities.lives = 3
}

func (game *Game) Run() {
	userInputChannel := make(chan input.InputAction, 16)
	go input.GetInput(game.screen, userInputChannel)

	for game.running {
		width, height := game.screen.Size()
		game.handleInput(width, height, userInputChannel)
		switch game.gameState {
		case StateTitle:
			game.titleScreenEntities.showSubtitle = !game.titleScreenEntities.showSubtitle
			time.Sleep(400 * time.Millisecond)
		case StatePlaying:
			scoredGained, remainingBricks := game.updatePhysics(width, height)
			game.runningGameEntities.score += scoredGained
			if remainingBricks == 0 {
				game.gameState = StateGameOver
			}
			game.runningGameEntities.ball.Move()
			time.Sleep(50 * time.Millisecond)
		}
		game.renderScreen(width, height)
	}
}

func (game *Game) renderScreen(width int, height int) {
	switch game.gameState {
	case StateTitle:
		render.DrawTitleScreen(game.screen, width, height, game.titleScreenEntities.showSubtitle)
	case StatePlaying:
		render.RenderRunningGameScreen(game.screen, width, height, game.runningGameEntities.lives, game.runningGameEntities.score, game.runningGameEntities.paddle, game.runningGameEntities.bricks, game.runningGameEntities.ball)
	case StateGameOver:
		gameWon := entities.AreAllBricksDead(game.runningGameEntities.bricks)
		render.DrawGameOverScreen(game.screen, width, height, game.runningGameEntities.score, gameWon)
	}
}

func (game *Game) updatePhysics(width int, height int) (int, int) {
	physics.DetectWallCollision(width, game.runningGameEntities.ball)
	isAlive := physics.DetectPaddleCollisionAndCheckIfAlive(height, game.runningGameEntities.ball, game.runningGameEntities.paddle)
	if !isAlive {
		game.runningGameEntities.lives--
		if game.runningGameEntities.lives > 0 {
			game.runningGameEntities.paddle.ResetPaddle(width)
			game.runningGameEntities.ball.ResetBall(width, height)
		} else {
			game.gameState = StateGameOver
		}
	}
	scoreGained, remainingBricks := physics.DetectBrickCollisionAndGetScoreGainedAndRemainingBricks(game.runningGameEntities.ball, game.runningGameEntities.bricks)
	return scoreGained, remainingBricks
}

func (game *Game) handleInput(screenWidth int, screenHeight int, userInputChannel chan input.InputAction) {
	for {
		select {
		case userAction := <-userInputChannel:
			switch userAction {
			case input.ActionEnterKeyPressed:
				if game.gameState == StateTitle {
					game.gameState = StatePlaying
					game.setupRunningGameEntities(screenWidth, screenHeight)
				}

			case input.ActionLeftKeyPressed:
				if game.gameState == StatePlaying {
					game.runningGameEntities.paddle.Move(-1, screenWidth)
				}

			case input.ActionRightKeyPressed:
				if game.gameState == StatePlaying {
					game.runningGameEntities.paddle.Move(1, screenWidth)
				}

			case input.ActionExit:
				game.running = false

			case input.ActionRKeyPressed:
				if game.gameState == StatePlaying || game.gameState == StateGameOver {
					game.setupRunningGameEntities(screenWidth, screenHeight)
					game.gameState = StatePlaying
				}
			}
		default:
			return
		}
	}
}
