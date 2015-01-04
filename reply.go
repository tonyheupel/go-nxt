package nxt

import (
	"fmt"
	"io"
)

// ReplyStatus is the possible status codes returned as part of a ReplyTelegram's Status value
type ReplyStatus byte

const (
	Success                                   ReplyStatus = 0x00
	PendingCommunicationTransactionInProgress             = 0x20
	SpecifiedMailboxQueueIsEmpty                          = 0x40
	RequestFailed                                         = 0xBD
	UnknownCommandOpcode                                  = 0xBE
	InsanePacket                                          = 0xBF
	DataContainsOutOfRangeValues                          = 0xC0
	CommunicationBusError                                 = 0xDD
	NoFreeMemoryInCommunicationBuffer                     = 0xDE
	SpecifiedConnectionIsNotValid                         = 0xDF
	SpecifiedConnectionIsNotConfiguredOrBusy              = 0xE0
	NoActiveProgram                                       = 0xEC
	IllegalSizeSpecified                                  = 0xED
	IllegalMailboxQueueIDSpecified                        = 0xEE
	AttemptedToAccessInvalidFieldOfStructure              = 0xEF
	BadInputOrOutputSpecified                             = 0xF0
	InsufficientMemoryAvailable                           = 0xFB
	BadArguments                                          = 0xFF
)

// ReplyTelegram is the response to a command when the caller waits for the reply
type ReplyTelegram struct {
	*Telegram
	Status ReplyStatus
}

// Represent the ReplyTelegram as a string
func (r ReplyTelegram) String() string {
	return fmt.Sprintf("Status: 0x%02x, %v", r.Status, r.Telegram)
}

// IsSuccess returns true when the reply indicates a successful operation
func (r ReplyTelegram) IsSuccess() bool {
	return r.Status == Success
}

// NewReply creates a new ReplyTelegram for the given command, status, and an optional reply message; pass nil if there is no reply message for the command
func NewReply(replyForCommand CommandCode, status ReplyStatus, message []byte) *ReplyTelegram {
	return &ReplyTelegram{
		Telegram: &Telegram{
			Type:    Reply,
			Command: replyForCommand,
			Message: message,
		},
		Status: status,
	}
}

func newReplyFromBytes(replyBytes []byte, replyLength int) *ReplyTelegram {
	var replyMessage []byte
	if replyLength == len(replyBytes) {
		replyMessage = replyBytes
	} else if replyLength <= 3 {
		replyMessage = nil
	} else {
		replyMessage = replyBytes[3:replyLength]
	}

	// TODO: Make sure the first byte is 0x02 to indicate a reply

	return NewReply(CommandCode(replyBytes[1]), ReplyStatus(replyBytes[2]), replyMessage)
}

func getReplyFromReader(reader io.Reader) *ReplyTelegram {
	response := make([]byte, 64)

	numRead, _ := reader.Read(response)

	// TODO: Do not ignore the error here

	return newReplyFromBytes(response, numRead)
}
