package main

type Camera struct {
	Position    Vec3
	Target      Vec3
	AspectRatio float64
	PixelWidth  int
	FocalLength float64
	x, y        int
	w, h        int
}

func (c Camera) ResetIterator() {
	c.x = 0
	c.y = 0
	c.w = c.PixelWidth
	c.h = int(float64(c.PixelWidth) * c.AspectRatio)
}

func (c Camera) NextRay() Ray {
	o := Vec3{0, 0, 0}
	d := Vec3{0, 0, 1}
	r := Ray{o, d}
	c.x++
	if c.x >= c.w {
		c.y++
		c.x = 0
	}
	return r
}

func (c Camera) HasNextRay() bool {
	return c.x < c.w && c.y < c.h
}
