package toeplitz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isToeplitzMatrix(t *testing.T) {
	a := [][]int{
		{1, 2, 3, 4},
		{5, 1, 2, 3},
		{9, 5, 1, 2}}
	assert.True(t, isToeplitzMatrix(a))

	b := [][]int{
		{1, 2},
		{2, 2}}
	assert.False(t, isToeplitzMatrix(b))

	c := [][]int{
		{36, 59, 71, 15, 26, 82, 87},
		{56, 36, 59, 71, 15, 26, 82},
		{15, 0, 36, 59, 71, 15, 26}}
	assert.False(t, isToeplitzMatrix(c))

	d := [][]int{
		{44, 35, 39},
		{15, 44, 35},
		{17, 15, 44},
		{80, 17, 15},
		{43, 80, 0},
		{77, 43, 80}}

	assert.False(t, isToeplitzMatrix(d))
}
