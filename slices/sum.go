package arrays

// Sum a slice of integers.
func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

// Sum an indefinite amount of slices returning a slice of sums.
func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range(numbersToSum) {
		sums = append(sums, Sum(numbers))
	}
	return
}
