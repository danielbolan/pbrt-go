package main

import "math"

type SpherePrimitive struct {
	Center Vec3
	Radius float64
}

func (s SpherePrimitive) Intersect(ray Ray) float64 {
	oc := ray.O.Sub(s.Center)
	a := ray.D.Len2()
	b := 2.0 * oc.Dot(ray.D)
	c := oc.Len2() - s.Radius*s.Radius
	d := b*b - 4*a*c
	if d < 0 {
		return -1
	} else {
		return (-b - math.Sqrt(d)) / (2 * a)
	}
}

func (s SpherePrimitive) DoesIntersect(ray Ray) bool {
	return s.Intersect(ray) > 0
}
