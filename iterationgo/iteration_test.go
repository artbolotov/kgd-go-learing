package iterationgo

import "testing"


func TestRepeat(t *testing.T) {

	assertCorrectMessageInt := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	assertCorrectMessageString := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("test repeat", func(t *testing.T) {
		repeated := Repeat("a")
		expected := "aaaaa"
		assertCorrectMessageString(t, repeated, expected)
	})

	t.Run("test sum all numbers", func(t *testing.T) {
		repeated := SumAllNumbers(1, 3, 2000, -4, -543)
		expected := 1457
		assertCorrectMessageInt(t, repeated, expected)
	})

	t.Run("test sum positive numbers", func(t *testing.T) {
		repeated := SumPositiveNumbers(1, 3, 2000, -4)
		expected := 2004
		assertCorrectMessageInt(t, repeated, expected)
	})
}