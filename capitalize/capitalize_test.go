package capitalize

import "testing"

func TestCapitalize(t *testing.T) {
	want := "CAPITALIZE"
	got := Capitalize("capitalize")
	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}
