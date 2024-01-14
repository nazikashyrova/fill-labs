package main

import (
	"reflect"
	"testing"
)

func TestSortByACharacterCount(t *testing.T) {
	testCases := []struct {
		input  []string
		output []string
	}{
		// Add more test cases as needed
		{[]string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"},
			[]string{"aaaasd", "aaabcd", "aab", "a", "lklklklklklklklkl", "cssssssd", "fdz", "ef", "kf", "zc", "l"}},
		{[]string{"apple", "zyx", "banana", "grape", "aardvark", "avocado"},
			[]string{"aardvark", "banana", "avocado", "apple", "grape", "zyx"}},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := SortByACharacterCount(tc.input)
			if !reflect.DeepEqual(result, tc.output) {
				t.Errorf("Expected %v, got %v", tc.output, result)
			}
		})
	}
}
