package ifswitch

import "testing"

func TestSheOrHe(t *testing.T) {
	got := SheOrHe("maLe")
	want := "He"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
