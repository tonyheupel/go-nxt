package nxt

func startProgram(n NXT, requireResponse bool, filename string) error {
	file := append([]byte(filename), 0) // null-terminated string

	telegram := NewDirectCommand(requireResponse, 0x00, file)

	_, err := n.connection.Write(telegram.Bytes())

	if err != nil {
		return err
	}

	return nil
}

func (n NXT) StartProgram(filename string) error {
	return startProgram(n, false, filename)
}

func (n NXT) StartProgramSync(filename string) (*ReplyTelegram, error) {

	err := startProgram(n, true, filename)

	if err != nil {
		return nil, err
	}

	return getReplyFromReader(n.connection), nil
}

func stopProgram(n NXT, requireResponse bool) error {
	telegram := NewDirectCommand(requireResponse, 0x01, nil)

	_, err := n.connection.Write(telegram.Bytes())

	if err != nil {
		return err
	}

	return nil
}

func (n NXT) StopProgram() error {
	return stopProgram(n, false)
}

func (n NXT) StopProgramSync() (*ReplyTelegram, error) {

	err := stopProgram(n, true)

	if err != nil {
		return nil, err
	}

	return getReplyFromReader(n.connection), nil
}

func (n NXT) GetCurrentProgramName() (string, error) {
	telegram := NewDirectCommand(true, 0x11, nil)

	command := telegram.Bytes()

	_, err := n.connection.Write(command)

	if err != nil {
		return "", err
	}

	reply := getReplyFromReader(n.connection)

	return string(reply.Message), nil
}
