package input

import "github.com/gdamore/tcell/v3"

type InputAction int

const (
	ActionNone InputAction = iota
	ActionExit
	ActionLeftKeyPressed
	ActionRightKeyPressed
	ActionUnhandledKeyPressed
)

func GetInput(s tcell.Screen) InputAction {
	select {
	case ev := <-s.EventQ():
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape, tcell.KeyCtrlC:
				return ActionExit
			case tcell.KeyLeft:
				return ActionLeftKeyPressed
			case tcell.KeyRight:
				return ActionRightKeyPressed
			}
			return ActionUnhandledKeyPressed
		}
	default:
		return ActionNone
	}
	return ActionNone
}
