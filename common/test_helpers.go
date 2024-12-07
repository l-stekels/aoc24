package common

import "testing"

// AssertEqualIntSlices checks if two slices of integers are equal.
func AssertEqualIntSlices(t *testing.T, a, b []int) {
	t.Helper()
	if len(a) != len(b) {
		t.Errorf("slices have different lengths: got %q, want %q", a, b)
	}
	for i := range a {
		if a[i] != b[i] {
			t.Errorf("slices differ at index %d: got %d, want %d", i, a[i], b[i])
		}
	}
}

// AssertEqual2DIntSlices checks if two 2D slices of integers are equal.
func AssertEqual2DIntSlices(t *testing.T, a, b [][]int) {
	t.Helper()
	if len(a) != len(b) {
		t.Errorf("slices have different lengths: got %q, want %q", a, b)
	}
	for i := range a {
		AssertEqualIntSlices(t, a[i], b[i])
	}
}
