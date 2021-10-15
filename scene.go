package main

import (
	"image"
	"image/color"
)

type Scene struct {
	Aggregate Primitive
	Camera    Camera
}

func (s Scene) Render() image.Image {
	w := 100
	h := 100
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			u := float64(x)/float64(w) - 0.5
			v := float64(y)/float64(h) - 0.5
			o := Vec3{u, v, 0}
			d := Vec3{0, 0, 1}
			ray := Ray{o, d}
			if s.Aggregate.DoesIntersect(ray) {
				img.Set(x, y, color.White)
			} else {
				img.Set(x, y, color.Black)
			}
		}
	}
	return img
}
