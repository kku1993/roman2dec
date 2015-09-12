package roman2dec

import (
	"fmt"
	"strings"
)

// RomanString2Number converts a Roman Number presented as a string into a
// value of Roman Number type. An error is returned if the input cannot be
// parsed.
func RomanString2Number(s string) (RomanNumber, error) {
	// Convert input to all uppercase.
	s = strings.ToUpper(s)

	romanNumber := make(RomanNumber)

	// Set previousDigit to the largest RomanDigit possible.
	previousDigit := RomanDigit(M)

	for i := 0; i < len(s); i += 1 {
		var digit RomanDigit
		var foundDigit bool

		// Check for subtractive notation.
		if i < len(s)-1 {
			if digit, foundDigit = RomanDigitDict[s[i:i+2]]; foundDigit {
				i += 1
			}
		}

		if !foundDigit {
			// Try to parse single digit.
			if digit, foundDigit = RomanDigitDict[s[i:i+1]]; !foundDigit {
				return nil, fmt.Errorf("cannot parse %s", s)
			}
		}

		// Make sure the digits are given in descending order.
		if previousDigit < digit {
			return nil, fmt.Errorf("%s is not a valid Roman Numeral", s)
		}
		previousDigit = digit

		if _, ok := romanNumber[digit]; ok {
			romanNumber[digit] += 1
		} else {
			romanNumber[digit] = 1
		}
	}

	// Check the RomanNumber is valid.
	if !romanNumber.IsValid() {
		return nil, fmt.Errorf("%s is not a valid Roman Numeral", s)
	}

	return romanNumber, nil
}

// Roman2Dec converts a Roman Numeral number to decimal number.
func Roman2Dec(r RomanNumber) int {
	sum := 0
	for digit, count := range r {
		sum += int(digit) * count
	}
	return sum
}

// Dec2Roman converts a decimal number to Roman Numeral.
func Dec2Roman(n int) (RomanNumber, error) {
	if n <= 0 {
		// The Romans didn't need 0 or negative numbers.
		return nil, fmt.Errorf("%d does not have a Roman Numeral representation", n)
	}

	romanNumber := make(RomanNumber)

	for _, digit := range SortedRomanDigits {
		count := n / int(digit)
		n -= count * int(digit)

		romanNumber[digit] = count
	}

	return romanNumber, nil
}
