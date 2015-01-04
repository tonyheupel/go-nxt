package nxt

type Command struct {
	Telegram     *Telegram
	ReplyChannel chan ReplyTelegram
}

func newCommand(responseRequiredCommandType CommandType, noResponseRequiredCommandType CommandType, command CommandCode, message []byte, replyChannel chan ReplyTelegram) *Command {
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

func NewDirectCommand(command CommandCode, message []byte, replyChannel chan ReplyTelegram) *Command {
	return newCommand(DirectRequiresResponse, DirectNoResponse, command, message, replyChannel)
}

func NewSystemCommand(command CommandCode, message []byte, replyChannel chan ReplyTelegram) *Command {
	return newCommand(SystemRequiresResponse, SystemNoResponse, command, message, replyChannel)
}
