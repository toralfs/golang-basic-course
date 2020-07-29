package main

import (
	"fmt"
	"math"
)

type square struct {
	w float64
	h float64
}

func (s square) area() float64 {
	return s.w * s.h
}

type circle struct {
	r float64
}

func (c circle) area() float64 {
	return math.Pi * (c.r * c.r)
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println("the area of the shape is: ", s.area())
}

func main() {
	s := square{
		w: 5,
		h: 4,
	}

	c := circle{
		r: 3,
	}

	info(s)
	info(c)
}
