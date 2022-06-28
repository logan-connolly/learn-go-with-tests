package pbt

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic int) string {
	var result strings.Builder
	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}
	return result.String()
}

func ConvertToArabic(roman string, count int) int {
	if roman == "" {
		return count
	}

	for _, numeral := range allRomanNumerals {
		if len(numeral.Symbol) == 2 {
			if len(roman) > 1 && roman[:2] == numeral.Symbol {
				return ConvertToArabic(roman[2:], count+numeral.Value)
			}
		} else {
			if roman[:1] == numeral.Symbol {
				return ConvertToArabic(roman[1:], count+numeral.Value)
			}
		}
	}

	return -1
}
