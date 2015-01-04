package nxt

import "fmt"

func GetBatteryLevel(replyChannel chan *ReplyTelegram) *Command {
	return NewDirectCommand(0x0B, nil, replyChannel)
}

type GetBatteryLevelReply struct {
	*ReplyTelegram
	BatteryLevelMillivolts int
}

func ParseGetBatteryLevelReply(reply *ReplyTelegram) *GetBatteryLevelReply {
	return &GetBatteryLevelReply{
		ReplyTelegram:          reply,
		BatteryLevelMillivolts: calculateIntFromLSBAndMSB(reply.Message[0], reply.Message[1]),
	}
}

func (n NXT) GetBatteryLevelMillivolts() (int, error) {
	reply := make(chan *ReplyTelegram)

	n.CommandChannel <- GetBatteryLevel(reply)
	batteryLevelReply := ParseGetBatteryLevelReply(<-reply)

	if !batteryLevelReply.IsSuccess() {
		return 0, fmt.Errorf("Error getting battery level: %v", batteryLevelReply)
	}

	return batteryLevelReply.BatteryLevelMillivolts, nil

}
