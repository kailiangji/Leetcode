package reshape

func matrixReshape(nums [][]int, r int, c int) [][]int {
	if r*c != len(nums)*len(nums[0]) {
		return nums
	}
	ans, ansLine, i := make([][]int, 0), make([]int, 0), 0
	for _, line := range nums {
		for _, n := range line {
			ansLine = append(ansLine, n)
			i++
			if i == c {
				ans, i = append(ans, ansLine), 0
				ansLine = make([]int, 0)
			}
		}
	}
	return ans
}
