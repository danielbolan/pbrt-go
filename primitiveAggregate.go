package main

type AggregatePrimitive struct {
	Primitives []Primitive
}

func (ap *AggregatePrimitive) Add(p Primitive) {
	ap.Primitives = append(ap.Primitives, p)
}

func (ap AggregatePrimitive) Intersect(ray Ray) float64 {
	t := -1.0
	for _, p := range ap.Primitives {
		it := p.Intersect(ray)
		if it > 0 {
			if it < t || t == -1 {
				t = it
			}
		}
	}
	return t
}

func (ap AggregatePrimitive) DoesIntersect(ray Ray) bool {
	for _, p := range ap.Primitives {
		if p.DoesIntersect(ray) {
			return true
		}
	}
	return false
}
