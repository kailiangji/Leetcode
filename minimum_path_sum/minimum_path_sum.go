package minimum_path_sum

// Given a m x n grid filled with non-negative numbers, find a path
// from top left to bottom right which minimizes the sum of all
// numbers along its path.

func minPathSum(grid [][]int) int {
	//x := len(grid) - 1
	//y := len(grid[0]) - 1
	// return minPathSum1(grid, x, y)
	//return minPathSum2(grid, x, y, grid[x][y])
	return minPathSum4(grid)
}

func minPathSum1(grid [][]int, x, y int) int {
	if x == 0 && y == 0 {
		return grid[0][0]
	}

	if x > 0 && y > 0 {
		return grid[x][y] + min(minPathSum1(grid, x-1, y), minPathSum1(grid, x, y-1))
	}

	if x == 0 {
		return grid[x][y] + minPathSum1(grid, x, y-1)
	}

	return grid[x][y] + minPathSum1(grid, x-1, y)
}

func minPathSum2(grid [][]int, x, y int, sum int) int {
	if x == 0 && y == 0 {
		return sum
	}

	if x > 0 && y > 0 {
		return min(minPathSum2(grid, x-1, y, grid[x-1][y]+sum),
			minPathSum2(grid, x, y-1, grid[x][y-1]+sum))
	}

	if x == 0 {
		return minPathSum2(grid, x, y-1, grid[x][y-1]+sum)
	}

	return minPathSum2(grid, x-1, y, grid[x-1][y]+sum)
}

// minPathSum3
// Use memorization
func minPathSum3(grid [][]int) int {
	x := len(grid)
	y := len(grid[0])
	mem := make([][]int, x)
	for i := range mem {
		mem[i] = make([]int, y)
	}

	mem[0][0] = grid[0][0]

	for j := 1; j < y; j++ {
		mem[0][j] = grid[0][j] + mem[0][j-1]
	}

	for i := 1; i < x; i++ {
		mem[i][0] = grid[i][0] + mem[i-1][0]
	}

	for i := 1; i < x; i++ {
		for j := 1; j < y; j++ {
			mem[i][j] = min(mem[i][j-1], mem[i-1][j]) + grid[i][j]
		}
	}

	return mem[x-1][y-1]
}

// minPathSum4
// Use memorization and take grid as memorization matrix
func minPathSum4(grid [][]int) int {
	x := len(grid)
	y := len(grid[0])

	for j := 1; j < y; j++ {
		grid[0][j] = grid[0][j] + grid[0][j-1]
	}

	for i := 1; i < x; i++ {
		grid[i][0] = grid[i][0] + grid[i-1][0]
	}

	for i := 1; i < x; i++ {
		for j := 1; j < y; j++ {
			grid[i][j] = min(grid[i][j-1], grid[i-1][j]) + grid[i][j]
		}
	}

	return grid[x-1][y-1]
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}
