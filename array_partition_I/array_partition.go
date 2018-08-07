package array_partition

import "sort"

func arrayPairSum(nums []int) int {
	l := len(nums)
	if l&1 != 0 {
		panic("there should be even elements in the array")
	}
	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
	sum := 0
	for i := 1; i < l; i += 2 {
		sum += nums[i]
	}
	return sum
}
