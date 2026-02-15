package input

import (
	"github.com/gdamore/tcell/v3"
)

type InputAction int

const (
	ActionNone InputAction = iota
	ActionExit
	ActionLeftKeyPressed
	ActionRightKeyPressed
	ActionEnterKeyPressed
	ActionRKeyPressed
)

func GetInput(s tcell.Screen, userInputChannel chan InputAction) {
	for {
		ev := <-s.EventQ()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape, tcell.KeyCtrlC:
				userInputChannel <- ActionExit
			case tcell.KeyLeft:
				userInputChannel <- ActionLeftKeyPressed
			case tcell.KeyRight:
				userInputChannel <- ActionRightKeyPressed
			case tcell.KeyEnter:
				userInputChannel <- ActionEnterKeyPressed
			case tcell.KeyRune:
				keyPressed := ev.Name()
				switch keyPressed {
				case "Rune[r]", "Rune[R]":
					userInputChannel <- ActionRKeyPressed
				}
			}
		}
	}
}
