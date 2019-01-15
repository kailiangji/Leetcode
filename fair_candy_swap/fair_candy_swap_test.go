package fair_candy_swap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fairCandySwap(t *testing.T) {
	testFun := func(a, b, o []int) {
		var sum1, sum2 int
		for i := range a {
			sum1 += a[i]
		}
		for j := range b {
			sum2 += b[j]
		}
		assert.Equal(t, sum1+(o[1]-o[0])*2, sum2)
	}
	a1 := []int{1, 1}
	b1 := []int{2, 2}

	a2 := []int{1, 2}
	b2 := []int{2, 3}

	a3 := []int{2}
	b3 := []int{1, 3}

	a4 := []int{1, 2, 5}
	b4 := []int{2, 4}

	testFun(a1, b1, fairCandySwap(a1, b1))
	testFun(a2, b2, fairCandySwap(a2, b2))
	testFun(a3, b3, fairCandySwap(a3, b3))
	testFun(a4, b4, fairCandySwap(a4, b4))
}
