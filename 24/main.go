/*
	Разработать программу нахождения расстояния между двумя точками,
	которые представлены в виде структуры Point с инкапсулированными
	параметрами x,y и конструктором.
*/
package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p Point) getX() float64 {
	return p.x
}

func (p Point) getY() float64 {
	return p.y
}
func Distance(p1, p2 Point) float64 {
	x1, y1 := p1.getX(), p1.getY()
	x2, y2 := p2.getX(), p2.getY()

	return math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))
}

func main() {
	p1 := NewPoint(1, 2)
	p2 := NewPoint(6, 2)

	fmt.Printf("p1: %v, p2: %v, dist(p1,p2): %.1f\n", p1, p2, Distance(*p1, *p2))

}
