package french

import (
	"testing"
)

var testCases = []struct {
	number int
	result string
}{
	{0, "zéro"},
	{1, "un"},
	{2, "deux"},
	{3, "trois"},
	{4, "quatre"},
	{5, "cinq"},
	{6, "six"},
	{7, "sept"},
	{8, "huit"},
	{9, "neuf"},
	{10, "dix"},
	{11, "onze"},
	{12, "douze"},
	{13, "treize"},
	{14, "quatorze"},
	{15, "quinze"},
	{16, "seize"},
	{17, "dix-sept"},
	{18, "dix-huit"},
	{19, "dix-neuf"},
	{20, "vingt"},
	{30, "trente"},
	{40, "quarante"},
	{50, "cinquante"},
	{60, "soixante"},
	{70, "septante"},
	{80, "quatre-vingt"},
	{90, "quatre-vingt-dix"},
	{100, "cent"},
	{101, "cent et un"},
	{102, "cent et deux"},
	{103, "cent et trois"},
	{104, "cent et quatre"},
	{105, "cent et cinq"},
	{106, "cent et six"},
	{107, "cent et sept"},
	{108, "cent et huit"},
	{109, "cent et neuf"},
	{110, "cent dix"},
	{111, "cent onze"},
	{112, "cent douze"},
	{113, "cent treize"},
	{114, "cent quatorze"},
	{115, "cent quinze"},
	{116, "cent seize"},
	{117, "cent dix-sept"},
	{118, "cent dix-huit"},
	{119, "cent dix-neuf"},
	{120, "cent vingt"},
	{121, "cent vingt et un"},
	{122, "cent vingt et deux"},
	{123, "cent vingt et trois"},
	{124, "cent vingt et quatre"},
	{125, "cent vingt et cinq"},
	{126, "cent vingt et six"},

	{200, "deux cents"},
	{201, "deux cents et un"},
	{202, "deux cents et deux"},
	{203, "deux cents et trois"},

	{300, "trois cents"},
	{301, "trois cents et un"},
	{302, "trois cents et deux"},
	{303, "trois cents et trois"},

	{400, "quatre cents"},
	{401, "quatre cents et un"},
	{402, "quatre cents et deux"},

	{500, "cinq cents"},

	{600, "six cents"},

	{700, "sept cents"},

	{800, "huit cents"},

	{900, "neuf cents"},

	{1000, "mille"},

	{1001, "mille et un"},
	{1002, "mille et deux"},
	{1003, "mille et trois"},
	{1004, "mille et quatre"},

	{1100, "onze mille"},
	{1101, "onze mille et un"},
	{1102, "onze mille et deux"},

	{1200, "douze mille"},
	{1201, "douze mille et un"},

	{1300, "treize mille"},

	{1400, "quatorze mille"},

	{1500, "quinze mille"},

	{1600, "seize mille"},

	{1700, "dix-sept mille"},

	{1800, "dix-huit mille"},

	{1900, "dix-neuf mille"},

	{2000, "vingt mille"},
	{2001, "vingt mille et un"},
	{2002, "vingt mille et deux"},

	{2100, "vingt et un mille"},
	{2101, "vingt et un mille et un"},
	{2102, "vingt et un mille et deux"},

	{2200, "vingt-deux mille"},

	{2300, "vingt-trois mille"},

	{2400, "vingt-quatre mille"},

	{2500, "vingt-cinq mille"},

	{2600, "vingt-six mille"},

	{2700, "vingt-sept mille"},

	{2800, "vingt-huit mille"},

	{2900, "vingt-neuf mille"},

	{3000, "trois mille"},

	{3100, "trois mille cent"},

	{3200, "trois mille deux cent"},

	{3300, "trois mille trois cent"},

	{10_000, "dix mille"},
	{11_000, "onze mille"},
	{12_000, "douze mille"},
	{13_000, "treize mille"},
	{14_000, "quatorze mille"},
	{15_000, "quinze mille"},
	{16_000, "seize mille"},
	{17_000, "dix-sept mille"},
	{18_000, "dix-huit mille"},
	{19_000, "dix-neuf mille"},
	{20_000, "vingt mille"},
	{21_000, "vingt et un mille"},
	{22_000, "vingt-deux mille"},
	{23_000, "vingt-trois mille"},

	{100_000, "cent mille"},
	{101_000, "cent mille et un"},
	{102_000, "cent mille et deux"},
	{103_000, "cent mille et trois"},
	{104_000, "cent mille et quatre"},
	{105_000, "cent mille et cinq"},
	{106_000, "cent mille et six"},
	{107_000, "cent mille et sept"},
	{108_000, "cent mille et huit"},
	{109_000, "cent mille et neuf"},
	{110_000, "cent mille dix"},
}

func TestNumberToFrench(t *testing.T) {
	for _, testCase := range testCases {
		result := NumberToFrench(testCase.number)
		if result != testCase.result {
			t.Errorf("with NumberToFrench(%d) => get `%s`, want `%s`", testCase.number, result, testCase.result)
		}
	}
}
