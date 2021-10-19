package main

import "math"

type SpherePrimitive struct {
	Center Vec3
	Radius float64
	Color  Vec3
}

func (s SpherePrimitive) Intersect(ray Ray) *SurfaceInteraction {
	oc := ray.O.Sub(s.Center)
	a := ray.D.Len2()
	b := 2.0 * oc.Dot(ray.D)
	c := oc.Len2() - s.Radius*s.Radius
	d := b*b - 4*a*c
	if d < 0 {
		return nil
	} else {
		T := (-b - math.Sqrt(d)) / (2 * a)
		Position := ray.At(T)
		Normal := Position.Sub(s.Center).Norm()
		Color := s.Color.Times(Normal.Dot(ray.D) * -1)
		si := SurfaceInteraction{
			T, Position, Normal, Color,
		}
		return &si
	}
}

func (s SpherePrimitive) DoesIntersect(ray Ray) bool {
	oc := ray.O.Sub(s.Center)
	a := ray.D.Len2()
	b := 2.0 * oc.Dot(ray.D)
	c := oc.Len2() - s.Radius*s.Radius
	d := b*b - 4*a*c
	return d >= 0
}
