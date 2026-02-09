package main

import (
	"breakout/entities"
	"breakout/input"
	"breakout/render"
	"log"

	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
)

func main() {
	defStyle := tcell.StyleDefault.Background(color.Black).Foreground(color.White)

	// Initialize screen
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(defStyle)
	s.EnableMouse()
	s.EnablePaste()
	s.Clear()

	quit := func() {
		// You have to catch panics in a defer, clean up, and
		// re-raise them - otherwise your application can
		// die without leaving any diagnostic trace.
		maybePanic := recover()
		s.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	width, _ := s.Size()
	paddle := entities.NewPaddle(width, 19)
	renderer := render.NewRenderer(s)

	dx := 1
	s.Show()
	for {
		// Update screen
		UserAction := input.GetInput(s)
		if UserAction == input.ActionExit {
			return
		}
		switch UserAction {
		case input.ActionLeftKeyPressed:
			paddle.Move(-dx)
		case input.ActionRightKeyPressed:
			paddle.Move(dx)
		}
		renderer.Clear()
		renderer.DrawPaddle(paddle)
		s.Show()
	}
}
