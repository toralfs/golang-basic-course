package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	y := Years(6)
	if y != 42 {
		t.Error("got", y, "expected", 42)
	}
}

func TestYearsTwo(t *testing.T) {
	y2 := YearsTwo(6)
	if y2 != 42 {
		t.Error("got", y2, "expected", 42)
	}
}

func ExampleYears() {
	fmt.Printf("6 in human years is %v in dog years.\n", Years(6))
	// Output
	// 42
}

func ExampleYearsTwo() {
	fmt.Printf("6 in human years is %v in dog years.\n", YearsTwo(6))
	// Output
	// 42
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}
