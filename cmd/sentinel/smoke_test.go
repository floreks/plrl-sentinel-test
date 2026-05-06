package main

import (
	"strings"
	"testing"
)

func buildStatusLine(pkg, test string, passed bool) string {
	result := "FAIL"
	if passed {
		result = "PASS"
	}
	return result + " " + pkg + "." + test
}

func TestBuildStatusLine(t *testing.T) {
	t.Parallel()

	if got, want := buildStatusLine("pkg/a", "TestOne", true), "PASS pkg/a.TestOne"; got != want {
		t.Fatalf("buildStatusLine(...) = %q, want %q", got, want)
	}
	if got, want := buildStatusLine("pkg/a", "TestTwo", false), "FAIL pkg/a.TestTwo"; got != want {
		t.Fatalf("buildStatusLine(...) = %q, want %q", got, want)
	}
}

func TestBuildStatusLineFormatting(t *testing.T) {
	t.Parallel()

	got := buildStatusLine("pkg/runner", "TestOutput", true)
	if !strings.HasPrefix(got, "PASS ") {
		t.Fatalf("expected PASS prefix, got %q", got)
	}
	if !strings.Contains(got, ".") {
		t.Fatalf("expected package.test format, got %q", got)
	}
}

func TestSkipExample(t *testing.T) {
	t.Parallel()
	t.Skip("example skipped test for runner reporting")
}
