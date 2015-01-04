package main

import (
	"fmt"
	"github.com/tonyheupel/go-nxt"
	"github.com/tonyheupel/go-nxt/cmd"
	"time"
)

func main() {
	device := nxt.NewNXT("heupel-home-bot", "/dev/tty.NXT-DevB")

	fmt.Println(device)
	//traditional(device)
	channels(device)
}

func channels(n *nxt.NXT) {
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

//func traditional(n *nxt.NXT) {
//	err := n.Connect()
//
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	fmt.Println("Connected!")
//
//	n.StartProgramSync("DREW.rxe")
//
//	// Normally use StartProgram but we want to see the name of the running program
//	// so we need to wait
//
//	runningProgram, err := n.GetCurrentProgramName()
//
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	fmt.Println("Current running program:", runningProgram)
//
//	fmt.Println("Stopping running program...")
//	stopProgramReply, err := n.StopProgramSync()
//
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	if stopProgramReply.IsSuccess() {
//		fmt.Println("Stopped running program successfully!")
//	} else {
//		fmt.Println("Was unable to stop the program.")
//	}
//
//	batteryMillivolts, err := n.GetBatteryLevelMillivolts()
//
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	fmt.Println("Battery level (mv):", batteryMillivolts)
//
//	n.Disconnect()
//}
