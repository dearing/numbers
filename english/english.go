package english

func englishToTwenty(number int) (value string) {
	switch number {
	case 0:
		value = "zero"
	case 1:
		value = "one"
	case 2:
		value = "two"
	case 3:
		value = "three"
	case 4:
		value = "four"
	case 5:
		value = "five"
	case 6:
		value = "six"
	case 7:
		value = "seven"
	case 8:
		value = "eight"
	case 9:
		value = "nine"
	case 10:
		value = "ten"
	case 11:
		value = "eleven"
	case 12:
		value = "twelve"
	case 13:
		value = "thirteen"
	case 14:
		value = "fourteen"
	case 15:
		value = "fifteen"
	case 16:
		value = "sixteen"
	case 17:
		value = "seventeen"
	case 18:
		value = "eighteen"
	case 19:
		value = "nineteen"
	}
	return value
}

// return the english tens of a number
func tens(number int) (value string) {
	switch number {
	case 2:
		value = "twenty"
	case 3:
		value = "thirty"
	case 4:
		value = "forty"
	case 5:
		value = "fifty"
	case 6:
		value = "sixty"
	case 7:
		value = "seventy"
	case 8:
		value = "eighty"
	case 9:
		value = "ninety"
	}
	return value
}

// recursive return the english version of a number
func NumberToEnglish(number int) (value string) {

	minus := false

	if number < 0 {
		minus = true
		number = -number
	}

	if number < 20 {
		value = englishToTwenty(number)
	} else if number < 100 {
		if number%10 == 0 {
			value = tens(number / 10)
		} else {
			value = tens(number/10) + "-" + englishToTwenty(number%10)
		}
	} else if number < 1_000 {
		value = englishToTwenty(number/100) + " hundred"
		if number%100 != 0 {
			value += " and " + NumberToEnglish(number%100)
		}
	} else if number < 1_000_000 {
		value = NumberToEnglish(number/1_000) + " thousand"
		if number%1_000 != 0 {
			value += ", " + NumberToEnglish(number%1_000)
		}
	} else if number < 1_000_000_000 {
		value = NumberToEnglish(number/1_000_000) + " million"
		if number%1_000_000 != 0 {
			value += ", " + NumberToEnglish(number%1_000_000)
		}
	} else if number < 1_000_000_000_000 {
		value = NumberToEnglish(number/1_000_000_000) + " billion"
		if number%1_000_000_000 != 0 {
			value += ", " + NumberToEnglish(number%1_000_000_000)
		}

	} else if number < 1_000_000_000_000_000 {
		value = NumberToEnglish(number/1_000_000_000_000) + " trillion"
		if number%1_000_000_000_000 != 0 {
			value += ", " + NumberToEnglish(number%1_000_000_000_000)
		}
	} else {
		value = NumberToEnglish(number/1_000_000_000_000_000) + " quadrillion"
		if number%1_000_000_000_000_000 != 0 {
			value += ", " + NumberToEnglish(number%1_000_000_000_000_000)
		}
	}

	if minus {
		value = "minus " + value
	}

	return value
}
