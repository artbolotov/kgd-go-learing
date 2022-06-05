package integersgo

import "testing"

func TestAddeder(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("test addeder", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4
		assertCorrectMessage(t, sum, expected)
	})
	
	t.Run("test multiply", func(t *testing.T) {
		sum := Multiply(10, 2)
		expected := 20
		assertCorrectMessage(t, sum, expected)
	})
}
