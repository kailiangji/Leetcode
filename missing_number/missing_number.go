package missing_number

func missingNumber(nums []int) int {
	var sum int

	for i, v := range nums {
		sum += i
		if v < len(nums) {
			sum -= v
		}
	}

	if sum == 0 {
		for _, v := range nums {
			if v == 0 {
				return len(nums)
			}
		}
		return 0
	}

	return sum
}
