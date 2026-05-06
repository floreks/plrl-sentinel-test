package pipeline

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func writeAndRead(path string, content string) (string, error) {
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		return "", err
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(raw), nil
}

func resolveMode() string {
	mode := os.Getenv("PIPELINE_MODE")
	if mode == "" {
		return "default"
	}
	return strings.ToLower(mode)
}

func TestWriteAndRead(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	path := filepath.Join(dir, "artifact.txt")
	want := "job=prepare\nstatus=ok\n"
	got, err := writeAndRead(path, want)
	if err != nil {
		t.Fatalf("writeAndRead error: %v", err)
	}
	if got != want {
		t.Fatalf("writeAndRead(%s) = %q, want %q", path, got, want)
	}
}

func TestResolveMode(t *testing.T) {
	t.Setenv("PIPELINE_MODE", "")
	if got := resolveMode(); got != "default" {
		t.Fatalf("resolveMode() = %q, want %q", got, "default")
	}

	t.Setenv("PIPELINE_MODE", "CI")
	if got := resolveMode(); got != "ci" {
		t.Fatalf("resolveMode() = %q, want %q", got, "ci")
	}
}

func TestResolveModeParallel(t *testing.T) {
	modes := []string{"LOCAL", "Staging", "Prod"}
	for _, mode := range modes {
		mode := mode
		t.Run(mode, func(t *testing.T) {
			t.Setenv("PIPELINE_MODE", mode)
			if got := resolveMode(); got != strings.ToLower(mode) {
				t.Fatalf("resolveMode() = %q, want %q", got, strings.ToLower(mode))
			}
		})
	}
}
