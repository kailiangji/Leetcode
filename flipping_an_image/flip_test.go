package flip

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlipAndInvertImage(t *testing.T) {
	assertFun := func(s1 [][]int, s2 [][]int) {
		s3 := flipAndInvertImage(s2)
		for i, v1 := range s1 {
			for j, v2 := range v1 {
				assert.Equal(t, v2, s3[i][j])
			}
		}
	}
	a1 := [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}
	c1 := [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}}
	assertFun(c1, a1)

	a2 := [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}
	c2 := [][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}}
	assertFun(c2, a2)
}
