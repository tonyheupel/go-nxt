package nxt

import (
	"testing"
)

func TestNewDirectCommandThatRequiresResponse(t *testing.T) {
	command := CommandCode(2)
	reply := make(chan *ReplyTelegram)
	dc := NewDirectCommand(command, nil, reply)

	if dc.Telegram.Type != DirectRequiresResponse {
		t.Errorf("Type of DirectRequiresResponse (0x%02x), but got 0x%02x instead", DirectRequiresResponse, dc.Telegram.Type)
	}

	if dc.Telegram.Command != command {
		t.Errorf("Expected Command of 0x%02x but got 0x%02x instead", command, dc.Telegram.Command)
	}
}
