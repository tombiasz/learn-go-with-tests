package integer

import "testing"

func TestAdder(t *testing.T) {
	got := Add(2, 2)
	expect := 4

	if got != expect {
		t.Errorf("Expected %d got %d", expect, got)
	}
}
