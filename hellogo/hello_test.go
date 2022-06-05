package hellogo

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := hello("Chris", "English")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Russian", func(t *testing.T) {
		got := hello("Артем", "Russian")
		want := "Привет, Артем"
		assertCorrectMessage(t, got, want)
	})
	t.Run("get Prefix", func(t *testing.T) {
		got := greetingPrefix("Russian")
		want := "Привет, "
		assertCorrectMessage(t, got, want)
	})
}