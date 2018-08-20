package triangle

func minimumTotal(triangle [][]int) int {
	//return minimumTotal1(triangle, 0, 0)
	return minimumTotal2(triangle, triangle[len(triangle)-1], len(triangle)-1)
}

func minimumTotal1(triangle [][]int, root int, level int) int {
	if level == len(triangle)-1 {
		return triangle[level][root]
	}

	m1 := minimumTotal1(triangle, root, level+1)
	m2 := minimumTotal1(triangle, root+1, level+1)

	if m1 <= m2 {
		return triangle[level][root] + m1
	}
	return triangle[level][root] + m2
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func minimumTotal2(triangle [][]int, roots []int, level int) int {
	if level == 0 {
		return roots[0]
	}

	for i := len(triangle[level-1]) - 1; i >= 0; i-- {
		roots[i+1] = triangle[level-1][i] +
			min(roots[i], roots[i+1])
	}
	roots = roots[1:]

	return minimumTotal2(triangle, roots, level-1)
}
