package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranspose(t *testing.T) {
	a := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	b := [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}

	c := transpose(a)

	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			assert.Equal(t, b[i][j], c[i][j])
		}
	}
}

func BenchmarkTranspose(b *testing.B) {
	a := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for i := 0; i < b.N; i++ {
		_ = transpose(a)
	}
}
