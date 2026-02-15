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
			}
		}
	}
}
