package missing_number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_missingNumber(t *testing.T) {
	a1 := []int{3, 0, 1}
	a2 := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	a3 := []int{0}
	a4 := []int{1, 2, 3}

	assert.Equal(t, 2, missingNumber(a1))
	assert.Equal(t, 8, missingNumber(a2))
	assert.Equal(t, 1, missingNumber(a3))
	assert.Equal(t, 0, missingNumber(a4))
}
