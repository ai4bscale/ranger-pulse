package version

import "testing"

func TestString(t *testing.T) {
	got := String()
	want := "dev (unknown)"
	if got != want {
		t.Errorf("String() = %q, want %q", got, want)
	}
}
