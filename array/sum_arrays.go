package sum

func SumAll(arrayOfNums ...[]int) []int {
	lengthOfArrays := len(arrayOfNums)
	sums := make([]int, lengthOfArrays)

	for i, nums := range arrayOfNums {
		sums[i] = Sum(nums)
	}
	return sums
}
