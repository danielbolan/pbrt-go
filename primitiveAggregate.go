package main

import "math"

type AggregatePrimitive struct {
	Primitives []Primitive
}

func (ap *AggregatePrimitive) Add(p Primitive) {
	ap.Primitives = append(ap.Primitives, p)
}

func (ap AggregatePrimitive) Intersect(ray Ray) *SurfaceInteraction {
	t := math.Inf(1)
	var retval *SurfaceInteraction = nil
	for _, p := range ap.Primitives {
		si := p.Intersect(ray)
		if si != nil && si.T > 0 {
			if si.T < t {
				t = si.T
				retval = si
			}
		}
	}
	return retval
}

func (ap AggregatePrimitive) DoesIntersect(ray Ray) bool {
	for _, p := range ap.Primitives {
		if p.DoesIntersect(ray) {
			return true
		}
	}
	return false
}
