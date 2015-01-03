package nxt

import (
	"io"
)

type lightReadWriteCloser struct {
	// TODO: Put in testing flags to simulate scenarios

	ReadCalled  bool
	WriteCalled bool
	CloseCalled bool
}

func (r *lightReadWriteCloser) Read(p []byte) (n int, err error) {
	r.ReadCalled = true

	return 0, nil
}

func (r *lightReadWriteCloser) Write(p []byte) (n int, err error) {
	r.WriteCalled = true

	return 0, nil
}

func (r *lightReadWriteCloser) Close() error {
	r.CloseCalled = true

	return nil
}

type mockConnection struct {
	devicePort string
}

func (c mockConnection) DevicePort() string {
	return c.devicePort
}

func (c *mockConnection) Open() (io.ReadWriteCloser, error) {
	return &lightReadWriteCloser{}, nil
}

func NewMockConnection(devicePort string) Connection {
	return &mockConnection{devicePort: devicePort}
}
