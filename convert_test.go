package roman2dec

import "testing"

var tests = []struct {
	r              RomanNumber
	userInput      string // String representation of the Roman Number as given by the user.
	d              int    // Decimal representation of the number.
	expectParseErr bool   // Expect failure to parse the user input.
}{
	{
		map[RomanDigit]int{
			XL: 1,
		},
		"XL",
		40,
		false,
	},
	{
		map[RomanDigit]int{
			M:  1,
			CM: 1,
			IV: 1,
		},
		"MCMIV",
		1904,
		false,
	},
	{
		map[RomanDigit]int{
			M:  1,
			CM: 1,
			L:  1,
			IV: 1,
		},
		"MCMLIV",
		1954,
		false,
	},
	{
		map[RomanDigit]int{
			M:  2,
			X:  1,
			IV: 1,
		},
		"MMXIV",
		2014,
		false,
	},
	{
		map[RomanDigit]int{
			C: 2,
			V: 1,
			I: 2,
		},
		"CCVII",
		207,
		false,
	},
	{
		map[RomanDigit]int{
			M: 1,
			L: 1,
			X: 1,
			V: 1,
			I: 1,
		},
		"MLXVI",
		1066,
		false,
	},
	{
		// Test lowercase string input.
		map[RomanDigit]int{
			M: 1,
			L: 1,
			X: 1,
			V: 1,
			I: 1,
		},
		"mlxvi",
		1066,
		false,
	},
	{
		// Test invalid "CDCD" input.
		map[RomanDigit]int{
			D: 1,
			C: 3,
		},
		"CDCD",
		800,
		true,
	},
}

func TestRoman2Dec(t *testing.T) {
	for _, test := range tests {
		result := Roman2Dec(test.r)
		if result != test.d {
			t.Errorf("%s != %d, expected %d", test.r.String(), result, test.d)
		}
	}
}

func TestDec2Roman(t *testing.T) {
	for _, test := range tests {
		result, err := Dec2Roman(test.d)

		if err != nil {
			t.Errorf("Got unexpected error %v", err)
		} else if result.String() != test.r.String() {
			t.Errorf("%d != %s, expected %s", test.d, result.String(),
				test.r.String())
		}
	}

	// Test converting negative numbers and 0.
	if _, err := Dec2Roman(0); err == nil {
		t.Errorf("Expected error for Dec2Roman(0)")
	} else if _, err := Dec2Roman(-100); err == nil {
		t.Errorf("Expected error for Dec2Roman(-100)")
	}
}

func TestRomanString2Number(t *testing.T) {
	for _, test := range tests {
		result, err := RomanString2Number(test.userInput)
		if err != nil && !test.expectParseErr {
			t.Errorf("unexpected error: %v", err)
		} else if err == nil {
			if test.expectParseErr {
				t.Errorf("got %s, expected error", result.String())
			} else if result.String() != test.r.String() {
				t.Errorf("%s != %s, expected %s", test.userInput, result.String(),
					test.r.String())
			}
		}
	}
}
