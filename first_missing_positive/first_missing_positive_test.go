package first_missing_positive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstMissingPositive(t *testing.T) {
	a1 := []int{1, 2, 0}
	a2 := []int{3, 4, -1, 1}
	a3 := []int{7, 8, 9, 11, 12}
	a4 := []int{2}

	assert.Equal(t, 3, firstMissingPositive(a1))
	assert.Equal(t, 2, firstMissingPositive(a2))
	assert.Equal(t, 1, firstMissingPositive(a3))
	assert.Equal(t, 1, firstMissingPositive(a4))
}
