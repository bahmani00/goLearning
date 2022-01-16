package iteration

import "testing"

//to run this: go test ./iteration
func TestRepeat(t *testing.T) {
	t.Run("repeat 'a' 0 times", func(t *testing.T) {
		got := Repeat("a", 0)
		expected := ""

		if got != expected {
			t.Errorf("expected %q but got %q", expected, got)
		}
	})

	t.Run("repeat 'a' 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		expected := "aaaaa"

		if got != expected {
			t.Errorf("expected %q but got %q", expected, got)
		}
	})
}

//to run this: go test ./iteration -bench=.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", i)
	}
}
