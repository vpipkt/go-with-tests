package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("two plus three", func(t *testing.T) {
		sum := Add(2, 3)
		expected := 5
		assertCorrectSum(t, expected, sum)
	})
	t.Run("minus two plus three", func(t *testing.T) {
		sum := Add(-2, 3)
		expected := 1
		assertCorrectSum(t, expected, sum)
	})
}

func assertCorrectSum(t testing.TB, expected, sum int) {
	t.Helper()

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
