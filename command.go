package nxt

// Command represents a command that is sent to the NXT and the
// optional reply.  The ReplyChannel is the channel that the
// NXT should send the reply to.
type Command struct {
	Telegram     *Telegram
	ReplyChannel chan *ReplyTelegram
}

func newCommand(responseRequiredCommandType CommandType, noResponseRequiredCommandType CommandType, command CommandCode, message []byte, replyChannel chan *ReplyTelegram) *Command {
	var commandType CommandType

	if replyChannel != nil {
		commandType = responseRequiredCommandType
	} else {
		commandType = noResponseRequiredCommandType
	}

	return &Command{
		Telegram:     newTelegramWithMessage(commandType, command, message),
		ReplyChannel: replyChannel,
	}
}

// NewDirectCommand creates a Command that can be sent to the NXT. To wait for a reply, pass in a replyChannel
// value; to not wait for a reply, pass in nil for the replyChannel argument.
// Direct commands are the most typical commands that are sent to the NXT.
func NewDirectCommand(command CommandCode, message []byte, replyChannel chan *ReplyTelegram) *Command {
	return newCommand(DirectRequiresResponse, DirectNoResponse, command, message, replyChannel)
}

// NewSystemCommand creates a Command that can be sent to the NXT. To wait for a reply, pass in a replyChannel
// value; to not wait for a reply, pass in nil for the replyChannel argument.
// System commands are not usually the most typical commands sent to the NXT.
func NewSystemCommand(command CommandCode, message []byte, replyChannel chan *ReplyTelegram) *Command {
	return newCommand(SystemRequiresResponse, SystemNoResponse, command, message, replyChannel)
}
