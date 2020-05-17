package property_based_tests

import "testing"

func TestRomanNumerals(t *testing.T) {
	got := ConvertToRoman(1)
	want := "1"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}