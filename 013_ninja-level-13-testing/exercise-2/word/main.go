// Package word can be used to get information about word counts
package word

import (
	"strings"
)

// UseCount returns a map of how many times each word has been used.
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns the number of words in a string
func Count(s string) int {
	ss := strings.Split(s, " ")
	return len(ss)
}

// Count2 returns the number of words in a string
func Count2(s string) int {
	xs := strings.Fields(s)
	return len(xs)
}
