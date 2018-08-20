package triangle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTriangle(t *testing.T) {
	m := make([][]int, 4)
	m[0] = []int{2}
	m[1] = []int{3, 4}
	m[2] = []int{6, 5, 7}
	m[3] = []int{4, 1, 8, 3}

	assert.Equal(t, 11, minimumTotal(m))
}
