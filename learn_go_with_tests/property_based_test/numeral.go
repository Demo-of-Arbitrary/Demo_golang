package numeral

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

type RomanNumerals []RomanNumeral

func (n RomanNumerals) ValueOf(symbols ...byte) int {
	symbol := string(symbols)
	for _, v := range n {
		if v.Symbol == symbol {
			return v.Value
		}
	}
	return 0
}

var allRomanNumeral = RomanNumerals{
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
	for _, r := range allRomanNumeral {
		for arabic >= r.Value {
			result.WriteString(r.Symbol)
			arabic -= r.Value
		}
	}
	return result.String()
}

func ConvertToArabic(roman string) int {
	total := 0
	for i := 0; i < len(roman); i++ {
		symbol := roman[i]
		if couldBeSubstractive(i, symbol, roman) {
			nextSymbol := roman[i+1]

			if value := allRomanNumeral.ValueOf(symbol, nextSymbol); value > 0 {
				total += value
				i++
			} else {
				total++
			}
		} else {
			total += allRomanNumeral.ValueOf(symbol)
		}
	}
	return total
}

func couldBeSubstractive(index int, currentSymbol uint8, roman string) bool {
	return index+1 < len(roman) && currentSymbol == 'I'
}
