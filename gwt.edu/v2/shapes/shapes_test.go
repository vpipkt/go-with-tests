package shapes

import "testing"

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
		shape Shape
		want  float64
	}{
		{Rectangle{9.0, 9.0}, 81.0},
		{Rectangle{2.0, 1.25}, 2.5},
		{Circle{10.0}, 314.1592653589793},
	}

	for _, tt := range areaTest {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got '%g' want '%g'", got, tt.want)
		}
	}
}
