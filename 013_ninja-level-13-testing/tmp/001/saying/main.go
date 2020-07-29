package saying

import "fmt"

// Greet returns a string saying welcome (s string)
func Greet(s string) string {
	return fmt.Sprint("Welcome ", s)
}
