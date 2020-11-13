package shapes

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 4}, hasArea: math.Pi * 16},
		{name: "Triangle", shape: Triangle{Base: 10, Height: 6}, hasArea: 30.0},
	}

	for _, tt := range areaTests {
		t.Run("", func(t *testing.T) {
			got := tt.shape.Area()
			want := tt.hasArea
			if got != want {
				t.Errorf("%#v got %f but want %f", tt.shape, got, tt.hasArea)
			}
		})
	}

}

func TestPerimeter(t *testing.T) {

	checkPerimeter := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("%#v want %f but got %f", shape, want, got)
		}
	}

	t.Run("test Circle Perimeter", func(t *testing.T) {
		shape := Circle{10.0}
		checkPerimeter(t, shape, math.Pi*10.0*2)
	})

	t.Run("test Rectangle Perimeter", func(t *testing.T) {
		shape := Rectangle{10, 10}
		checkPerimeter(t, shape, 40.0)
	})
}
