package plus_one

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlusOne(t *testing.T) {
	a1 := []int{1, 2, 3}
	a2 := []int{4, 3, 2, 1}
	a3 := []int{1, 3, 9}
	a4 := []int{9, 9}

	sliceEqual := func(s1, s2 []int) {
		assert.Equal(t, len(s1), len(s2))
		for i, _ := range s1 {
			assert.Equal(t, s1[i], s2[i])
		}
	}

	sliceEqual([]int{1, 2, 4}, plusOne(a1))
	sliceEqual([]int{4, 3, 2, 2}, plusOne(a2))
	sliceEqual([]int{1, 4, 0}, plusOne(a3))
	sliceEqual([]int{1, 0, 0}, plusOne(a4))
}
