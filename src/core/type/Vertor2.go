package _type

type Vector2 struct {
	x int32
	y int32
}

func NewVector2(x int32, y int32) Vector2 {
	return Vector2{
		x: x,
		y: y,
	}
}

func (v Vector2) Get() (int32, int32) {
	return v.x, v.y
}

func (v *Vector2) Add(v2 Vector2) *Vector2 {
	v.x += v2.x
	v.y += v2.y
	return v
}
