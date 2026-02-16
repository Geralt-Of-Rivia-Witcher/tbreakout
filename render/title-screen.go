package render

import (
	"strings"

	"github.com/gdamore/tcell/v3"
)

func DrawTitleScreen(
	screen tcell.Screen,
	screenWidth, screenHeight int,
	showSubtitle bool,
) {
	screen.SetStyle(baseStyle())
	screen.Clear()

	title := buildPixelTitle("BREAKOUT")
	subtitle := "PRESS ENTER TO PLAY"

	titleHeight := len(title)
	titleWidth := len(title[0])
	startX := (screenWidth - titleWidth) / 2
	startY := (screenHeight-titleHeight)/2 - 2

	// Outer frame and starfield establish the title scene.
	drawTitleBackdrop(screen, screenWidth, screenHeight, dividerLineStyle())
	drawTitleFrame(screen, screenWidth, screenHeight)

	// Header text above the marquee area.
	drawCenteredText(screen, screenWidth, startY-2, "ARCADE SYSTEM", hudLabelStyle())

	// Draw title with a subtle horizontal shadow to feel less flat.
	for y, line := range title {
		for x, ch := range line {
			if ch == ' ' {
				continue
			}
			screen.SetContent(startX+x+1, startY+y+1, '.', nil, dividerLineStyle())
			screen.SetContent(startX+x, startY+y, ch, nil, titleRowStyle(y))
		}
	}

	dividerY := startY + titleHeight + 1
	for x := startX; x < startX+titleWidth; x++ {
		screen.SetContent(x, dividerY, '=', nil, borderStyle())
	}

	subY := dividerY + 2
	subX := (screenWidth - len(subtitle)) / 2
	if showSubtitle {
		for i, ch := range subtitle {
			screen.SetContent(subX+i, subY, ch, nil, subtitleTextStyle())
		}
	} else {
		for i := 0; i < len(subtitle); i++ {
			screen.SetContent(subX+i, subY, ' ', nil, baseStyle())
		}
	}

	screen.Show()
}

func buildPixelTitle(text string) []string {
	glyphs := map[rune][]string{
		' ': {"00000", "00000", "00000", "00000", "00000"},
		'A': {"01110", "10001", "11111", "10001", "10001"},
		'B': {"11110", "10001", "11110", "10001", "11110"},
		'E': {"11111", "10000", "11110", "10000", "11111"},
		'K': {"10001", "10010", "11100", "10010", "10001"},
		'O': {"01110", "10001", "10001", "10001", "01110"},
		'R': {"11110", "10001", "11110", "10010", "10001"},
		'T': {"11111", "00100", "00100", "00100", "00100"},
		'U': {"10001", "10001", "10001", "10001", "01110"},
	}

	const (
		glyphHeight = 5
		pixelOn     = '#'
		pixelOff    = ' '
		letterGap   = "  "
	)

	lines := make([]string, glyphHeight)
	for i, ch := range strings.ToUpper(text) {
		pattern, ok := glyphs[ch]
		if !ok {
			pattern = glyphs[' ']
		}

		for row := 0; row < glyphHeight; row++ {
			for _, bit := range pattern[row] {
				if bit == '1' {
					lines[row] += string(pixelOn)
				} else {
					lines[row] += string(pixelOff)
				}
			}
			if i < len(text)-1 {
				lines[row] += letterGap
			}
		}
	}

	return lines
}

func drawTitleBackdrop(screen tcell.Screen, screenWidth, screenHeight int, style tcell.Style) {
	for y := 1; y < screenHeight-1; y += 3 {
		for x := (y % 7); x < screenWidth-1; x += 9 {
			screen.SetContent(x, y, '.', nil, style)
		}
	}
}

func drawTitleFrame(screen tcell.Screen, screenWidth, screenHeight int) {
	frameStyle := borderStyle()
	left := 2
	right := screenWidth - 3
	top := 1
	bottom := screenHeight - 2

	for x := left + 1; x < right; x++ {
		screen.SetContent(x, top, '-', nil, frameStyle)
		screen.SetContent(x, bottom, '-', nil, frameStyle)
	}
	for y := top + 1; y < bottom; y++ {
		screen.SetContent(left, y, '|', nil, frameStyle)
		screen.SetContent(right, y, '|', nil, frameStyle)
	}
	screen.SetContent(left, top, '+', nil, frameStyle)
	screen.SetContent(right, top, '+', nil, frameStyle)
	screen.SetContent(left, bottom, '+', nil, frameStyle)
	screen.SetContent(right, bottom, '+', nil, frameStyle)
}

func titleRowStyle(row int) tcell.Style {
	palette := []tcell.Color{
		tcell.NewRGBColor(255, 110, 230),
		tcell.NewRGBColor(255, 65, 205),
		tcell.NewRGBColor(255, 40, 170),
		tcell.NewRGBColor(240, 30, 150),
		tcell.NewRGBColor(220, 20, 130),
	}
	return baseStyle().Foreground(palette[row%len(palette)]).Bold(true)
}
