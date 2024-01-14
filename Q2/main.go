package main

func RecursiveFunc(n int) []int {
	var output []int // Local slice to hold the result

	var recursiveHelper func(int)
	recursiveHelper = func(x int) {
		if x/2 > 1 {
			recursiveHelper(x / 2)
		}
		output = append(output, x)
	}

	recursiveHelper(n)

	return output
}
