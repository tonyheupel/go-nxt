package nxt

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

type ReplyTelegram struct {
	*Telegram
	Status ReplyStatus
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

func NewReply(replyForCommand Command, status ReplyStatus, message []byte) *ReplyTelegram {
	return &ReplyTelegram{
		Telegram: &Telegram{
			Type:    Reply,
			Command: replyForCommand,
			Message: message,
		},
		Status: status,
	}
}
