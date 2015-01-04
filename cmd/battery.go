package cmd

import "github.com/tonyheupel/go-nxt"

func GetBatteryLevel(replyChannel chan *nxt.ReplyTelegram) *nxt.Command {
	return nxt.NewDirectCommand(0x0B, nil, replyChannel)
}

type GetBatteryLevelReply struct {
	*nxt.ReplyTelegram
	BatteryLevelMillivolts int
}

func ParseGetBatteryLevelReply(reply *nxt.ReplyTelegram) *GetBatteryLevelReply {
	return &GetBatteryLevelReply{
		ReplyTelegram: reply,
		BatteryLevelMillivolts: nxt.CalculateIntFromLSBAndMSB(reply.Message[0], reply.Message[1]),
	}
}
