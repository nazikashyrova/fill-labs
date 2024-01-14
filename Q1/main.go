package main

import (
	"sort"
	"strings"
)

func SortByACharacterCount(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		// First, compare by the number of 'a's
		countA_i := strings.Count(words[i], "a")
		countA_j := strings.Count(words[j], "a")

		if countA_i > countA_j {
			return true
		} else if countA_i < countA_j {
			return false
		}

		// If the number of 'a's is the same, compare by word lengths
		return len(words[i]) > len(words[j])
	})
	return words
}

func main() {
}
