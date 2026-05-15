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

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, SumSlice(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {

		if len(numbers) > 0 {
			tail := numbers[1:]
			sums = append(sums, SumSlice(tail))
		} else {
			sums = append(sums, 0)
		}

	}

	return sums
}
