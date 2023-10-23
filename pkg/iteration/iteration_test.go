package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeatedVal := Repeat("a", 4)
	expected := "aaaa"

	if repeatedVal != expected {
		t.Errorf("expected %s - got %s", expected, repeatedVal)
	}
}
