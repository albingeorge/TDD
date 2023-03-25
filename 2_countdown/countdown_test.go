package countdown

import (
	"bytes"
	"testing"
)

// Non testable
func TestCountdownNonTdd(t *testing.T) {
	CountdownNonTdd()
}

// Can test what's getting printed
func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	Countdown(buffer)
	got := buffer.String()
	want := `3
2
1
Go!`
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

type SleeperMock struct {
	count int
}

func (m *SleeperMock) Sleep() {
	m.count += 1
}

// Can also test how many sleeps are done in the function
func TestCountdownImproved(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleep := SleeperMock{}

	CountdownImproved(buffer, &sleep)

	got := buffer.String()
	want := `3
2
1
Go!`
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	sleepWanted := 3
	if sleep.count != sleepWanted {
		t.Errorf("sleep count: got %q want %q", sleep.count, sleepWanted)
	}
}
