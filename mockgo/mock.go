package mockgo

import (
	"fmt"
	"io"
	"os"
	"time"
)

type DefaultSleeper struct{}

type Sleeper interface {
	Sleep()
}

const finalWord = "Go!"
const countdownStart = 3

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// Countdown prints a countdown from 3 to out.
func Countdown(out io.Writer, sleeper Sleeper) {

	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func Mocking(){
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)	
}



