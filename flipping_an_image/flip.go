package flip

func flipAndInvertImage(A [][]int) [][]int {
	l, c := len(A), len(A[0])
	c1 := c / 2
	for i := 0; i < l; i++ {
		for j := 0; j < c1; j++ {
			A[i][j], A[i][c-1-j] = 1-A[i][c-1-j], 1-A[i][j]
		}

		if c&1 == 1 {
			A[i][c1] = 1 - A[i][c1]
		}
	}
	return A
}
