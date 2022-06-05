package dependencyinjection

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
	fmt.Print("\n")
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Artem")
}