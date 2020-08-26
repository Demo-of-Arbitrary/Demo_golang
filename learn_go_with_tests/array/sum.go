package arrays

func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbers []int, numbers2 []int) [2]int {
	var sum [2]int
	sum1 := Sum(numbers)
	sum2 := Sum(numbers2)
	sum[0] = sum1
	sum[1] = sum2
	return sum
}
