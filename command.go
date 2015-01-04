package nxt

import (
	"fmt"
)

type CommandType byte

const (
	DirectRequiresResponse CommandType = 0x00
	SystemRequiresResponse             = 0x01
	Reply                              = 0x02
	DirectNoResponse                   = 0x80
	SystemNoResponse                   = 0x81
)

type Command byte

type Telegram struct {
	Type    CommandType
	Command Command
	Message []byte
}

func (t Telegram) Bytes() []byte {
	commandInfo := []byte{byte(t.Type), byte(t.Command)}
	return append(commandInfo, t.Message...)
}

func (t Telegram) IsResponseRequired() bool {
	return t.Type == DirectRequiresResponse || t.Type == SystemRequiresResponse
}

func (t Telegram) String() string {
	return fmt.Sprintf("Type: 0x%02x, Command: 0x%02x, Message: %v", t.Type, t.Command, t.Message)
}

func newTelegramWithMessage(commandType CommandType, command Command, message []byte) *Telegram {
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

func newCommand(requiresResponse bool, responseRequiredCommandType CommandType, noResponseRequiredCommandType CommandType, command Command, message []byte) *Telegram {
	var commandType CommandType

	if requiresResponse {
		commandType = responseRequiredCommandType
	} else {
		commandType = noResponseRequiredCommandType
	}

	return newTelegramWithMessage(commandType, command, message)
}

func NewDirectCommand(requiresResponse bool, command Command, message []byte) *Telegram {
	return newCommand(requiresResponse, DirectRequiresResponse, DirectNoResponse, command, message)
}

func NewSystemCommand(requiresResponse bool, command Command, message []byte) *Telegram {
	return newCommand(requiresResponse, SystemRequiresResponse, SystemNoResponse, command, message)
}
