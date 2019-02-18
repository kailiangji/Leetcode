package count_and_say

import "strconv"

func say(s string) string {
	result := ""
	l := len(s)

	j := 0
	for i := 0; i < l; {
		for j = i; j < l; j++ {
			if s[j] != s[i] {
				break
			}
		}
		result += strconv.Itoa(j-i) + string(s[i])
		i += j - i
	}

	return result
}

func countAndSay(n int) string {
	if n < 1 {
		panic("invalid input number")
	}

	s := "1"

	for i := 2; i <= n; i++ {
		s = say(s)
	}

	return s
}
