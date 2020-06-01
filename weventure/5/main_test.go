package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T){
	t.Run("5 times a", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("10 times b", func(t *testing.T) {
		repeated := Repeat("b", 3)
		expected := "bbb"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B){
	for i:= 0; i <b.N; i++ {
		Repeat("a",5)
	}
}

func ExampleRepeat() {
	result := Repeat("t",4)
	fmt.Println(result)
	// Output: tttt
}