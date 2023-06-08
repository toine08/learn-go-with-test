package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}
type Shape interface {
	Area() float64
	Perimeter() float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Perimeter() float64 {
	return 2 * (math.Pi * c.Radius)
}
func (c Circle) Area() float64 {
	return math.Pi * (c.Radius * c.Radius)
}

func (t Triangle) Perimeter() float64 {
	return t.Base * t.Height
}
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

func printArea(shape Shape) {
	fmt.Printf("The result of the area is %.2f cm2\n", shape.Area())
}
func printPerimeters(shape Shape) {
	fmt.Printf("The result of the perimeter is:%.2f cm\n", shape.Perimeter())
}

func main() {

	var calcType string
	var choiceShapes string
	var w, h, r, b float64

	fmt.Print("Do you want Area(y) or Perimeters(n) ?")
	fmt.Scanln(&calcType)
	fmt.Print("Which shapes do you want to calculate? ")
	fmt.Scan(&choiceShapes)

	switch choiceShapes {
	case "rectangle":
		fmt.Print("Enter the width and height:")
		fmt.Scan(&w, &h)
		rectangle := Rectangle{w, h}
		if calcType == "y" {
			printArea(rectangle)
		} else {
			printPerimeters(rectangle)
		}
	case "circle":
		fmt.Print("Enter the radius:")
		fmt.Scan(&r)
		circle := Circle{r}
		if calcType == "y" {
			printArea(circle)
		} else {
			printPerimeters(circle)
		}
	case "triangle":
		fmt.Print("Enter the base and height:")
		fmt.Scan(&b, &h)
		triangle := Triangle{b, h}
		if calcType == "y" {
			printArea(triangle)
		} else {
			printPerimeters(triangle)
		}
	}
}
