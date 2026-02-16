package main

import (
	"breakout/game"
	"log"

	"github.com/gdamore/tcell/v3"
)

func main() {
	defStyle := tcell.StyleDefault

	// Initialize screen
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	screen.SetStyle(defStyle)
	screen.EnableMouse()
	screen.EnablePaste()
	screen.Clear()

	quit := func() {
		// You have to catch panics in a defer, clean up, and
		// re-raise them - otherwise your application can
		// die without leaving any diagnostic trace.
		maybePanic := recover()
		screen.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	game := game.NewGame(screen)
	game.Run()
}
