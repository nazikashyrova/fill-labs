package main

func FindMostRepeated(arr []string) string {
	// Create a map to store the frequency of each element
	frequency := make(map[string]int)

	// Iterate through the array and update the frequency map
	for _, element := range arr {
		frequency[element]++
	}

	// Find the element with the bigest frequency
	var mostRepeated string
	maxFrequency := 0

	for element, count := range frequency {
		if count > maxFrequency {
			mostRepeated = element
			maxFrequency = count
		}
	}

	return mostRepeated
}
