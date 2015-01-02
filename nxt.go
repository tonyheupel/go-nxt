// The nxt package provides tools to allow one to control a
// Lego Mindstorms NXT 2.0.
package nxt

import (
	"fmt"
)

type NXT struct {
	Name       string
	DevicePath string
	// Communication some-kind-of-io-interface-around-bluetooth
}

func NewNXT(name string, devicePath string) *NXT {
	return &NXT{
		Name:       name,
		DevicePath: devicePath,
	}
}

func (n *NXT) String() string {
	return fmt.Sprintf("NXT named %s, at %s", n.Name, n.DevicePath)
}

func (n *NXT) Connect() {
	fmt.Println("Called \"Connect\" but it is not yet implemented")
}

func (n *NXT) Disconnect() {
	fmt.Println("Called \"Disconnect\" but it is not yet implemented")
}
