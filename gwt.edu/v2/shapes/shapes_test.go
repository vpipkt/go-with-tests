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
	t.Run("square", func(t *testing.T) {
		r := Rectangle{9.0, 9.0}
		result := r.Area()
		expected := 81.0
		if result != expected {
			t.Errorf("expected %f got %f", expected, result)
		}
	})
	t.Run("rect", func(t *testing.T) {
		r := Rectangle{2.0, 1.25}
		result := r.Area()
		expected := 2.5
		if result != expected {
			t.Errorf("expected %f got %f", expected, result)
		}
	})
	t.Run("circle", func(t *testing.T) {
		c := Circle{1.25}
		expected := 4.9087
		result := c.Area()
		if result != expected {
			t.Errorf("got '%g' want '%f", result, expected)
		}

	})
}
