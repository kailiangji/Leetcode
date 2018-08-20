package plus_one

func plusOne(digits []int) []int {
	l := len(digits)

	for i := l - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	head := []int{1}
	return append(head, digits...)
}
