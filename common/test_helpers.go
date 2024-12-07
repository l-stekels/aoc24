package common

import "testing"

// AssertEqualSlices checks if two slices are equal.
func AssertEqualSlices[T int | string](t *testing.T, a, b []T) {
	t.Helper()
	if len(a) != len(b) {
		t.Errorf("slices have different lengths: got %q, want %q", a, b)
	}
	for i := range a {
		if a[i] != b[i] {
			t.Errorf("slices differ at index %v: got %v, want %v", i, a[i], b[i])
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
		AssertEqualSlices[int](t, a[i], b[i])
	}
}
