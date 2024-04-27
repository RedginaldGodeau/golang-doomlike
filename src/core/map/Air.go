package _map

import (
	_type "github.com/RedginaldGodeau/GoDoom/src/core/type"
	"github.com/RedginaldGodeau/GoDoom/src/objects"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Air struct {
	objects.Object
	Size  _type.Vector2
	Color rl.Color
}

func (b *Air) Draw() {
	var x, y = b.GetPosition().Get()
	var w, h = b.Size.Get()

	rl.DrawRectangleLines(x, y, w, h, b.Color)
}

func (b *Air) GetZIndex() int {
	return b.ZIndex
}

func (b *Air) GetPosition() _type.Vector2 {
	return b.Position
}

func (b *Air) GetCanCollide() bool {
	return b.CanCollide
}

func NewAir(pos _type.Vector2) objects.ObjectInterface {
	var x, y = pos.Get()
	return &Wall{
		Object: objects.Object{
			CanCollide: true,
			ZIndex:     0,
			Position:   _type.NewVector2(x*25, y*25),
		},
		Size:  _type.NewVector2(25, 25),
		Color: rl.Color{120, 120, 120, 255},
	}
}
