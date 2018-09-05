package minimum_size_subarray_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minSubArrayLen(t *testing.T) {
	nums := []int{2, 3, 1, 2, 4, 3}
	nums1 := []int{10, 5, 13, 4, 8, 4, 5, 11, 14, 9, 16, 10, 20, 8}

	assert.Equal(t, 2, minSubArrayLen(7, nums))
	assert.Equal(t, 6, minSubArrayLen(80, nums1))

}
