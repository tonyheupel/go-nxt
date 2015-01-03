package nxt

import (
	"github.com/tarm/goserial"
	"io"
)

type bluetoothConnection struct {
	config  *serial.Config
	conduit *io.ReadWriteCloser
}

func (b bluetoothConnection) DevicePort() string {
	return b.config.Name
}

func (b *bluetoothConnection) Open() (io.ReadWriteCloser, error) {

	conduit, err := serial.OpenPort(b.config)

	if err != nil {
		return nil, err
	}

	b.conduit = &conduit

	return conduit, nil
}

func newBluetoothConnection(name string) Connection {
	config := &serial.Config{Name: name, Baud: 57600}

	return &bluetoothConnection{
		config:  config,
		conduit: nil,
	}
}
