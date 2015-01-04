package nxt

import "testing"

func TestCalulateIntFromLSBAndMSB(t *testing.T) {
	var expected int = 38726 // 1001 0111 0100 0110  or  0x9746 or [151, 70] Little endian is [70, 151]

	actual := calculateIntFromLSBAndMSB(70, 151)

	if expected != actual {
		t.Errorf("Expected: %d, Actual: %d", expected, actual)
	}

}
