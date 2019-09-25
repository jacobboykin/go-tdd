package repeater

import (
	"fmt"
	"testing"
)

func TestRepeater(t *testing.T) {
	repeated := Repeater("a", 7)
	expected := "aaaaaaa"

	if repeated != expected {
		t.Errorf("excpected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeater("a", 5)
	}
}

func ExampleRepeater() {
	repeated := Repeater("c", 4)
	fmt.Println(repeated)
	// Output: cccc
}
