package main

import (
	"github.com/RedginaldGodeau/GoDoom/src/core/game"
	g_type "github.com/RedginaldGodeau/GoDoom/src/core/type"
	"github.com/RedginaldGodeau/GoDoom/src/objects"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	var win = game.NewWindow("DoomLike", g_type.NewVector2(800, 800), 60)
	var scene = game.NewScene("map/level.mp")
	var ga = game.NewGame(win, scene, nil, nil)

	var player = objects.NewPlayer(g_type.NewVector2(100, 100))
	scene.AddObjects(&player)

	ga.SetDraw(func(game *game.Game, scene *game.Scene) {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		scene.Draw()

		rl.EndDrawing()
	})

	ga.SetRun(func(game *game.Game, scene *game.Scene) {
		if rl.IsKeyDown(rl.KeyW) {
			player.Move(g_type.NewVector2(0, -1))
		} else if rl.IsKeyDown(rl.KeyS) {
			player.Move(g_type.NewVector2(0, 1))
		}
		if rl.IsKeyDown(rl.KeyA) {
			player.Move(g_type.NewVector2(-1, 0))
		} else if rl.IsKeyDown(rl.KeyD) {
			player.Move(g_type.NewVector2(1, 0))
		}
	})

	ga.Start()
}
