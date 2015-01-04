package directcommands

func NewGetBatteryLevelCommand(replyChannel chan ReplyTelegram) *Command {
	return NewDirectCommand(0x0B, nil, replyChannel)
}

type GetBatteryLevelReply struct {
	ReplyTelegram
	BatteryLevelMillivolts int
}

func ParseGetBatteryLevelReply(reply ReplyTelegram) *GetBatteryLevelReply {
	return &GetBatteryLevelReply{
		ReplyTelegram: reply,
		BatteryLevelMillivolts: calculateIntFromLSBAndMSB(reply.Message[0], reply.Message[1])
	}
}
