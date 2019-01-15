package degree_of_an_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findShortestSubArray(t *testing.T) {
	input1 := []int{1, 2, 2, 3, 1}
	assert.Equal(t, 2, findShortestSubArray(input1))

	input2 := []int{1, 2, 2, 3, 1, 4, 2}
	assert.Equal(t, 6, findShortestSubArray(input2))
}

func Test_findShortestSubArray1(t *testing.T) {
	input1 := []int{1, 2, 2, 3, 1}
	assert.Equal(t, 2, findShortestSubArray1(input1))

	input2 := []int{1, 2, 2, 3, 1, 4, 2}
	assert.Equal(t, 6, findShortestSubArray1(input2))
}
