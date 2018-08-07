package matrix

func transpose(A [][]int) [][]int {
	l := len(A)
	c := len(A[0])

	B := make([][]int, c)
	for i := range B {
		B[i] = make([]int, l)
		for j := 0; j < l; j++ {
			B[i][j] = A[j][i]
		}
	}

	return B
}
