package main

type Primitive interface {
	Intersect(ray Ray) *SurfaceInteraction
	DoesIntersect(ray Ray) bool
}
