package array_slice

import (
	"fmt"
	"slices"
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

func TestSumAll(t *testing.T) {
	t.Run("two slices summed", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{4, 5}
		result := SumAll(a, b)
		expected := []int{6, 9}
		if !slices.Equal(expected, result) {
			t.Errorf("expected '%d' got '%d' adding '%v' and '%v'", expected, result, a, b)
		}
	})
	t.Run("one slice summed", func(t *testing.T) {
		slice := []int{11, 12, 13}
		expected := []int{36}
		result := SumAll(slice)
		if !slices.Equal(result, expected) {
			t.Errorf("expected '%v' got '%v' for SumAll of '%v'", expected, result, slice)
		}
	})
}

func BenchmarkSum(b *testing.B) {
	arr := []int{4, 5, 6, 7, 8, 9, 12, 15, 18}
	for range b.N {
		Sum(arr)
	}
}
