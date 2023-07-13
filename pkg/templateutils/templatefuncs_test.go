package templateutils

import "testing"

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
