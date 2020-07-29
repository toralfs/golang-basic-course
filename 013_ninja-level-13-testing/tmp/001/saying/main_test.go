package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("Dolly")

	if s != "Welcome Dolly" {
		t.Error("got ", s, "expected ", "Welcome Dolly")
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Dolly"))
	// Output:
	// Welcome Dolly
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Dolly")
	}
}
