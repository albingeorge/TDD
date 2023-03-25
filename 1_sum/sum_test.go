package sum

import "testing"

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	expected := 15

	actual := Sum(numbers)

	if actual != expected {
		t.Errorf("input: %v; expected: %d; actual: %d\n", numbers, expected, actual)
	}
}
