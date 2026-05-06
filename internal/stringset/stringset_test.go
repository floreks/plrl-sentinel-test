package stringset

import (
	"slices"
	"strings"
	"testing"
)

func normalize(items []string) []string {
	out := make([]string, 0, len(items))
	for _, item := range items {
		trimmed := strings.TrimSpace(strings.ToLower(item))
		if trimmed == "" {
			continue
		}
		out = append(out, trimmed)
	}
	slices.Sort(out)
	return slices.Compact(out)
}

func contains(set []string, target string) bool {
	for _, item := range set {
		if item == target {
			return true
		}
	}
	return false
}

func TestNormalize(t *testing.T) {
	t.Parallel()

	in := []string{"  Alpha", "beta", "ALPHA", "", "  ", "Beta", "gamma"}
	want := []string{"alpha", "beta", "gamma"}

	got := normalize(in)
	if !slices.Equal(got, want) {
		t.Fatalf("normalize(%v) = %v, want %v", in, got, want)
	}
}

func TestContains(t *testing.T) {
	t.Parallel()

	set := []string{"alpha", "beta", "gamma"}
	if !contains(set, "beta") {
		t.Fatalf("contains(%v, beta) = false, want true", set)
	}
	if contains(set, "delta") {
		t.Fatalf("contains(%v, delta) = true, want false", set)
	}
}

func TestContainsParallelSubtests(t *testing.T) {
	t.Parallel()

	set := []string{"one", "two", "three"}
	tests := []struct {
		name string
		key  string
		want bool
	}{
		{name: "present", key: "two", want: true},
		{name: "missing", key: "five", want: false},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := contains(set, tc.key); got != tc.want {
				t.Fatalf("contains(%v, %q) = %v, want %v", set, tc.key, got, tc.want)
			}
		})
	}
}
