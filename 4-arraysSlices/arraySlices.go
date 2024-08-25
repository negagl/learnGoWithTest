package arrayslices

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}

	return
}

func SumAll(numbers ...[]int) []int {
	length := len(numbers)
	sums := make([]int, length)

	for i, nums := range numbers {
		sums[i] = Sum(nums)
	}

	return sums
}