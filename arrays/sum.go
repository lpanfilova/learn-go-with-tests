package main

// size of an array is encoded in its type
// we cannot pass numbers [4]int to this function
func SumArray(numbers [5]int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumSlice(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}
