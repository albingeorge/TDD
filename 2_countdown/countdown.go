package countdown

import (
	"fmt"
	"io"
	"time"
)

// Does not matter if the time.Sleep() method is called here
// The responsibility of what Sleep() does is passed on to the
// consumer of this function.
type SleeperInterface interface {
	Sleep()
}

func CountdownNonTdd() {
	num := 3
	for i := num; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}

	fmt.Print("Go!")
}

// Accept the io.Writer interface
func Countdown(writer io.Writer) {
	num := 3
	for i := num; i > 0; i-- {
		// Replaces the Println with Fprintln, which accepts an io.Writer interface
		fmt.Fprintln(writer, i)
		time.Sleep(1 * time.Second)
	}

	fmt.Fprint(writer, "Go!")
}

func CountdownImproved(writer io.Writer, s SleeperInterface) {
	num := 3
	for i := num; i > 0; i-- {
		fmt.Fprintln(writer, i)
		s.Sleep()
	}

	fmt.Fprint(writer, "Go!")
}
