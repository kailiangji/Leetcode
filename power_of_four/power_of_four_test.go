package power_of_four

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPowerOfFour(t *testing.T) {
	assert.True(t, isPowerOfFour(1))
	assert.True(t, isPowerOfFour(4))
	assert.True(t, isPowerOfFour(16))
	assert.False(t, isPowerOfFour(8))
}
