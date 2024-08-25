package arrayslices

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}

	return
}

func SumAll(numbersToAdd ...[]int) []int {
	var sums []int
	for _, nums := range numbersToAdd {
		sums = append(sums, Sum(nums))
	}

	return sums
}

func SumAllTails(numbersToAdd ...[]int) []int {
	var sums []int
	for _, nums := range numbersToAdd {
		if len(nums) == 0 {
			sums = append(sums, 0)
			} else {
				tail := nums[1:]
				sums = append(sums, Sum(tail))
			}
	}
	return sums
}