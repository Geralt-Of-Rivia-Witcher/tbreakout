package render

import (
	"github.com/gdamore/tcell/v3"
)

func DrawTitleScreen(
	screen tcell.Screen,
	screenWidth, screenHeight int,
	showSubtitle bool,
) {
	title := []string{
		"######  ######  ####### ####### ##   ##  ######  ##    ## ########",
		"##   ## ##   ## ##      ##   ## ##  ##  ##    ## ##    ##    ##",
		"######  ######  #####   ####### #####   ##    ## ##    ##    ##",
		"##   ## ##   ## ##      ##   ## ##  ##  ##    ## ##    ##    ##",
		"######  ##   ## ####### ##   ## ##   ##  ######   ######     ##",
	}

	subtitle := "PRESS ENTER TO START"

	titleHeight := len(title)
	titleWidth := len(title[0])

	startX := (screenWidth - titleWidth) / 2
	startY := (screenHeight-titleHeight)/2 - 2

	titleStyle := titleTextStyle()
	dividerStyle := dividerLineStyle()
	subStyle := subtitleTextStyle()

	// Draw title
	for y, line := range title {
		for x, ch := range line {
			screen.SetContent(startX+x, startY+y, ch, nil, titleStyle)
		}
	}

	// Divider
	dividerY := startY + titleHeight + 1
	for x := 0; x < titleWidth; x++ {
		screen.SetContent(startX+x, dividerY, '-', nil, dividerStyle)
	}

	// Subtitle (blinking)
	subY := dividerY + 2
	subX := (screenWidth - len(subtitle)) / 2

	if showSubtitle {
		for i, ch := range subtitle {
			screen.SetContent(subX+i, subY, ch, nil, subStyle)
		}
	} else {
		// Clear subtitle line only
		for i := 0; i < len(subtitle); i++ {
			screen.SetContent(subX+i, subY, ' ', nil, subStyle)
		}
	}

	screen.Show()
}
