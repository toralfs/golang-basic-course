package word

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	n := Count("one two three")
	if n != 3 {
		t.Error("got", n, "expected", 3)
	}
}

func TestCount2(t *testing.T) {
	n := Count2("one two three")
	if n != 3 {
		t.Error("got", n, "expected", 3)
	}
}

// This is easier to understand, but less flexible, so gonna think some more on this...
func TestUseCount(t *testing.T) {
	m := UseCount("one two two three three three")

	for k, v := range m {
		switch k {
		case "one":
			if v != 1 {
				t.Error("got", v, "expected", 1)
			}
		case "two":
			if v != 2 {
				t.Error("got", v, "expected", 2)
			}
		case "three":
			if v != 3 {
				t.Error("got", v, "expected", 3)
			}
		}
	}
}

// This might not work correctly as maps are unordered, not sure...
/* func TestUseCount(t *testing.T) {
	m := UseCount("one two two three three three")

	mTest := map[string]int{"one": 1, "two": 2, "three": 3}

	for k, v := range m {
		if key, ok := mTest[k]; !ok {
			t.Error("got", k, "expected", key)
		}
		if v != mTest[k] {
			t.Error("got", v, "expected", mTest[k])
		}
	}

} */

func ExampleCount() {
	fmt.Println(Count("one two three"))
	// Output:
	// 3
}

func ExampleCount2() {
	fmt.Println(Count2("one two three"))
	// Output:
	// 3
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count("One Two Three")
	}
}

//This one seems to be a tiny bit faster
func BenchmarkCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count2("One Two Three")
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount("one two two three three three")
	}
}
