package reshape

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatrixReshape(t *testing.T) {
	assertFun := func(m1, m2 [][]int) {
		for i := range m1 {
			for j := range m1[i] {
				assert.Equal(t, m1[i][j], m2[i][j])
			}
		}
	}

	m1 := [][]int{
		{1, 2},
		{3, 4}}
	m2 := [][]int{{1, 2, 3, 4}}
	assertFun(m2, matrixReshape(m1, 1, 4))

	m3 := [][]int{
		{1, 2, 3},
		{4, 5, 6}}
	m4 := [][]int{
		{1, 2},
		{3, 4},
		{5, 6}}
	assertFun(m4, matrixReshape(m3, 3, 2))

}
