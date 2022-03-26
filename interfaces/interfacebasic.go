package interfaces

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func school(s shape) {
	fmt.Println("Şeklin Alanı :", s.area())
}

func InterfaceBasicTestFunc() {
	r := rectangle{width: 5, height: 7}
	school(r)

	c := circle{radius: 4}
	school(c)
}
