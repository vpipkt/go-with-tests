package repeater

import (
	"fmt"
	"testing"
)

func TestRepeater(t *testing.T) {
	t.Run("repeat a string 5 times", func(t *testing.T) {
		expected := "aaaaa"
		got := Repeat("a", 5)
		assertString(t, expected, got)
	})
	t.Run("repeat a string 4 times", func(t *testing.T) {
		expected := "ahahahah"
		got := Repeat("ah", 4)
		assertString(t, expected, got)
	})
	t.Run("edge case: empty string", func(t *testing.T) {
		expected := ""
		got := Repeat("", 4)
		assertString(t, expected, got)
	})
	t.Run("edge case: zero", func(t *testing.T) {
		expected := ""
		got := Repeat("hello", 0)
		assertString(t, expected, got)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	result := Repeat("hi", 3)
	fmt.Print(result)
	// Output: hihihi
}

func assertString(t testing.TB, expected, got string) {
	t.Helper()
	if got != expected {
		t.Errorf("expected '%s' got '%s'", expected, got)
	}
}
