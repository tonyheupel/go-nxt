package nxt

import (
	"fmt"
)

func startProgram(n *NXT, requireResponse bool, filename string) (reply *Telegram, err error) {
	return nil, nil
}

func (n *NXT) StartProgram(filename string) {
	fmt.Println("Called StartProgram for filename", filename)
}
