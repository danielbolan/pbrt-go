package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
)

func ParseScene() Scene {
	c := Camera{}
	c.PixelWidth = 100
	c.AspectRatio = 1.0
	c.FocalLength = 1.0
	c.Position = Vec3{0, 0, -1}
	c.Target = Vec3{0, 0, 0}

	p := AggregatePrimitive{}
	sphere := SpherePrimitive{}
	sphere.Center = Vec3{-0.2, -0.2, 1}
	sphere.Radius = 0.25
	p.Add(sphere)
	sphere.Center = Vec3{0.2, 0.2, 1}
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
