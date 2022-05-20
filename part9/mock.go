package part9

import (
	"fmt"
	"io"

)

const finalWord = "Go!"
const countdownStart = 3

// Countdown prints a countdown from 3 to out.
func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, finalWord)
}




