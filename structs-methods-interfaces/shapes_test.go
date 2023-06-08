package main

import "testing"

func TestPerimeter(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasPeri float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasPeri: 36.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasPeri: 62.83185307179586},
		{name: "Triangle", shape: Triangle{Base: 5, Height: 5}, hasPeri: 25.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Perimeter()
		if got != tt.hasPeri {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasPeri)
		}
	}

}

/*func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	//f = float64 and .2 means print 2 decimals
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}*/

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 5, Height: 5}, hasArea: 12.5},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
		}
	}

}
