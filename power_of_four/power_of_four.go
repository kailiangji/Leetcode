package power_of_four

func isPowerOfFour(num int) bool {
	for i := 0; i < 16; i++ {
		if num == 1 {
			return true
		}

		if num&3 != 0 {
			return false
		}

		num = num >> 2
	}

	return false
}
