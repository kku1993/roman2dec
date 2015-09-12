// roman2dec package implements a converter from Roman Numerals to decimal and
// vice versa.
package roman2dec

// RomanDigit type represents a digit in Roman Numeral.
type RomanDigit int

const (
	I  RomanDigit = 1
	IV            = 4
	V             = 5
	IX            = 9
	X             = 10
	XL            = 40
	L             = 50
	XC            = 90
	C             = 100
	CD            = 400
	D             = 500
	CM            = 900
	M             = 1000
)

func (d RomanDigit) String() string {
	switch d {
	case I:
		return "I"
	case IV:
		return "IV"
	case V:
		return "V"
	case IX:
		return "IX"
	case X:
		return "X"
	case XL:
		return "XL"
	case L:
		return "L"
	case XC:
		return "XC"
	case C:
		return "C"
	case CD:
		return "CD"
	case D:
		return "D"
	case CM:
		return "CM"
	default:
		return "M"
	}
}

var RomanDigitDict map[string]RomanDigit = map[string]RomanDigit{
	"I":  I,
	"IV": IV,
	"V":  V,
	"IX": IX,
	"X":  X,
	"XL": XL,
	"L":  L,
	"XC": XC,
	"C":  C,
	"CD": CD,
	"D":  D,
	"CM": CM,
	"M":  M,
}

// Sorted Roman Digits in descending order.
var SortedRomanDigits = []RomanDigit{M, CM, D, CD, C, XC, L, XL, X, IX, V, IV, I}

// RomanNumber represents a number in Roman Numeral.
type RomanNumber map[RomanDigit]int

func (r RomanNumber) String() string {
	output := ""
	for _, d := range SortedRomanDigits {
		if count, ok := r[d]; ok {
			for i := 0; i < count; i += 1 {
				output += d.String()
			}
		}
	}

	return output
}

// IsValid returns whether the given RomanNumber is valid.
// It checks that we should never have a subtractive notation appear more than
// once.
func (r RomanNumber) IsValid() bool {
	for digit, count := range r {
		switch digit {
		case IV:
			fallthrough
		case IX:
			fallthrough
		case XL:
			fallthrough
		case XC:
			fallthrough
		case CD:
			fallthrough
		case CM:
			if count > 1 {
				return false
			}
		}
	}

	return true
}
