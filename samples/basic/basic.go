package main

import (
	"fmt"
	"github.com/tonyheupel/go-nxt"
	"time"
)

func main() {
	n := nxt.NewNXT("heupel-home-bot", "/dev/tty.NXT-DevB")

	fmt.Println(n)
	err := n.Connect()

	if err != nil {
		fmt.Println("Could not connect:", err)
		return
	}

	fmt.Println("Connected!")

	friendly(n)
	time.Sleep(2 * time.Second)
	channels(n)

	n.Disconnect()
}

func friendly(n *nxt.NXT) {

	// Normally use StartProgram but we want to see the name of the running program
	// so we need to wait
	startProgramReply, err := n.StartProgramSync("DREW.rxe")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Reply from StartProgram:", startProgramReply)

	runningProgram, err := n.GetCurrentProgramName()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Current running program:", runningProgram)

	time.Sleep(3 * time.Second) // Wait 3 seconds before trying to stop

	fmt.Println("Stopping running program...")
	stopProgramReply, err := n.StopProgramSync()

	if err != nil {
		fmt.Println(err)
		return
	}

	if stopProgramReply.IsSuccess() {
		fmt.Println("Stopped running program successfully!")
	} else {
		fmt.Println("Was unable to stop the program.")
	}

	batteryMillivolts, err := n.GetBatteryLevelMillivolts()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Battery level (mv):", batteryMillivolts)
}

func channels(n *nxt.NXT) {
	reply := make(chan *nxt.ReplyTelegram)

	n.CommandChannel <- nxt.StartProgram("DREW.rxe", reply)
	fmt.Println("Reply from StartProgram:", <-reply)

	// Normally would pass in nil for the reply channel and not wait,
	//but we want to see the name of the running program so we need to wait

	n.CommandChannel <- nxt.GetCurrentProgramName(reply)
	runningProgramReply := nxt.ParseGetCurrentProgramNameReply(<-reply)
	fmt.Println("Current running program:", runningProgramReply.Filename)

	time.Sleep(3 * time.Second) // Wait 3 seconds

	fmt.Println("Stopping running program...")
	n.CommandChannel <- nxt.StopProgram(reply)

	stopProgramReply := <-reply

	if stopProgramReply.IsSuccess() {
		fmt.Println("Stopped running program successfully!")
	} else {
		fmt.Println("Was unable to stop the program.")
	}

	n.CommandChannel <- nxt.GetBatteryLevel(reply)
	batteryLevelReply := nxt.ParseGetBatteryLevelReply(<-reply)

	if batteryLevelReply.IsSuccess() {
		fmt.Println("Battery level (mv):", batteryLevelReply.BatteryLevelMillivolts)
	} else {
		fmt.Println("Was unable to get the current battery level")
	}
}
