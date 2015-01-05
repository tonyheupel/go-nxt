package nxt

import "fmt"

// GetBatteryLevel creates a command that can be sent to the NXT
// to retrieve the battery level on the device, represented
// in millivolts.
func GetBatteryLevel(replyChannel chan *ReplyTelegram) *Command {
	return NewDirectCommand(0x0B, nil, replyChannel)
}

// GetBatteryLevelReply is the reply telegram for the GetBatteryLevel command.
// The battery level is accessed via the BatteryLevelMillivolts member.
type GetBatteryLevelReply struct {
	*ReplyTelegram
	BatteryLevelMillivolts int
}

// ParseGetBatteryLevelReply takes a raw ReplyTelegram and converts it to
// a GetBatteryLevelReply.
func ParseGetBatteryLevelReply(reply *ReplyTelegram) *GetBatteryLevelReply {
	return &GetBatteryLevelReply{
		ReplyTelegram:          reply,
		BatteryLevelMillivolts: calculateIntFromLSBAndMSB(reply.Message[0], reply.Message[1]),
	}
}

// GetBatteryLevelMillivolts gets the battery level of the NXT device,
// represented in millivolts.
func (n NXT) GetBatteryLevelMillivolts() (int, error) {
	reply := make(chan *ReplyTelegram)

	n.CommandChannel <- GetBatteryLevel(reply)
	batteryLevelReply := ParseGetBatteryLevelReply(<-reply)

	if !batteryLevelReply.IsSuccess() {
		return 0, fmt.Errorf("%v", batteryLevelReply.Status)
	}

	return batteryLevelReply.BatteryLevelMillivolts, nil

}
