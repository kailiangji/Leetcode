package search_insert

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	middle := len(nums) / 2

	if target == nums[middle] {
		return middle
	}

	if target < nums[middle] {
		if middle == 0 || target > nums[middle-1] {
			return middle
		}
		return searchInsert(nums[0:middle], target)
	}

	if middle == 0 {
		return middle + 1
	}
	return middle + searchInsert(nums[middle:], target)
}
