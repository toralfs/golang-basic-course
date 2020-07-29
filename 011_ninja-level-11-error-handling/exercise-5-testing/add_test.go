package main

import (
	"testing"
)

type testpair struct {
	values []int
	sum    int
}

var tests = []testpair{
	{[]int{2, 4}, 6},
	{[]int{1, 1}, 2},
	{[]int{-1, -2}, -3},
	{[]int{-1, 1}, 0},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 45},
	{[]int{-1, -15, 13, -16, 123, 2, 9999, -9998}, 107},
}

func TestAdd(t *testing.T) {
	for _, pair := range tests {
		v := add(pair.values...)
		if v != pair.sum {
			t.Error("For ", pair.values, "expected ", pair.sum, "got ", v)
		}
	}
}
