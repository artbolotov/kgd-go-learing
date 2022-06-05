package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Artem")

	got := buffer.String()
	want := "Hello, Artem"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}