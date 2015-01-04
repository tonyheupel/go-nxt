package main

import (
	"fmt"
	"github.com/tonyheupel/go-nxt"
	"github.com/tonyheupel/go-nxt/cmd"
	"time"
)

func main() {
	n := nxt.NewNXT("heupel-home-bot", "/dev/tty.NXT-DevB")

	fmt.Println(n)
	reply := make(chan *nxt.ReplyTelegram)
	err := n.Connect()

	if err != nil {
		fmt.Println("Could not connect:", err)
		return
	}

	fmt.Println("Connected!")

	n.CommandChannel <- cmd.StartProgram("DREW.rxe", reply)
	fmt.Println("Reply from StartProgram:", <-reply)

	// Normally would pass in nil for the reply channel and not wait,
	//but we want to see the name of the running program so we need to wait

	n.CommandChannel <- cmd.GetCurrentProgramName(reply)
	runningProgramReply := cmd.ParseGetCurrentProgramNameReply(<-reply)
	fmt.Println("Current running program:", runningProgramReply.Filename)

	time.Sleep(3 * time.Second) // Wait 3 seconds

	fmt.Println("Stopping running program...")
	n.CommandChannel <- cmd.StopProgram(reply)

	stopProgramReply := <-reply

	if stopProgramReply.IsSuccess() {
		fmt.Println("Stopped running program successfully!")
	} else {
		fmt.Println("Was unable to stop the program.")
	}

	n.CommandChannel <- cmd.GetBatteryLevel(reply)
	batteryLevelReply := cmd.ParseGetBatteryLevelReply(<-reply)

	if batteryLevelReply.IsSuccess() {
		fmt.Println("Battery level (mv):", batteryLevelReply.BatteryLevelMillivolts)
	} else {
		fmt.Println("Was unable to get the current battery level")
	}

	n.Disconnect()
}
