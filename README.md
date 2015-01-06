# NXT
A Go package for controlling Lego Mindstorms NXT 2.0

## Installation
Either install it from the command-line:

```shell
$ go install github.com/tonyheupel/go-nxt
```

or add as a dependency in your code files:
```go
import "github.com/tonyheupel/go-nxt"
```
```shell
$ go get
```

## Usage
The most basic usage can be done using the NXT class by connecting your NXT 2.0
brick over Bluetooth.

There are two main ways to interact with an NXT instance:

1. Traditional method calls on the NXT instance, checking for errors
2. Interacting with the CommandChannel on the NXT directly.

Interacting with the channels gives you lower-level access and requires a little
more work, but if you prefer using channels as your communication paradigm, it's
there for you.  The NXT commands are all defined as Command structs and can
be passed around, so if you are managing multiple bots at once and want to
issue the same command to multiple bots, this approach may make more sense.

The example below has one method that first uses the more traditional style
interaction, and the second uses channels. They should produce the same output
and have the same effect on the bot.

The brick is connected at ``` /dev/tty.NXT-DevB ```.

```go
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

	n.CommandChannel <- nxt.GetBatteryLevel(reply)
	batteryLevelReply := nxt.ParseGetBatteryLevelReply(<-reply)

	if batteryLevelReply.IsSuccess() {
		fmt.Println("Battery level (mv):", batteryLevelReply.BatteryLevelMillivolts)
	} else {
		fmt.Println("Was unable to get the current battery level")
	}
}
