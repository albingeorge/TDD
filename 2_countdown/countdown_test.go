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
	// Create a []byte buffer
	buffer := &bytes.Buffer{}

	// Pass the buffer as a dependency to the implementation
	Countdown(buffer)

	// Fetch the content of the bufer to compute the output
	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// Create a mock implementation, which calls the Sleep() function
type SleeperMock struct {
	count int
}

// In this test, we intend to test how many times Sleep() were called
// from within the function. We can extend this for more functionality if needed.
func (m *SleeperMock) Sleep() {
	m.count++
}

func TestCountdownImproved(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleep := SleeperMock{}

	// For lack of a better name
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
