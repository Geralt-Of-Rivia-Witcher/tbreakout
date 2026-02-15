package render

import (
	"fmt"

	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
)

func DrawGameOverScreen(
	screen tcell.Screen,
	screenWidth, screenHeight int,
	score int,
	gameWon bool,
) {
	screen.Clear()

	var title string
	var titleStyle tcell.Style

	if gameWon {
		title = "YOU WIN!"
		titleStyle = tcell.StyleDefault.
			Foreground(color.Green).
			Background(color.Black).
			Bold(true)
	} else {
		title = "GAME OVER"
		titleStyle = tcell.StyleDefault.
			Foreground(color.Red).
			Background(color.Black).
			Bold(true)
	}
	dimStyle := tcell.StyleDefault.
		Foreground(color.DarkGray).
		Background(color.Black)

	valueStyle := tcell.StyleDefault.
		Foreground(color.White).
		Background(color.Black)

	restartStyle := tcell.StyleDefault.
		Foreground(color.Gray).
		Background(color.Black)

	centerY := screenHeight / 2

	// Title
	drawCenteredText(screen, screenWidth, centerY-4, title, titleStyle)

	// Divider
	drawCenteredLine(screen, screenWidth, centerY-2, '-', 25, dimStyle)

	// Score
	scoreText := fmt.Sprintf("SCORE  %d", score)
	drawCenteredText(screen, screenWidth, centerY, scoreText, valueStyle)

	// Restart hint
	drawCenteredText(
		screen,
		screenWidth,
		centerY+3,
		"PRESS R TO RESTART | ESC TO QUIT",
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
	for i := 0; i < width; i++ {
		screen.SetContent(startX+i, y, ch, nil, style)
	}
}
