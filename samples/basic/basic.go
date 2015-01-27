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

	// Use a more traditional-looking method/check-for-error style
	methodStyle(n)

	// Pause in between styles to ensure the old commands are done executing
	time.Sleep(2 * time.Second)

	// Use the raw channels style
	channelStyle(n)

	n.Disconnect()
}

func methodStyle(n *nxt.NXT) {
	// Normally use StartProgram but we want to see the name of the running program
	// so we need to wait
	startProgramReply, err := n.StartProgramSync("DREW.rxe")

	if err != nil {
		fmt.Println("Error starting a program:", err)
	}

	fmt.Println("Reply from StartProgram:", startProgramReply)

	runningProgram, err := n.GetCurrentProgramName()

	if err != nil {
		fmt.Println("Error getting current program name:", err)
	} else {
		fmt.Println("Current running program:", runningProgram)
	}

	time.Sleep(3 * time.Second) // Wait 3 seconds before trying to stop

	fmt.Println("Stopping running program...")
	_, err = n.StopProgramSync()

	if err != nil {
		fmt.Println("Error stopping the running program:", err)
	}

	playSoundFileReply, err := n.PlaySoundFileSync("Green.rso", false)

	if err != nil {
		fmt.Println("Error playing the sound file \"Green.rso\":", err)
	}

	fmt.Println("Reply from PlaySoundFile:", playSoundFileReply)

	fmt.Println("Playing Convert A for 3 seconds...")
	playToneReply, err := n.PlayToneSync(440, 3000)

	if err != nil {
		fmt.Println("Error playing the tone:", err)
	}

	fmt.Println("Reply from PlayTone:", playToneReply)

	batteryMillivolts, err := n.GetBatteryLevelMillivolts()

	if err != nil {
		fmt.Println("Error getting the battery level:", err)
	} else {
		fmt.Println("Battery level (mv):", batteryMillivolts)
	}
}

func channelStyle(n *nxt.NXT) {
	// All reply messages will be sent to this channel
	reply := make(chan *nxt.ReplyTelegram)

	// Normally would pass in nil for the reply channel and not wait,
	//but we want to see the name of the running program so we need to wait
	n.CommandChannel <- nxt.StartProgram("DREW.rxe", reply)
	fmt.Println("Reply from StartProgram:", <-reply)

	n.CommandChannel <- nxt.GetCurrentProgramName(reply)
	runningProgramReply := nxt.ParseGetCurrentProgramNameReply(<-reply)
	fmt.Println("Current running program:", runningProgramReply.Filename)

	time.Sleep(3 * time.Second) // Wait 3 seconds before trying to stop

	fmt.Println("Stopping running program...")
	n.CommandChannel <- nxt.StopProgram(reply)

	stopProgramReply := <-reply

	if stopProgramReply.IsSuccess() {
		fmt.Println("Stopped running program successfully!")
	} else {
		fmt.Println("Was unable to stop the program.")
	}

	fmt.Println("Playing sound file \"Green\"...")
	n.CommandChannel <- nxt.PlaySoundFile("Green.rso", false, reply)

	playSoundFileReply := <-reply

	if playSoundFileReply.IsSuccess() {
		fmt.Println("Played sound file successfully!")
	} else {
		fmt.Println("Was unable to play the sound file:", playSoundFileReply)
	}

	fmt.Println("Playing Concert A (440Hz) for 3 seconds...")
	n.CommandChannel <- nxt.PlayTone(440, 3000, reply)

	playToneReply := <-reply

	if playToneReply.IsSuccess() {
		fmt.Println("Played Concert A!")
	} else {
		fmt.Println("Was unable to play the tone:", playToneReply)
	}

	n.CommandChannel <- nxt.GetBatteryLevel(reply)
	batteryLevelReply := nxt.ParseGetBatteryLevelReply(<-reply)

	if batteryLevelReply.IsSuccess() {
		fmt.Println("Battery level (mv):", batteryLevelReply.BatteryLevelMillivolts)
	} else {
		fmt.Println("Was unable to get the current battery level")
	}
}
