package main

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, SumSlice(numbers))
	}

	return sums
}
