package toeplitz

func isToeplitzMatrix(matrix [][]int) bool {
	x := len(matrix)
	y := len(matrix[0])
	var m1 []int
	var m2 []int

	for i := 0; i < x-1; i++ {
		m1 = matrix[i]
		m2 = matrix[i+1]
		for j := 0; j < y-1; j++ {
			if m1[j] != m2[j+1] {
				return false
			}
		}
	}

	return true
}
