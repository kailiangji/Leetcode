package first_missing_positive

func firstMissingPositive(nums []int) int {
	min := len(nums) + 1

	nums1 := make([]int, min)
	nums1[0] = 1

	for _, v := range nums {
		if v < min && v > 0 {
			nums1[v] = 1
		}
	}

	for i, v := range nums1 {
		if v == 0 {
			return i
		}
	}

	return min
}
