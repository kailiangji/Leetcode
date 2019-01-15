package next_permutation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_nextPermuation(t *testing.T) {
	testFun := func(a, b []int) {
		for i := range a {
			assert.Equal(t, a[i], b[i])
		}
	}
	a1 := []int{1, 2, 3}
	nextPermutation(a1)
	testFun([]int{1, 3, 2}, a1)

	a2 := []int{3, 2, 1}
	nextPermutation(a2)
	testFun([]int{1, 2, 3}, a2)

	a3 := []int{1, 1, 5}
	nextPermutation(a3)
	testFun([]int{1, 5, 1}, a3)

	a4 := []int{1, 3, 2}
	nextPermutation(a4)
	testFun([]int{2, 1, 3}, a4)

	a5 := []int{2, 3, 1}
	nextPermutation(a5)
	testFun([]int{3, 1, 2}, a5)

	a6 := []int{5, 4, 7, 5, 3, 2}
	nextPermutation(a6)
	testFun([]int{5, 5, 2, 3, 4, 7}, a6)

	a7 := []int{2, 2, 7, 5, 4, 3, 2, 2, 1}
	nextPermutation(a7)
	testFun([]int{2, 3, 1, 2, 2, 2, 4, 5, 7}, a7)

	a8 := []int{4, 2, 0, 2, 3, 2, 0}
	nextPermutation(a8)
	testFun([]int{4, 2, 0, 3, 0, 2, 2}, a8)
}
