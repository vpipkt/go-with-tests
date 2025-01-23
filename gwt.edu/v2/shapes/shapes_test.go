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

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got '%g' want '%g'", got, want)
		}
	}

	t.Run("square", func(t *testing.T) {
		r := Rectangle{9.0, 9.0}
		checkArea(t, r, 81.0)
	})
	t.Run("rect", func(t *testing.T) {
		r := Rectangle{2.0, 1.25}
		checkArea(t, r, 2.5)
	})
	t.Run("circle", func(t *testing.T) {
		c := Circle{10.}
		expected := 314.1592653589793
		checkArea(t, c, expected)

	})
}
