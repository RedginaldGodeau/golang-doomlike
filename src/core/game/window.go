package game

import (
	_type "github.com/RedginaldGodeau/GoDoom/src/core/type"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type WindowInterface interface {
	WinInit()
}

type Window struct {
	title string
	fps   int32
	size  _type.Vector2
}

func NewWindow(winTitle string, winSize _type.Vector2, fps int32) *Window {
	return &Window{
		title: winTitle,
		fps:   fps,
		size:  winSize,
	}
}

func (w Window) WinInit() {
	width, height := w.size.Get()
	rl.InitWindow(width, height, w.title)
	rl.SetTargetFPS(w.fps)
}
