package main

type Ray struct {
	O, D Vec3
}

func (ray Ray) At(t float64) Vec3 {
	return ray.O.Add(ray.D.Times(t))
}
