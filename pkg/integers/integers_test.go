package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := adder(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected %d but got %d", expected, sum)
	}
}
