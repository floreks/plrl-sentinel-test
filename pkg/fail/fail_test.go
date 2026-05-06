package fail

import "testing"

func TestFail(t *testing.T) {
	t.Errorf("This test should fail")
}
