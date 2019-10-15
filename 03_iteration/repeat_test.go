package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a")
	expect := "aaaaa"

	if got != expect {
		t.Errorf("expected %q got %q", expect, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func ExampleRepeat() {
	result := Repeat("z")
	fmt.Println(result)
	// Output: zzzzz
}
