package array_slice

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("add up a small array", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		expected := 15
		got := Sum(numbers)

		assertArraySum(t, expected, got, numbers)

	})
	t.Run("add up a empty array", func(t *testing.T) {
		numbers := []int{}
		expected := 0
		got := Sum(numbers)

		assertArraySum(t, expected, got, numbers)
	})
}

func assertArraySum(t testing.TB, expected, sum int, arr []int) {
	t.Helper()
	if sum != expected {
		t.Errorf("wanted '%d' got '%d'; array %v", expected, sum, arr)

	}
}

func ExampleSum() {
	arr := []int{2, 3, 4}
	fmt.Print(Sum(arr))
	// Output: 9
}

func BenchmarkSum(b *testing.B) {
	arr := []int{4, 5, 6, 7, 8, 9, 12, 15, 18}
	for range b.N {
		Sum(arr)
	}
}
