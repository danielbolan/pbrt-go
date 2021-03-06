package main

import (
	"image/color"
	"math"
)

type Vec3 struct {
	X, Y, Z float64
}

func (a Vec3) Add(b Vec3) Vec3 {
	return Vec3{
		a.X + b.X,
		a.Y + b.Y,
		a.Z + b.Z,
	}
}

func (a Vec3) Sub(b Vec3) Vec3 {
	return Vec3{
		a.X - b.X,
		a.Y - b.Y,
		a.Z - b.Z,
	}
}

func (a Vec3) Times(t float64) Vec3 {
	return Vec3{
		a.X * t,
		a.Y * t,
		a.Z * t,
	}
}

func (a Vec3) Dot(b Vec3) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func (a Vec3) Len2() float64 {
	return a.Dot(a)
}

func (a Vec3) Len() float64 {
	return math.Sqrt(a.Len2())
}

func (a Vec3) Norm() Vec3 {
	return a.Times(1 / a.Len())
}

func (u Vec3) Cross(v Vec3) Vec3 {
	x := u.Y*v.Z - u.Z*v.Y
	y := u.Z*v.X - u.X*v.Z
	z := u.X*v.Y - u.Y*v.X
	return Vec3{x, y, z}
}

// The Hadamard product does piecewise multiplication
// Should be helpful when working with vectors representing colors
func (u Vec3) Hadamard(v Vec3) Vec3 {
	x := u.X * v.X
	y := u.Y * v.Y
	z := u.Z * v.Z
	return Vec3{x, y, z}
}

func (u Vec3) ToRGBA() color.RGBA {
	return color.RGBA{
		uint8(u.X * 255.0),
		uint8(u.Y * 255.0),
		uint8(u.Z * 255.0),
		255,
	}
}
