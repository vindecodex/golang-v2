package ifswitch

import "testing"

func TestCountries2LetterCode(t *testing.T) {
	got := Countries2LetterCode("japAn")
	want := "JP"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
