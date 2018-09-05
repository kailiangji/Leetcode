package minimum_size_subarray_sum

import "fmt"

func minSubArrayLen(s int, nums []int) int {
	var end, l int
	var sum int

	for i := range nums {
		sum += nums[i]
		if sum >= s {
			end = i
			l = end + 1
			break
		}
	}

	if l == 0 {
		return 0
	}

	for i := 1; i < len(nums); i++ {
		sum -= nums[i-1]
		if sum >= s {
			if l > end-i+1 {
				l = end - i + 1
				fmt.Println(nums[i : end+1])
			}
			continue
		}
		for j := end + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum >= s {
				end = j
				break
			}
		}
	}

	return l
}
