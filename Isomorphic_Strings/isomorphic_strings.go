package isomorphic_strings

func isIsomorphic(s string, t string) bool {
	ls := len(s)
	lt := len(t)
	if ls != lt {
		return false
	}

	var id1, id2 int
	ms := make(map[byte]int)
	mt := make(map[byte]int)

	for i := 0; i < ls; i++ {
		if _, ex := ms[s[i]]; !ex {
			ms[s[i]] = id1
			id1++
		}
		if _, ex := mt[t[i]]; !ex {
			mt[t[i]] = id2
			id2++
		}

		if ms[s[i]] != mt[t[i]] {
			return false
		}
	}

	return true
}
