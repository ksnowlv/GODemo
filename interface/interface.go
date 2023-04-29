package main

import (
	"fmt"
	"math"
)

type Shape interface {
	getArea() int
}

type Circumference interface {
	getCircumference() int
}

type Rectangle struct {
	width  int
	height int
}

type Circle struct {
	radius float64
}

func (r Rectangle) getArea() int {
	return r.width * r.height
}

// 使用指针接收实现
func (r *Rectangle) getCircumference() int {
	return 2 * (r.width + r.height)
}

func (c Circle) getArea() int {
	return int(math.Pi * math.Pow(c.radius, 2))
}

// 使用指针接收实现
func (c *Circle) getCircumference() int {
	return int(2 * math.Pi * c.radius)
}

func outputAllArea(shapes []Shape) {

	for i := 0; i < len(shapes); i++ {
		fmt.Println("area:", shapes[i].getArea())
	}
}

func outputAllCircumference(circumferences []Circumference) {
	for i := 0; i < len(circumferences); i++ {
		fmt.Println("Circumference:", circumferences[i].getCircumference())
	}
}

func main() {

	r1 := Rectangle{2, 5}
	r2 := Rectangle{3, 4}
	c1 := Circle{3}
	c2 := Circle{4}

	shapes := []Shape{r1, r2, c1, c2}
	outputAllArea(shapes)

	circumferences := []Circumference{&r1, &r2, &c1, &c2}
	outputAllCircumference(circumferences)
}
