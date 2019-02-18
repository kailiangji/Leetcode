package valid_sudoku

func isValidSudoku(board [][]byte) bool {
	numExists := make(map[byte]bool, 9)

	for _, l := range board {
		for _, e := range l {
			if e == '.' {
				continue
			}
			if e < '1' || e > '9' {
				panic("invalid input")
			}
			if numExists[e] {
				return false
			}
			numExists[e] = true
		}

		for k, _ := range numExists {
			numExists[k] = false
		}
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			e := board[j][i]
			if e == '.' {
				continue
			}
			if numExists[e] {
				return false
			}
			numExists[e] = true
		}

		for k, _ := range numExists {
			numExists[k] = false
		}
	}

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			for u := i; u < i+3; u++ {
				for v := j; v < j+3; v++ {
					e := board[u][v]
					if e == '.' {
						continue
					}
					if numExists[e] {
						return false
					}
					numExists[e] = true
				}
			}
			for k, _ := range numExists {
				numExists[k] = false
			}
		}
	}

	return true
}
