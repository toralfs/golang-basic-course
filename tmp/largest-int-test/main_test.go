package main

import "testing"

func BenchmarkLargest(b *testing.B) {
	ii := []int{1, 9, 3, 6, 4, 10, 7, 5, 8, 2}

	for i := 0; i < b.N; i++ {
		largest(ii)
	}
}

func BenchmarkLargest2(b *testing.B) {
	ii := []int{1, 9, 3, 6, 4, 10, 7, 5, 8, 2}

	for i := 0; i < b.N; i++ {
		largest2(ii)
	}
}
