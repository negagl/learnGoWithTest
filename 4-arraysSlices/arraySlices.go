package arrayslices

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}

	return
}

func SumAll(numbers ...[]int) []int {
	var sums []int
	for _, nums := range numbers {
		sums = append(sums, Sum(nums))
	}

	return sums
}