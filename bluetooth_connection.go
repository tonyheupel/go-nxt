package nxt

import (
	"github.com/tarm/goserial"
	"io"
)

// bluetoothConnection represents the default connection to an NXT device.
type bluetoothConnection struct {
	config  *serial.Config
	conduit io.ReadWriteCloser
}

func newBluetoothConnection(name string) Connection {
	config := &serial.Config{Name: name, Baud: 115200}

	return &bluetoothConnection{
		config:  config,
		conduit: nil,
	}
}

func (b bluetoothConnection) Port() string {
	return b.config.Name
}

func (b *bluetoothConnection) Open() error {

	conduit, err := serial.OpenPort(b.config)

	if err != nil {
		return err
	}

	b.conduit = conduit

	return nil
}

func (b bluetoothConnection) Read(p []byte) (n int, err error) {
	bluetoothMessage := make([]byte, 66)

	bytesRead, err := b.conduit.Read(bluetoothMessage)

	if err != nil {
		return bytesRead, err
	}

	var length int
	if bytesRead >= 2 {
		// Remove bluetooth two-byte length headers and only get length back
		length = calculateIntFromLSBAndMSB(bluetoothMessage[0], bluetoothMessage[1])
		copy(p, bluetoothMessage[2:])
	}

	return length, nil
}

func (b *bluetoothConnection) Write(p []byte) (n int, err error) {
	telegramLength := len(p)
	// Bluetooth messages also require a two-byte header representing the length
	// of the message being sent (NOT including the two-byte header)
	bluetoothHeader := []byte{calculateLSB(telegramLength), calculateMSB(telegramLength)}

	bluetoothMessage := append(bluetoothHeader, p...)

	return b.conduit.Write(bluetoothMessage)
}
func (b bluetoothConnection) Close() error {
	return b.conduit.Close()
}
