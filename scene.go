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
	h := 100
	w := int(float64(h) * s.Camera.AspectRatio)
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			u := float64(x) / float64(w)
			v := float64(y) / float64(h)
			ray := s.Camera.Ray(u, v)
			si := s.Aggregate.Intersect(ray)
			if si != nil {
				img.Set(x, y, si.Color.ToRGBA())
			} else {
				img.Set(x, y, color.Black)
			}
		}
	}
	return img
}
