package roman_to_integer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_romanToInt(t *testing.T) {
	assert.Equal(t, 3, romanToInt("III"))
	assert.Equal(t, 4, romanToInt("IV"))
	assert.Equal(t, 7, romanToInt("VII"))
	assert.Equal(t, 9, romanToInt("IX"))
	assert.Equal(t, 10, romanToInt("X"))
	assert.Equal(t, 13, romanToInt("XIII"))
	assert.Equal(t, 14, romanToInt("XIV"))
	assert.Equal(t, 19, romanToInt("XIX"))
	assert.Equal(t, 40, romanToInt("XL"))
	assert.Equal(t, 43, romanToInt("XLIII"))
	assert.Equal(t, 44, romanToInt("XLIV"))
	assert.Equal(t, 49, romanToInt("XLIX"))
	assert.Equal(t, 50, romanToInt("L"))
	assert.Equal(t, 53, romanToInt("LIII"))
	assert.Equal(t, 90, romanToInt("XC"))
	assert.Equal(t, 100, romanToInt("C"))
	assert.Equal(t, 110, romanToInt("CX"))
	assert.Equal(t, 58, romanToInt("LVIII"))
}
