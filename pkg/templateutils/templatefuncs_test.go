package templateutils

import (
	"fmt"
	"testing"
)

func TestContains(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []string
		str      string
		expected bool
	}{
		{
			name:     "contains element",
			slice:    []string{"apple", "banana", "cherry"},
			str:      "banana",
			expected: true,
		},
		{
			name:     "does not contain element",
			slice:    []string{"apple", "banana", "cherry"},
			str:      "grape",
			expected: false,
		},
		{
			name:     "empty slice",
			slice:    []string{},
			str:      "apple",
			expected: false,
		},
		{
			name:     "element in slice with duplicates",
			slice:    []string{"apple", "banana", "banana"},
			str:      "banana",
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if result := Contains(tc.slice, tc.str); result != tc.expected {
				t.Errorf("Contains(%v, %v): expected %v, got %v", tc.slice, tc.str, tc.expected, result)
			}
		})
	}
}

func BenchmarkContains(b *testing.B) {
	benchmarks := []struct {
		name string
		s    []string
		str  string
	}{
		{"Small slice - at beginning", []string{"a", "b", "c", "d", "e"}, "a"},
		{"Small slice - at end", []string{"a", "b", "c", "d", "e"}, "e"},
		{"Large slice - at beginning", make([]string, 1e6), "0"},
		{"Large slice - at end", make([]string, 1e6), "999999"},
		{"Large slice - not found", make([]string, 1e6), "not found"},
	}

	// Populate large slices
	for i := range benchmarks[2].s {
		benchmarks[2].s[i] = fmt.Sprint(i)
	}
	for i := range benchmarks[3].s {
		benchmarks[3].s[i] = fmt.Sprint(i)
	}
	for i := range benchmarks[4].s {
		benchmarks[4].s[i] = fmt.Sprint(i)
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Contains(benchmark.s, benchmark.str)
			}
		})
	}
}
