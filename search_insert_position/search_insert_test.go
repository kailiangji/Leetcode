package search_insert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchInsert(t *testing.T) {
	a := []int{1, 3, 5, 6}

	assert.Equal(t, 2, searchInsert(a, 5))
	assert.Equal(t, 1, searchInsert(a, 2))
	assert.Equal(t, 4, searchInsert(a, 7))
	assert.Equal(t, 0, searchInsert(a, 0))
}
