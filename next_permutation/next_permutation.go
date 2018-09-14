package next_permutation

func nextPermutation(nums []int) {
	l := len(nums)
	if l <= 1 {
		return
	}

	for i := l - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			for j := l - 1; j >= i; j-- {
				if nums[j] > nums[i-1] {
					nums[i-1], nums[j] = nums[j], nums[i-1]
					for k := 0; k < (l-i)/2; k++ {
						nums[k+i], nums[l-1-k] = nums[l-1-k], nums[k+i]
					}
					return
				}
			}
		}
	}

	for i := 0; i < l/2; i++ {
		nums[i], nums[l-1-i] = nums[l-1-i], nums[i]
	}
	return
}
