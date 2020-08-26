package arrays

func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAllTails(numberSums ...[]int) []int {
	var sums []int
	for _, numberSum := range numberSums {
		sums = append(sums, numberSum[len(numberSum)-1])
	}
	return sums
}