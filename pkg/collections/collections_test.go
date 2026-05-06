package collections

import (
	"maps"
	"slices"
	"testing"
)

func keysSorted[K ~string, V any](in map[K]V) []K {
	out := make([]K, 0, len(in))
	for k := range in {
		out = append(out, k)
	}
	slices.Sort(out)
	return out
}

func histogram[T comparable](items []T) map[T]int {
	out := make(map[T]int, len(items))
	for _, item := range items {
		out[item]++
	}
	return out
}

func TestKeysSorted(t *testing.T) {
	t.Parallel()

	in := map[string]int{"z": 1, "a": 2, "m": 3}
	want := []string{"a", "m", "z"}
	got := keysSorted(in)
	if !slices.Equal(got, want) {
		t.Fatalf("keysSorted(%v) = %v, want %v", in, got, want)
	}
}

func TestHistogram(t *testing.T) {
	t.Parallel()

	in := []string{"ok", "fail", "ok", "skip", "ok", "skip"}
	want := map[string]int{"ok": 3, "fail": 1, "skip": 2}
	got := histogram(in)
	if !maps.Equal(got, want) {
		t.Fatalf("histogram(%v) = %v, want %v", in, got, want)
	}
}

func TestHistogramEmpty(t *testing.T) {
	t.Parallel()

	if got := histogram([]int{}); len(got) != 0 {
		t.Fatalf("histogram(empty) = %v, want empty map", got)
	}
}
