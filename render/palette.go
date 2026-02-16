package render

import (
	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
)

func baseStyle() tcell.Style {
	return tcell.StyleDefault.Background(color.Default)
}

func borderStyle() tcell.Style {
	return baseStyle().Foreground(color.Silver).Bold(true)
}

func hudLabelStyle() tcell.Style {
	return baseStyle().Foreground(color.LightSteelBlue).Bold(true)
}

func scoreStyle() tcell.Style {
	return baseStyle().Foreground(color.Turquoise).Bold(true)
}

func livesStyle() tcell.Style {
	return baseStyle().Foreground(color.IndianRed).Bold(true)
}

func paddleStyle() tcell.Style {
	return baseStyle().Foreground(color.DeepSkyBlue).Bold(true)
}

func ballStyle() tcell.Style {
	return baseStyle().Foreground(color.MediumOrchid).Bold(true)
}

func inputHintStyle() tcell.Style {
	return baseStyle().Foreground(color.Gainsboro)
}

func titleTextStyle() tcell.Style {
	return baseStyle().Foreground(color.DeepSkyBlue).Bold(true)
}

func subtitleTextStyle() tcell.Style {
	return baseStyle().Foreground(color.WhiteSmoke).Reverse(true).Bold(true)
}

func dividerLineStyle() tcell.Style {
	return baseStyle().Foreground(color.Gray)
}

func gameWonStyle() tcell.Style {
	return baseStyle().Foreground(color.LimeGreen).Bold(true)
}

func gameLostStyle() tcell.Style {
	return baseStyle().Foreground(color.OrangeRed).Bold(true)
}

func gameOverHintStyle() tcell.Style {
	return baseStyle().Foreground(color.Silver)
}

func brickColorAtY(y int) tcell.Color {
	palette := []tcell.Color{
		color.Tomato,
		color.Gold,
		color.MediumSpringGreen,
		color.DodgerBlue,
		color.Orchid,
	}
	return palette[(y/2)%len(palette)]
}
