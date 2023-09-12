package utils

import (
	"math"
)

type Vector2 struct {
	X float32
	Y float32
}

func NewVector2(x float32, y float32) Vector2 {
	return Vector2{X: x, Y: y}
}

func (v Vector2) Add(v2 Vector2) Vector2 {
	return Vector2{X: v.X + v2.X, Y: v.Y + v2.Y}
}

func (v Vector2) Subtract(v2 Vector2) Vector2 {
	return Vector2{X: v.X - v2.X, Y: v.Y - v2.Y}
}

func (v Vector2) Multiply(v2 Vector2) Vector2 {
	return Vector2{X: v.X * v2.X, Y: v.Y * v2.Y}
}

func (v Vector2) Divide(v2 Vector2) Vector2 {
	return Vector2{X: v.X / v2.X, Y: v.Y / v2.Y}
}

func (v Vector2) AddScalar(s float32) Vector2 {
	return Vector2{X: v.X + s, Y: v.Y + s}
}

func (v Vector2) SubtractScalar(s float32) Vector2 {
	return Vector2{X: v.X - s, Y: v.Y - s}
}

func (v Vector2) MultiplyScalar(s float32) Vector2 {
	return Vector2{X: v.X * s, Y: v.Y * s}
}

func (v Vector2) DivideScalar(s float32) Vector2 {
	return Vector2{X: v.X / s, Y: v.Y / s}
}

func (v Vector2) Length() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}

func (v Vector2) Normalize() Vector2 {
	len := v.Length()
	if len == 0 {
		return v
	}
	return v.DivideScalar(len)
}

func (v Vector2) Invert() Vector2 {
	return Vector2{X: -v.X, Y: -v.Y}
}

func (v Vector2) SquaredLength() float32 {
	return v.X*v.X + v.Y*v.Y
}
