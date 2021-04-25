package main

import (
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	width  float64
	height float64
}

func (c circle) area() float64 {
	return math.Pi * (math.Pow(c.radius, 2))
}

func (r rectangle) area() float64 {
	return (r.width * r.height)
}
func main() {
	c1 := circle{2}
	r1 := rectangle{1, 2}
	shapes := []shape{c1, r1}

}
