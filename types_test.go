package roman2dec

import (
	"reflect"
	"testing"
)

// TestSortedRomanDigits tests the SortedRomanDigits correctly lists all the
// RomanDigits in descending order.
func TestSortedRomanDigits(t *testing.T) {
	expect := []RomanDigit{M, CM, D, CD, C, XC, L, XL, X, IX, V, IV, I}
	if !reflect.DeepEqual(SortedRomanDigits, expect) {
		t.Error("Sorted Roman Digit list is not correct.")
	}
}

// TestRomanNumberString tests we are correctly converting RomanNUmber to
// string.
func TestRomanNumberString(t *testing.T) {
	tests := []struct {
		r      RomanNumber
		expect string
	}{
		{
			map[RomanDigit]int{
				M:  1,
				CM: 1,
				IV: 1,
			},
			"MCMIV",
		},
		{
			map[RomanDigit]int{
				M:  1,
				CM: 1,
				L:  1,
				IV: 1,
			},
			"MCMLIV",
		},
		{
			map[RomanDigit]int{
				M:  2,
				X:  1,
				IV: 1,
			},
			"MMXIV",
		},
		{
			map[RomanDigit]int{
				C: 2,
				V: 1,
				I: 2,
			},
			"CCVII",
		},
		{
			map[RomanDigit]int{
				I:  2,
				IV: 1,
				V:  2,
				IX: 1,
				X:  2,
				XL: 1,
				L:  2,
				XC: 1,
				C:  2,
				CD: 1,
				D:  2,
				CM: 1,
				M:  2,
			},
			"MMCMDDCDCCXCLLXLXXIXVVIVII",
		},
	}

	for _, test := range tests {
		result := test.r.String()
		if result != test.expect {
			t.Errorf("Got %s, exepcted %s", result, test.expect)
		}
	}
}

func TestRomanNumberIsValid(t *testing.T) {
	tests := []struct {
		r      RomanNumber
		expect bool
	}{
		{
			map[RomanDigit]int{
				IV: 1,
				IX: 1,
				XL: 1,
				XC: 1,
				CD: 1,
				CM: 1,
			},
			true,
		},
		{
			map[RomanDigit]int{
				IV: 2,
				IX: 1,
				XL: 1,
				XC: 1,
				CD: 1,
				CM: 1,
			},
			false,
		},
		{
			map[RomanDigit]int{
				IV: 1,
				IX: 2,
				XL: 1,
				XC: 1,
				CD: 1,
				CM: 1,
			},
			false,
		},
		{
			map[RomanDigit]int{
				IV: 1,
				IX: 1,
				XL: 2,
				XC: 1,
				CD: 1,
				CM: 1,
			},
			false,
		},
		{
			map[RomanDigit]int{
				IV: 1,
				IX: 1,
				XL: 1,
				XC: 2,
				CD: 1,
				CM: 1,
			},
			false,
		},
		{
			map[RomanDigit]int{
				IV: 1,
				IX: 1,
				XL: 1,
				XC: 1,
				CD: 2,
				CM: 1,
			},
			false,
		},
		{
			map[RomanDigit]int{
				IV: 1,
				IX: 1,
				XL: 1,
				XC: 1,
				CD: 1,
				CM: 2,
			},
			false,
		},
	}

	for _, test := range tests {
		if test.r.IsValid() != test.expect {
			t.Errorf("IsValid(%s) != %v", test.r.String(), !test.expect)
		}
	}
}
