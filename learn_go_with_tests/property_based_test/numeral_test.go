package numeral

import (
	"fmt"
	"testing"
)

var cases = []struct {
	Arabic int
	Roman  string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{7, "VII"},
	{8, "VIII"},
	{9, "IX"},
	{10, "X"},
	{11, "XI"},
	{14, "XIV"},
	{19, "XIX"},
	{24, "XXIV"},
	{40, "XL"},
	{47, "XLVII"},
	{50, "L"},
	{100, "C"},
	{90, "XC"},
	{400, "CD"},
	{500, "D"},
	{1000, "M"},
	{900, "CM"},
	{1984, "MCMLXXXIV"},
}

func TestRomanNumeral(t *testing.T) {
	for _, v := range cases {
		t.Run(fmt.Sprintf("%d converted to %q", v.Arabic, v.Roman), func(t *testing.T) {
			got := ConvertToRoman(v.Arabic)
			if got != v.Roman {
				t.Errorf("expected %q, got %q", v.Roman, got)
			}
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, v := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", v.Roman, v.Arabic), func(t *testing.T) {
			got := ConvertToArabic(v.Roman)

			if got != v.Arabic {
				t.Errorf("got %d, want %d", got, v.Arabic)
			}
		})
	}
}
