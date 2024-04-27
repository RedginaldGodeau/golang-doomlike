package objects

import (
	g_type "github.com/RedginaldGodeau/GoDoom/src/core/type"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Object
	Color rl.Color
}

func NewPlayer(startPos g_type.Vector2) Player {
	return Player{
		Object: Object{
			CanCollide: true,
			ZIndex:     2,
			Position:   startPos,
		},
		Color: rl.Color{117, 0, 0, 255},
	}
}

func (p *Player) Draw() {
	var x, y = p.GetPosition().Get()
	rl.DrawCircle(x, y, 3, rl.Color{255, 0, 0, 255})
}

func (p *Player) GetZIndex() int {
	return p.ZIndex
}

func (p *Player) GetPosition() g_type.Vector2 {
	return p.Position
}

func (p *Player) GetCanCollide() bool {
	return p.CanCollide
}

func (p *Player) Move(vector2 g_type.Vector2) {
	p.Position.Add(vector2)
}
