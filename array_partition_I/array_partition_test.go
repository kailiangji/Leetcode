package array_partition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayPairSum(t *testing.T) {
	a := []int{1, 4, 3, 2}
	assert.Equal(t, 4, arrayPairSum(a))
}
