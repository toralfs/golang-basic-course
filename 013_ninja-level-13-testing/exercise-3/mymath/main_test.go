package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type testpair struct {
		values []int
		avg    float64
	}

	tests := []testpair{
		{[]int{1, 1, 1, 1}, 1},
		{[]int{1, 2, 3, 4}, 2.5},
		{[]int{10, 20, 30, 40, 50, 60, 70}, 40}}

	for _, v := range tests {
		f := CenteredAvg(v.values)
		if f != v.avg {
			t.Error("got", f, "expected", v.avg)
		}
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 2, 3, 4}))
	// Output:
	// 2.5
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{10, 20, 30, 40, 50, 60, 70})
	}
}
