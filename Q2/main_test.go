package main

import (
	"reflect"
	"testing"
)

func TestRecursiveFunc(t *testing.T) {
	testCases := []struct {
		input  int
		output []int
	}{
		// Add more test cases as needed
		{9, []int{2, 4, 9}},

		{16, []int{2, 4, 8, 16}},

		{25, []int{3, 6, 12, 25}},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := RecursiveFunc(tc.input)
			if !reflect.DeepEqual(result, tc.output) {
				t.Errorf("Expected %v, got %v", tc.output, result)
			}
		})
	}
}
