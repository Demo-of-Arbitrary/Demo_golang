package numeral

import "testing"

func TestRomanNumeral(t *testing.T) {
	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"1 gets converted to I", 1, "I"},
		{"2 gets converted to II", 2, "II"},
		{"3 gets converted to III", 3, "III"},
		{"4 gets converted to IV", 4, "IV"},
		{"5 gets converted to V", 5, "V"},
	}

	for _, v := range cases {
		t.Run(v.Description, func(t *testing.T) {
			got := ConvertToRoman(v.Arabic)
			if got != v.Want {
				t.Errorf("expected %q, got %q", v.Want, got)
			}
		})
	}
}
