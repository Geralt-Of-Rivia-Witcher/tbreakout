package render

import (
	"fmt"

	"github.com/gdamore/tcell/v3"
)

func DrawGameOverScreen(
	screen tcell.Screen,
	screenWidth, screenHeight int,
	score int,
	levelCleared, gameWon bool,
) {
	screen.SetStyle(baseStyle())
	screen.Clear()

	var title string
	var titleStyle tcell.Style

	if levelCleared {
		title = "LEVEL CLEARED!"
		titleStyle = gameWonStyle()
	} else if gameWon {
		title = "YOU WON!"
		titleStyle = gameWonStyle()
	} else {
		title = "RETRY?"
		titleStyle = gameLostStyle()
	}
	dimStyle := dividerLineStyle()
	valueStyle := scoreStyle()
	restartStyle := gameOverHintStyle()

	centerY := screenHeight / 2

	// Title
	drawCenteredText(screen, screenWidth, centerY-4, title, titleStyle)

	// Divider
	drawCenteredLine(screen, screenWidth, centerY-2, '=', 29, dimStyle)

	// Score
	scoreText := fmt.Sprintf("HIGH SCORE  %d", score)
	drawCenteredText(screen, screenWidth, centerY, scoreText, valueStyle)

	var hintText string
	if levelCleared {
		hintText = "R = CONTINUE    ESC = QUIT"
	} else if gameWon {
		hintText = "R = RESTART    ESC = QUIT"
	} else {
		hintText = "R = RETRY    ESC = QUIT"
	}

	// Restart hint
	drawCenteredText(
		screen,
		screenWidth,
		centerY+3,
		hintText,
		restartStyle,
	)

	screen.Show()
}

func drawCenteredText(
	screen tcell.Screen,
	screenWidth int,
	y int,
	text string,
	style tcell.Style,
) {
	startX := (screenWidth - len(text)) / 2
	for i, ch := range text {
		screen.SetContent(startX+i, y, ch, nil, style)
	}
}

func drawCenteredLine(
	screen tcell.Screen,
	screenWidth int,
	y int,
	ch rune,
	width int,
	style tcell.Style,
) {
	startX := (screenWidth - width) / 2
	for i := range width {
		screen.SetContent(startX+i, y, ch, nil, style)
	}
}
