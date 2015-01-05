package nxt

import "fmt"

// CommandType represents the different type of telegrams
// that can be sent to the NXT.
type CommandType byte

const (
	DirectRequiresResponse CommandType = 0x00
	SystemRequiresResponse             = 0x01
	Reply                              = 0x02
	DirectNoResponse                   = 0x80
	SystemNoResponse                   = 0x81
)

func (c CommandType) String() string {
	CommandTypeNames := map[CommandType]string{
		DirectRequiresResponse: "Direct requiring response",
		SystemRequiresResponse: "System requiring response",
		Reply:            "Reply",
		DirectNoResponse: "Direct",
		SystemNoResponse: "System",
	}

	return CommandTypeNames[c]
}

// CommandCode is the code used to represent the command
// to send to the NXT in a Telegram.
type CommandCode byte

// Telegram is the basic communication structure used to
// interact with the NXT.
type Telegram struct {
	Type    CommandType
	Command CommandCode
	Message []byte
}

// Bytes returns the slice of bytes that represents the Telegram in a
// format that can be understood by the NXT.
func (t Telegram) Bytes() []byte {
	commandInfo := []byte{byte(t.Type), byte(t.Command)}
	return append(commandInfo, t.Message...)
}

// IsResponseRequired returns true when the Telegram is going
// to wait for the reply from the NXT.  Only require a response
// when necessary, as it can add up to 60ms to the command time
// (per the NXT documentation).
func (t Telegram) IsResponseRequired() bool {
	return t.Type == DirectRequiresResponse || t.Type == SystemRequiresResponse
}

// String represents the Telegram as a string.
func (t Telegram) String() string {
	return fmt.Sprintf("Type: %v, Command: 0x%02x, Message: %v", t.Type, t.Command, t.Message)
}

// newTelegramWithMessage returns a Telegram with the given values.  If a telegram does
// not require a message, then pass in nil for the message.
func newTelegramWithMessage(commandType CommandType, command CommandCode, message []byte) *Telegram {
	const MAX_TELEGRAM_BYTES = 64
	const MAX_MESSAGE_BYTES = MAX_TELEGRAM_BYTES - 2 // remove Type and Command

	if message == nil {
		message = make([]byte, MAX_MESSAGE_BYTES)
	}

	return &Telegram{
		Type:    commandType,
		Command: command,
		Message: message,
	}
}
