package count_and_say

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countAndSay(t *testing.T) {
	assert.Equal(t, "1", countAndSay(1))
	assert.Equal(t, "11", countAndSay(2))
	assert.Equal(t, "21", countAndSay(3))
	assert.Equal(t, "1211", countAndSay(4))
	assert.Equal(t, "111221", countAndSay(5))
}
