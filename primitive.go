package main

type Primitive interface {
	Intersect(ray Ray) float64
	DoesIntersect(ray Ray) bool
}
