package nxt

import (
	"fmt"
	"testing"
)

func TestNXTStringer(t *testing.T) {
	name, devicePath := "testingbot", "/dev/tty-foobar"
	var n fmt.Stringer
	n = NewNXT(name, devicePath)

	if n == nil {
		t.Errorf("Calling NewNXT should not result in a nil NXT instance.")
		return
	}

	expected, actual := fmt.Sprintf("NXT \"%s\": %s", name, devicePath), n.String()
	if actual != expected {
		t.Errorf("Expected NXT.String to return \"%s\", but got \"%s\" instead", expected, actual)
	}
}
