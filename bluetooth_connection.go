package nxt

import (
	"fmt"
	"github.com/tarm/goserial"
	"io"
)

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

func (b bluetoothConnection) DevicePort() string {
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
	//TODO: Strip out the first two bytes of bluetooth length headers

	bluetoothMessage := make([]byte, 66)

	bytesRead, err := b.conduit.Read(bluetoothMessage)

	if err != nil {
		return bytesRead, err
	}

	var length int
	if bytesRead >= 2 {
		// Remove bluetooth headers and only get length back
		length = calculateIntFromLSBAndMSB(bluetoothMessage[0], bluetoothMessage[1])
		copy(p, bluetoothMessage[2:])
	}

	fmt.Println("Length:", length, "Response:", p)
	return length, nil
}

func (b *bluetoothConnection) Write(p []byte) (n int, err error) {
	//TODO: Add the first two bytes of bluetooth length headers
	telegramLength := len(p)
	bluetoothHeader := []byte{calculateLSB(telegramLength), calculateMSB(telegramLength)}

	bluetoothMessage := append(bluetoothHeader, p...)

	fmt.Println("Bluetooth Message:", bluetoothMessage)

	return b.conduit.Write(bluetoothMessage)
}
func (b bluetoothConnection) Close() error {
	return b.conduit.Close()
}
