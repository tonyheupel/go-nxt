// The nxt package provides tools to allow one to control a
// Lego Mindstorms NXT 2.0.
package nxt

import (
	"fmt"
)

// NXT represents the thing that a caller interacts with
// to control an NXT brick.
type NXT struct {
	Name          string
	DevicePath    string
	connection    Connection
}

// NewNXT creates a new NXT with the given name and
// will connect to the brick over Bluetooth using
// the serial port specified at the devicePath argument.
func NewNXT(name string, devicePath string) *NXT {
	return NewNXTUsingConnection(name, devicePath, newBluetoothConnection(devicePath))
}

func NewNXTUsingConnection(name string, devicePath string, connection Connection) *NXT {
	return &NXT{
		Name:       name,
		DevicePath: devicePath,
		connection: connection,
	}
}
func (n *NXT) String() string {
	return fmt.Sprintf("NXT named %s, at %s", n.Name, n.DevicePath)
}

func (n *NXT) Connect() error {
	return n.connection.Open()
}

func (n *NXT) Disconnect() error {
	return n.connection.Close()
}

