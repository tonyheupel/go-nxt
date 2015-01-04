package nxt

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
	*lightReadWriteCloser
	devicePort string
}

func (c mockConnection) Port() string {
	return c.devicePort
}

func (c *mockConnection) Open() error {
	return nil
}

func NewMockConnection(devicePort string) Connection {
	return &mockConnection{
		lightReadWriteCloser: &lightReadWriteCloser{},
		devicePort:           devicePort,
	}
}
