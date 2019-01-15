package integer_to_roman

func intToRoman(num int) string {
	if num > 3999 || num < 1 {
		panic("not valid input")
	}

	m := num / 1000
	rest := num % 1000
	c := rest / 100
	rest = rest % 100
	x := rest / 10
	rest = rest % 10
	i := rest

	roman := ""
	switch m {
	case 1:
		roman += "M"
	case 2:
		roman += "MM"
	case 3:
		roman += "MMM"
	}

	switch c {
	case 1:
		roman += "C"
	case 2:
		roman += "CC"
	case 3:
		roman += "CCC"
	case 4:
		roman += "CD"
	case 5:
		roman += "D"
	case 6:
		roman += "DC"
	case 7:
		roman += "DCC"
	case 8:
		roman += "DCCC"
	case 9:
		roman += "CM"
	}

	switch x {
	case 1:
		roman += "X"
	case 2:
		roman += "XX"
	case 3:
		roman += "XXX"
	case 4:
		roman += "XL"
	case 5:
		roman += "L"
	case 6:
		roman += "LX"
	case 7:
		roman += "LXX"
	case 8:
		roman += "LXXX"
	case 9:
		roman += "XC"
	}

	switch i {
	case 1:
		roman += "I"
	case 2:
		roman += "II"
	case 3:
		roman += "III"
	case 4:
		roman += "IV"
	case 5:
		roman += "V"
	case 6:
		roman += "VI"
	case 7:
		roman += "VII"
	case 8:
		roman += "VIII"
	case 9:
		roman += "IX"
	}

	return roman
}
