package hsd

import "testing"

func TestStringDistance(t *testing.T) {
	got := StringDistance("foo", "bar")
	want := 3
	if got != want {
		t.Fatalf("expected: %v, actual: %v",want, got)
	}
}
