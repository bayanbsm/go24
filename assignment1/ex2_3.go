package main

import (
    "fmt"
    "math"
)

type Shape interface(){
	Area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	w float64
	h float64
}

func (c Circle) Area() float64{
	return math.Pi*c.radius*c.radius
}

func (r Rectangle) Area() float64{
	return r.w*r.h
}

func PrintArea(s Shape) {
    fmt.Printf("The area is: %.2f\n", s.Area())
}

func main() {

    c := Circle{Radius: 12}
    r := Rectangle{Width: 5, Height: 7}
    

    PrintArea(c) 
    PrintArea(r) 
}
