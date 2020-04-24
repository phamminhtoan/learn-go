package compare

import (
	"fmt"
	"testing"
)

type testData struct {
	a    int
	b    int
	want int
}

func TestLarger(t *testing.T) {
	tests := []testData{
		{a: 2, b: 3, want: 3},
		{a: -12, b: -100, want: -12},
		{a: 0, b: 0, want: 0},
	}
	for _, test := range tests {
		got := Larger(test.a, test.b)
		if got != test.want {
			t.Errorf("Larger(%d,%d) = \"%d\", want: \"%d\"", test.a, test.b, got, test.want)
		}
	}
}

func TestFirstLarger(t *testing.T) {
	want := 2
	got := Larger(2, 1)
	if got != want {
		t.Error(errorString(2, 1, got, want))
	}
}

func TestSecondLarger(t *testing.T) {
	want := -2
	got := Larger(-2, -10)
	if got != want {
		t.Error(errorString(-2, -10, got, want))
	}
}

func errorString(a, b, got, want int) string {
	return fmt.Sprintf("Larger(%d, %d) = \"%d\", want \"%d\"", a, b, got, want)
}
