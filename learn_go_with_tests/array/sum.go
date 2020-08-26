package arrays

func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numberSums ...[]int) []int {
	lengthOfNumbers := len(numberSums)
	sums := make([]int, lengthOfNumbers)
	for i, numberSum := range numberSums {
		sums[i] = Sum(numberSum)
	}
	return sums
}
