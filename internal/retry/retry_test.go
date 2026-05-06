package retry

import (
	"errors"
	"testing"
)

func runWithRetries(maxAttempts int, fn func() error) (int, error) {
	if maxAttempts <= 0 {
		return 0, errors.New("maxAttempts must be > 0")
	}
	var err error
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		err = fn()
		if err == nil {
			return attempt, nil
		}
	}
	return maxAttempts, err
}

func TestRunWithRetriesSuccessEventually(t *testing.T) {
	t.Parallel()

	attempts := 0
	gotAttempts, err := runWithRetries(5, func() error {
		attempts++
		if attempts < 3 {
			return errors.New("transient")
		}
		return nil
	})
	if err != nil {
		t.Fatalf("runWithRetries returned error: %v", err)
	}
	if gotAttempts != 3 {
		t.Fatalf("attempts = %d, want 3", gotAttempts)
	}
}

func TestRunWithRetriesFailure(t *testing.T) {
	t.Parallel()

	gotAttempts, err := runWithRetries(3, func() error {
		return errors.New("always failing")
	})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if gotAttempts != 3 {
		t.Fatalf("attempts = %d, want 3", gotAttempts)
	}
}

func TestRunWithRetriesBadConfig(t *testing.T) {
	t.Parallel()

	if _, err := runWithRetries(0, func() error { return nil }); err == nil {
		t.Fatalf("expected validation error")
	}
}
