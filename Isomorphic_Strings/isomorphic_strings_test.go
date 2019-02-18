package isomorphic_strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isIsomrphic(t *testing.T) {
	assert.True(t, isIsomorphic("egg", "add"))
	assert.False(t, isIsomorphic("foo", "bar"))
	assert.True(t, isIsomorphic("paper", "title"))
}
