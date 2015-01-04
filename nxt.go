// The nxt package provides tools to allow one to control a
// Lego Mindstorms NXT 2.0.
package nxt

import (
	"fmt"
)

// NXT represents the thing that a caller interacts with
// to control an NXT brick.
type NXT struct {
	name           string
	port           string
	connection     Connection
	connected      bool
	CommandChannel chan *Command
}

// NewNXT creates a new NXT with the given name and
// will connect to the brick over Bluetooth using
// the serial port specified at the port argument.
func NewNXT(name string, port string) *NXT {
	return NewNXTUsingConnection(name, port, newBluetoothConnection(port))
}

// NewNXTUsingConnection creates a new NXT with the given name and
// a connection that implements the Connection interface.
// This method is good for testing when  passing in a test double for the connection.
func NewNXTUsingConnection(name string, port string, connection Connection) *NXT {
	return &NXT{
		name:       name,
		port:       port,
		connection: connection,
		connected:  false,
	}
}

// Name returns the friendly name of the NXT.
func (n NXT) Name() string { return n.name }

// Port returns the port that the NXT is to be connected to.
func (n NXT) Port() string { return n.port }

func (n NXT) String() string {
	return fmt.Sprintf("NXT \"%s\": %s", n.Name(), n.Port())
}

// Connect connects the NXT to the port and makes the NXT ready
// to receive commands.
func (n *NXT) Connect() error {
	n.CommandChannel = make(chan *Command)

	err := n.connection.Open()

	if err != nil {
		return err
	}

	n.connected = true

	go n.messageLoop()

	return nil
}

// Disconnect closes the connection to the port and
func (n *NXT) Disconnect() (err error) {
	// Closing the channel will signal the messageLoop with a "zero" message
	// and the messageLoop should stop listening for commands.
	close(n.CommandChannel)

	if n.connection != nil && n.connected {
		err = n.connection.Close()
		n.connected = false
	}

	return
}

// messageLoop is the message loop that listens for commands
func (n *NXT) messageLoop() {

	for {
		command, ok := <-n.CommandChannel
		if !ok {
			//Closed channel, stop listening
			return
		}

		n.connection.Write(command.Telegram.Bytes())

		if command.Telegram.IsResponseRequired() {
			command.ReplyChannel <- getReplyFromReader(n.connection)
		}
	}
}
