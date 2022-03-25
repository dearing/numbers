package french

// func to convert int to french
func frenchToSixteen(number int) (value string) {
	switch number {
	case 0:
		value = "zéro"
	case 1:
		value = "un"
	case 2:
		value = "deux"
	case 3:
		value = "trois"
	case 4:
		value = "quatre"
	case 5:
		value = "cinq"
	case 6:
		value = "six"
	case 7:
		value = "sept"
	case 8:
		value = "huit"
	case 9:
		value = "neuf"
	case 10:
		value = "dix"
	case 11:
		value = "onze"
	case 12:
		value = "douze"
	case 13:
		value = "treize"
	case 14:
		value = "quatorze"
	case 15:
		value = "quinze"
	case 16:
		value = "seize"
	}
	return value
}

// func to convert int to tens in french
func tens(number int) (value string) {
	switch number {
	case 1:
		value = "dix"
	case 2:
		value = "vingt"
	case 3:
		value = "trente"
	case 4:
		value = "quarante"
	case 5:
		value = "cinquante"
	case 6:
		value = "soixante"
	case 7:
		value = "septante"
	case 8:
		value = "quatre-vingt"
	case 9:
		value = "quatre-vingt-dix"
	}
	return value
}

// recursive return the french version of a number
func NumberToFrench(number int) (value string) {
	if number < 17 {
		value = frenchToSixteen(number)
	} else if number < 100 {
		if number%10 == 0 {
			value = tens(number / 10)
		} else {
			value = tens(number/10) + "-" + frenchToSixteen(number%10)
		}
	} else if number < 1000 {
		if number%100 == 0 {
			value = frenchToSixteen(number/100) + " " + "cent"
		}
		if number%100 != 0 {
			value = frenchToSixteen(number/100) + " " + "cent" + " " + NumberToFrench(number%100)
		}
	} else {
		value = "Number too big"
	}
	return value
}
