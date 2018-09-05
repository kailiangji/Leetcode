package set_matrix_zeroes

func setZeroes(matrix [][]int) {
	x := len(matrix)
	y := len(matrix[0])

	m := make([]uint, 1)
	n := make([]uint, 1)

	for i := range matrix {
		for j, v := range matrix[i] {
			if v == 0 {
				m = setOne(m, i)
				n = setOne(n, j)
			}
		}
	}

	for i := 0; i < x; i++ {
		if getOne(m, i) != 0 {
			for j := 0; j < y; j++ {
				matrix[i][j] = 0
			}
		}
	}
	for j := 0; j < y; j++ {
		if getOne(n, j) != 0 {
			for i := 0; i < x; i++ {
				matrix[i][j] = 0
			}
		}
	}
	return
}

func setOne(bv []uint, i int) []uint {
	j := i / 64
	i = i % 64
	pos := uint(1)
	pos = pos << uint(i)

	if j == len(bv) {
		bv = append(bv, uint(0))
	}

	bv[j] |= pos

	return bv
}

func getOne(bv []uint, i int) uint {
	j := i / 64
	i = i % 64
	return (bv[j] >> uint(i)) & uint(1)
}
