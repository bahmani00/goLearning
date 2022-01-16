package sum

func SumAll(arrayOfNums ...[]int) []int {
	var sums []int
	for _, numbers := range arrayOfNums {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(arrayOfNums ...[]int) []int {
	return
}
