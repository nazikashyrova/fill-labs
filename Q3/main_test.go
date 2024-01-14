package main

import (
	"reflect"
	"testing"
)

func TestFindMostRepeated(t *testing.T) {
	testCases := []struct {
		input  []string
		output string
	}{
		// Add more test cases as needed
		{[]string{"apple", "pie", "apple", "red", "red", "red"}, "red"},
		{[]string{"1", "3", "1", "4", "4", "1"}, "1"},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := FindMostRepeated(tc.input)
			if !reflect.DeepEqual(result, tc.output) {
				t.Errorf("Expected %v, got %v", tc.output, result)
			}
		})
	}
}
