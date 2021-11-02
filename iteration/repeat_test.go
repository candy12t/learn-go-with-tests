package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 12)
	expected := "aaaaaaaaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 12)
	}
}

func ExampleRepeat() {
	output := Repeat("x", 7)
	fmt.Println(output)
	// Output: xxxxxxx
}
