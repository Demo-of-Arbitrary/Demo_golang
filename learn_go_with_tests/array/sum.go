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
		if len(numberSum) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numberSum[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
