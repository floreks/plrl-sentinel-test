package parser

import (
	"strings"
	"testing"
)

func parseKV(line string) (key string, value string, ok bool) {
	idx := strings.IndexByte(line, '=')
	if idx <= 0 || idx == len(line)-1 {
		return "", "", false
	}
	return strings.TrimSpace(line[:idx]), strings.TrimSpace(line[idx+1:]), true
}

func TestParseKV(t *testing.T) {
	t.Parallel()

	tests := []struct {
		line     string
		wantKey  string
		wantVal  string
		wantGood bool
	}{
		{line: "suite = sentinel", wantKey: "suite", wantVal: "sentinel", wantGood: true},
		{line: "status=ok", wantKey: "status", wantVal: "ok", wantGood: true},
		{line: "=missing_key", wantGood: false},
		{line: "missing_value=", wantGood: false},
		{line: "no_separator", wantGood: false},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.line, func(t *testing.T) {
			t.Parallel()
			key, val, ok := parseKV(tc.line)
			if ok != tc.wantGood {
				t.Fatalf("parseKV(%q) ok=%v, want %v", tc.line, ok, tc.wantGood)
			}
			if ok {
				if key != tc.wantKey || val != tc.wantVal {
					t.Fatalf("parseKV(%q) = (%q, %q), want (%q, %q)", tc.line, key, val, tc.wantKey, tc.wantVal)
				}
			}
		})
	}
}
