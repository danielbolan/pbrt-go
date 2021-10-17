package main

import (
	"math"
)

type Camera struct {
	LookFrom      Vec3
	LookAt        Vec3
	Up            Vec3
	FieldOfView   float64
	AspectRatio   float64
	u, v, topLeft Vec3
}

func NewCamera(lookFrom Vec3, lookAt Vec3, up Vec3, fov float64, aspectRatio float64) Camera {
	c := Camera{}
	c.LookFrom = lookFrom
	c.LookAt = lookAt
	c.Up = up
	c.FieldOfView = fov
	c.AspectRatio = aspectRatio

	r := c.FieldOfView * math.Pi / 180.0
	h := math.Tan(r / 2)
	vh := 2 * h
	vw := c.AspectRatio * vh

	w := c.LookAt.Sub(c.LookFrom).Norm()
	c.u = c.Up.Cross(w).Norm()
	c.v = w.Cross(c.u).Times(-1)
	c.u = c.u.Times(vw)
	c.v = c.v.Times(vh)
	c.topLeft = c.LookFrom
	c.topLeft = c.topLeft.Sub(c.u.Times(0.5))
	c.topLeft = c.topLeft.Sub(c.v.Times(0.5))
	c.topLeft = c.topLeft.Add(w)

	return c
}

func (c Camera) Ray(u float64, v float64) Ray {
	o := c.LookFrom
	d := c.topLeft
	d = d.Add(c.u.Times(u))
	d = d.Add(c.v.Times(v))
	d = d.Sub(o)
	d = d.Norm()
	return Ray{o, d}
}
