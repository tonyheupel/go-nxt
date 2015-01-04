package directcommands

func NewStartProgramCommand(filename string, replyChannel chan ReplyTelegram) *Command 
	file := append([]byte(filename), 0) // null-terminated string

	return NewDirectCommand(0x00, file, replyChannel)
}

func NewStopProgramCommand(replyChannel chan ReplyTelegram) *Command {
	return NewDirectCommand(0x01, nil, replyChannel)
}

func NewGetCurrentProgramNameCommand(replyChannel chan ReplyTelegram) *Command {
	return NewDirectCommand(0x11, nil, replyChannel)
}

type GetCurrentProgramNameReply struct {
	ReplyTelegram
	Filename string
}

func ParseGetCurrentProgramNameReply(reply ReplyTelegram) *GetCurrentProgramNameReply {
	return &GetCurrentProgramNameReply{
		ReplyTelegram: reply,
		Filename: string(reply.Message),
	}
}
