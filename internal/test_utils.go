package internal

import "testing"

func AssertEqual[K Number](t *testing.T, got, want K) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertNotEqual[K Number](t *testing.T, got, want K) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want %v", got)
	}
}

func AssertEqualSlices[K Number](t *testing.T, got, want []K) {
	t.Helper()

	if len(got) != len(want) {
		t.Errorf("Array length not equal, got %v of size %v, want %v of size %v", got, len(got), want, len(want))
	}

	for x := range got {
		if got[x] != want[x] {
			t.Errorf("Entries not equal at index %v, got %v, want %v", x, got, want)
		}
	}
}

func AssertNotEqualSlices[K Number](t *testing.T, got, want []K) {
	if len(got) == len(want) {
		t.Errorf("Array length equal, got %v, want %v", got, want)
	}

	for x := range got {
		if got[x] == want[x] {
			t.Errorf("Entries equal at index %v, got %v, want %v", x, got, want)
		}
	}
}
