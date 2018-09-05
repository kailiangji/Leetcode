package minimum_path_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minPathSum(t *testing.T) {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1}}
	assert.Equal(t, 7, minPathSum(grid))
}
