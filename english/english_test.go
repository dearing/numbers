package english

import (
	"testing"
)

var testCases = []struct {
	number int
	result string
}{
	{0, "zero"},
	{1, "one"},
	{2, "two"},
	{3, "three"},
	{4, "four"},
	{5, "five"},
	{6, "six"},
	{7, "seven"},
	{8, "eight"},
	{9, "nine"},
	{10, "ten"},
	{11, "eleven"},
	{12, "twelve"},
	{13, "thirteen"},
	{14, "fourteen"},
	{15, "fifteen"},
	{16, "sixteen"},
	{17, "seventeen"},
	{18, "eighteen"},
	{19, "nineteen"},
	{20, "twenty"},
	{21, "twenty-one"},
	{22, "twenty-two"},
	{23, "twenty-three"},
	{24, "twenty-four"},
	{25, "twenty-five"},
	{26, "twenty-six"},
	{27, "twenty-seven"},
	{28, "twenty-eight"},
	{29, "twenty-nine"},
	{30, "thirty"},
	{40, "forty"},
	{50, "fifty"},
	{60, "sixty"},
	{70, "seventy"},
	{80, "eighty"},
	{90, "ninety"},
	{100, "one hundred"},
	{101, "one hundred and one"},
	{102, "one hundred and two"},
	{103, "one hundred and three"},
	{104, "one hundred and four"},
	{105, "one hundred and five"},
	{106, "one hundred and six"},
	{107, "one hundred and seven"},
	{108, "one hundred and eight"},
	{109, "one hundred and nine"},
	{110, "one hundred and ten"},
	{111, "one hundred and eleven"},
	{112, "one hundred and twelve"},
	{113, "one hundred and thirteen"},
	{114, "one hundred and fourteen"},
	{115, "one hundred and fifteen"},
	{116, "one hundred and sixteen"},
	{117, "one hundred and seventeen"},
	{118, "one hundred and eighteen"},
	{119, "one hundred and nineteen"},
	{120, "one hundred and twenty"},
	{121, "one hundred and twenty-one"},
	{130, "one hundred and thirty"},
	{140, "one hundred and forty"},
	{150, "one hundred and fifty"},
	{160, "one hundred and sixty"},
	{170, "one hundred and seventy"},
	{180, "one hundred and eighty"},
	{190, "one hundred and ninety"},
	{200, "two hundred"},
	{300, "three hundred"},
	{1_000, "one thousand"},
	{1_001, "one thousand, one"},
	{1_002, "one thousand, two"},
	{1_100, "one thousand, one hundred"},
	{1_101, "one thousand, one hundred and one"},
	{1_102, "one thousand, one hundred and two"},
	{1_200, "one thousand, two hundred"},
	{1_201, "one thousand, two hundred and one"},
	{2_000, "two thousand"},
	{10_000, "ten thousand"},
	{100_000, "one hundred thousand"},
	{1_000_000, "one million"},
	{1_000_000_000, "one billion"},
	{1_000_000_000_000, "one trillion"},
	{1_000_000_000_000_000, "one quadrillion"},
	{-1_000_000_000_000_000, "minus one quadrillion"},
}

func TestNumberToEnglish(t *testing.T) {
	// range over our test cases
	for _, testCase := range testCases {
		result := NumberToEnglish(testCase.number)
		if result != testCase.result {
			t.Errorf("with NumberToEnglish(%d) => get `%s`, want `%s`", testCase.number, result, testCase.result)
		}
	}

}
