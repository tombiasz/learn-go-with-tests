package main

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		tail := numbers[1:]
		sums = append(sums, SumSlice(tail))
	}

	return sums
}
