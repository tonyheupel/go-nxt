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
		fmt.Println(err)
		return
	}

	fmt.Println("Connected!")

	n.StartProgramSync("DREW.rxe")

	// Normally use StartProgram but we want to see the name of the running program
	// so we need to wait

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

	n.Disconnect()
}
