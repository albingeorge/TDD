package main

import (
	"os"
	"time"

	countdown "github.com/albingeorge/tdd/2_countdown"
)

type sleeperImplementation struct{}

func (s sleeperImplementation) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	countdown.Countdown(os.Stdout)

	sleeper := sleeperImplementation{}
	countdown.CountdownImproved(os.Stdout, sleeper)
}
