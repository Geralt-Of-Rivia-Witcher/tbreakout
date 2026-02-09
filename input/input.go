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
	ev := <-s.EventQ()

	switch ev := ev.(type) {
	case *tcell.EventKey:
		if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
			return ActionExit
		}
		if ev.Key() == tcell.KeyLeft {
			return ActionLeftKeyPressed
		}
		if ev.Key() == tcell.KeyRight {
			return ActionRightKeyPressed
		}
	}
	return ActionUnhandledKeyPressed
}
