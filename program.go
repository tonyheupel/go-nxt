package nxt

import (
	"fmt"
)

func startProgram(n *NXT, requireResponse bool, filename string) (reply *Telegram, err error) {
	return nil, nil
}

func (n NXT) StartProgram(filename string) {
	file := append([]byte(filename), 0) // null-terminated string

	telegram := NewDirectCommand(true, 0x00, file)

	command := telegram.Bytes()

	_, err := n.connection.Write(command)

	if err != nil {
		fmt.Println(err)
	}

	reply := getReplyFromReader(n.connection)

	fmt.Println("Reply:", reply)
}


func (n NXT) GetCurrentProgramName() string {
	telegram := NewDirectCommand(true, 0x11, nil)

	command := telegram.Bytes()

	_, err := n.connection.Write(command)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Called GetCurrentProgramName for", n.Name)

	reply := getReplyFromReader(n.connection)

	fmt.Println("Reply:", reply)

	return string(reply.Message)
}
