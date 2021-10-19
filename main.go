package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
)

func ParseScene() Scene {
	lookFrom := Vec3{0.8, 0.7, 0.3}
	lookAt := Vec3{0, 0, 1.1}
	up := Vec3{0, 1, 0}
	fov := 45.0
	aspectRatio := 1.0
	c := NewCamera(lookFrom, lookAt, up, fov, aspectRatio)

	p := AggregatePrimitive{}
	sphere := SpherePrimitive{}
	sphere.Center = Vec3{-0.2, -0.2, 1}
	sphere.Radius = 0.25
	sphere.Color = Vec3{1.0, 0.3, 0.5}
	p.Add(sphere)
	sphere.Center = Vec3{0.2, 0.2, 1}
	sphere.Color = Vec3{0.0, 0.6, 1.0}
	p.Add(sphere)

	s := Scene{}
	s.Camera = c
	s.Aggregate = p
	return s
}

func main() {
	s := ParseScene()
	img := s.Render()

	filename := "results/output.png"
	fd, err := os.Create(filename)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Saving to", filename)
	err = png.Encode(fd, img)
	if err != nil {
		log.Fatalln(err)
	}
}
