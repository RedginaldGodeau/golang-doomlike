package game

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameInterface interface {
	SetRun(handler Handler)
	SetDraw(handler Handler)
	Start()
}

type Handler func(game *Game, scene *Scene)

type Game struct {
	win   *Window
	scene *Scene

	drawHandler Handler
	runHandler  Handler
}

func NewGame(win *Window, scene *Scene, run Handler, draw Handler) *Game {
	return &Game{
		win:         win,
		scene:       scene,
		drawHandler: draw,
		runHandler:  run,
	}
}

func (g *Game) SetRun(handler Handler) {
	g.runHandler = handler
}

func (g *Game) SetDraw(handler Handler) {
	g.drawHandler = handler
}

func (g *Game) Start() {
	g.win.WinInit()
	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		g.runHandler(g, g.scene)
		g.drawHandler(g, g.scene)
	}

	fmt.Println(rl.LogError)
	rl.CheckErrors()
}
