package objects

import _type "github.com/RedginaldGodeau/GoDoom/src/core/type"

type ObjectInterface interface {
	Draw()
	GetZIndex() int
	GetPosition() _type.Vector2
	GetCanCollide() bool
}

type Object struct {
	CanCollide bool
	ZIndex     int
	Position   _type.Vector2
}
