package main

import (
	"fmt"
	"math"
)

type Shape interface {
	getArea() int
}

type Rectangle struct {
	width  int
	height int
}

type circle struct {
	radius float64
}

func (r Rectangle) getArea() int {
	return r.width * r.height
}

func (c circle) getArea() int {
	return int(math.Pi * math.Pow(c.radius, 2))
}

func outputAllArea(shapes []Shape) {

	for i := 0; i < len(shapes); i++ {
		fmt.Println("area:", shapes[i].getArea())
	}
}

func main() {

	r1 := Rectangle{2, 5}
	r2 := Rectangle{3, 4}
	c1 := circle{3}
	c2 := circle{4}

	shapes := []Shape{r1, r2, c1, c2}
	outputAllArea(shapes)
}
