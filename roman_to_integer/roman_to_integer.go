package roman_to_integer

func romanElemToInt(s string) []int {
	result := make([]int, len(s))
	for i, e := range s {
		switch e {
		case 'I':
			result[i] = 1
		case 'V':
			result[i] = 5
		case 'X':
			result[i] = 10
		case 'L':
			result[i] = 50
		case 'C':
			result[i] = 100
		case 'D':
			result[i] = 500
		case 'M':
			result[i] = 1000
		default:
			panic("%s is not valid input", s)
		}
	}

	return result
}

func romanToInt(s string) int {
	l := len(s)
	elems := romanElemToInt(s)
	sum := 0

	for i := 0; i < l; {
		if i+1 < l {
			if elems[i] < elems[i+1] {
				sum += elems[i+1] - elems[i]
				i += 2
			} else {
				sum += elems[i]
				i++
			}
		} else {
			sum += elems[i]
			i++
		}
	}

	return sum
}
