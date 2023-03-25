package countdown

import (
	"fmt"
	"io"
	"time"
)

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

func Countdown(writer io.Writer) {
	num := 3
	for i := num; i > 0; i-- {
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
