package nxt

import (
	"testing"
)

func TestNewDirectCommandThatRequiresResponse(t *testing.T) {
	command := Command(2)
	dc := NewDirectCommand(true, command, nil)

	if dc.Type != DirectRequiresResponse {
		t.Errorf("Type of DirectRequiresResponse (0x%02x), but got 0x%02x instead", DirectRequiresResponse, dc.Type)
	}

	if dc.Command != command {
		t.Errorf("Expected Command of 0x%02x but got 0x%02x instead", command, dc.Command)
	}
}
