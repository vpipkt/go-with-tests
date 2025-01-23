package shapes

import (
	"math"
	"testing"
)

func TestPerimter(t *testing.T) {
	t.Run("square", func(t *testing.T) {
		r := Rectangle{9.0, 9.0}
		result := r.Perimeter()
		expected := 36.0
		if result != expected {
			t.Errorf("expected %f got %f", expected, result)
		}
	})
	t.Run("rect", func(t *testing.T) {
		r := Rectangle{4.5, 1.12}
		result := r.Perimeter()
		expected := 11.24
		if result != expected {
			t.Errorf("expected %f got %f", expected, result)
		}
	})

}

func TestArea(t *testing.T) {
	// slice of anonymous struct, for table driven tests
	areaTest := []struct {
		shape   Shape
		hasArea float64
		eps     float64
	}{
		{shape: Rectangle{Width: 9.0, Height: 9.0}, hasArea: 81.0, eps: 0.0},
		{shape: Rectangle{Width: 2.0, Height: 1.25}, hasArea: 2.5, eps: 0.0},
		{shape: Circle{Radius: 10.0}, hasArea: 314.1592653589793, eps: 1e-18},
		{shape: Circle{Radius: 1.25}, hasArea: 4.908738521, eps: 1e-8},
		{shape: Triangle{Base: 3, Height: 2.5}, hasArea: 3.75, eps: 0.0},
	}

	for _, tt := range areaTest {
		got := tt.shape.Area()
		if math.Abs(got-tt.hasArea) > tt.eps {
			t.Errorf("%#v got '%g' hasArea '%g'", tt.shape, got, tt.hasArea)
		}
	}
}
