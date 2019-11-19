package main

import "testing"

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		area  float64
	}{
		{name: "Rectangle", shape: Rectangle{width: 12, height: 6}, area: 72.0},
		{name: "Circle", shape: Circle{radius: 10}, area: 314.1592653589793},
		{name: "Triangle", shape: Triangle{base: 12, height: 6}, area: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.area {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.area)
			}
		})
	}
}
