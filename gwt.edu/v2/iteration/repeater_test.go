package repeater

import "testing"

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
}

func assertString(t testing.TB, expected, got string) {
	t.Helper()
	if got != expected {
		t.Errorf("expected '%s' got '%s'", expected, got)
	}
}
