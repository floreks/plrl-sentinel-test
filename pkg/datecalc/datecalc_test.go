package datecalc

import (
	"testing"
	"time"
)

func beginningOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func addBusinessDays(start time.Time, days int) time.Time {
	cur := start
	step := 1
	if days < 0 {
		step = -1
	}

	for added := 0; added != days; {
		cur = cur.AddDate(0, 0, step)
		if cur.Weekday() == time.Saturday || cur.Weekday() == time.Sunday {
			continue
		}
		added += step
	}
	return cur
}

func TestBeginningOfDay(t *testing.T) {
	t.Parallel()

	loc, err := time.LoadLocation("Europe/Warsaw")
	if err != nil {
		t.Fatalf("LoadLocation: %v", err)
	}
	in := time.Date(2026, 5, 6, 15, 4, 9, 12345, loc)
	got := beginningOfDay(in)
	want := time.Date(2026, 5, 6, 0, 0, 0, 0, loc)
	if !got.Equal(want) {
		t.Fatalf("beginningOfDay(%v) = %v, want %v", in, got, want)
	}
}

func TestAddBusinessDays(t *testing.T) {
	t.Parallel()

	loc := time.UTC
	tests := []struct {
		name  string
		start time.Time
		days  int
		want  time.Time
	}{
		{
			name:  "forward_skips_weekend",
			start: time.Date(2026, 5, 8, 9, 0, 0, 0, loc),
			days:  1,
			want:  time.Date(2026, 5, 11, 9, 0, 0, 0, loc),
		},
		{
			name:  "forward_multiple_days",
			start: time.Date(2026, 5, 6, 9, 0, 0, 0, loc),
			days:  3,
			want:  time.Date(2026, 5, 11, 9, 0, 0, 0, loc),
		},
		{
			name:  "backward_skips_weekend",
			start: time.Date(2026, 5, 11, 9, 0, 0, 0, loc),
			days:  -1,
			want:  time.Date(2026, 5, 8, 9, 0, 0, 0, loc),
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := addBusinessDays(tc.start, tc.days); !got.Equal(tc.want) {
				t.Fatalf("addBusinessDays(%v, %d) = %v, want %v", tc.start, tc.days, got, tc.want)
			}
		})
	}
}
