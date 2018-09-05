package set_matrix_zeroes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_setZeroes(t *testing.T) {
	input1 := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1}}
	output1 := [][]int{
		{1, 0, 1},
		{0, 0, 0},
		{1, 0, 1}}

	setZeroes(input1)

	for i := range output1 {
		for j := range output1[i] {
			assert.Equal(t, output1[i][j], input1[i][j])
		}
	}

	input2 := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5}}
	output2 := [][]int{
		{0, 0, 0, 0},
		{0, 4, 5, 0},
		{0, 3, 1, 0}}

	setZeroes(input2)

	for i := range output2 {
		for j := range output2[i] {
			assert.Equal(t, output2[i][j], input2[i][j])
		}
	}

}
