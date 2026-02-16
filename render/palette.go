package render

import (
	"github.com/gdamore/tcell/v3"
)

var (
	arcadeBg      = tcell.NewRGBColor(7, 8, 18)
	arcadeBorder  = tcell.NewRGBColor(0, 220, 255)
	arcadeLabel   = tcell.NewRGBColor(255, 185, 0)
	arcadeScore   = tcell.NewRGBColor(57, 255, 20)
	arcadeLives   = tcell.NewRGBColor(255, 90, 160)
	arcadePaddle  = tcell.NewRGBColor(0, 255, 255)
	arcadeBall    = tcell.NewRGBColor(255, 240, 120)
	arcadeHint    = tcell.NewRGBColor(180, 200, 255)
	arcadeTitle   = tcell.NewRGBColor(255, 65, 205)
	arcadeDivider = tcell.NewRGBColor(80, 130, 220)
	arcadeWon     = tcell.NewRGBColor(125, 255, 85)
	arcadeLost    = tcell.NewRGBColor(255, 90, 90)
)

func baseStyle() tcell.Style {
	return tcell.StyleDefault.Background(arcadeBg).Foreground(tcell.ColorWhite)
}

func borderStyle() tcell.Style {
	return baseStyle().Foreground(arcadeBorder).Bold(true)
}

func hudLabelStyle() tcell.Style {
	return baseStyle().Foreground(arcadeLabel).Bold(true)
}

func scoreStyle() tcell.Style {
	return baseStyle().Foreground(arcadeScore).Bold(true)
}

func livesStyle() tcell.Style {
	return baseStyle().Foreground(arcadeLives).Bold(true)
}

func paddleStyle() tcell.Style {
	return baseStyle().Foreground(arcadePaddle).Bold(true)
}

func ballStyle() tcell.Style {
	return baseStyle().Foreground(arcadeBall).Bold(true)
}

func inputHintStyle() tcell.Style {
	return baseStyle().Foreground(arcadeHint)
}

func titleTextStyle() tcell.Style {
	return baseStyle().Foreground(arcadeTitle).Bold(true)
}

func subtitleTextStyle() tcell.Style {
	return baseStyle().Foreground(arcadeBg).Background(arcadeScore).Bold(true)
}

func dividerLineStyle() tcell.Style {
	return baseStyle().Foreground(arcadeDivider)
}

func gameWonStyle() tcell.Style {
	return baseStyle().Foreground(arcadeWon).Bold(true)
}

func gameLostStyle() tcell.Style {
	return baseStyle().Foreground(arcadeLost).Bold(true)
}

func gameOverHintStyle() tcell.Style {
	return baseStyle().Foreground(arcadeHint)
}

func brickColorAtY(y int) tcell.Color {
	palette := []tcell.Color{
		tcell.NewRGBColor(255, 80, 80),
		tcell.NewRGBColor(255, 200, 40),
		tcell.NewRGBColor(70, 255, 160),
		tcell.NewRGBColor(0, 195, 255),
		tcell.NewRGBColor(220, 120, 255),
	}
	return palette[(y/2)%len(palette)]
}
