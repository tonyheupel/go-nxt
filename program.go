package nxt

func startProgram(n NXT, requireResponse bool, filename string) error {
	file := append([]byte(filename), 0) // null-terminated string

	telegram := NewDirectCommand(requireResponse, 0x00, file)

	n.CommandChannel <- *telegram

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

	reply := <-n.ReplyChannel
	return &reply, nil
}

func stopProgram(n NXT, requireResponse bool) error {
	telegram := NewDirectCommand(requireResponse, 0x01, nil)

	n.CommandChannel <- *telegram

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

	reply := <-n.ReplyChannel
	return &reply, nil
}

func (n NXT) GetCurrentProgramName() (string, error) {
	telegram := NewDirectCommand(true, 0x11, nil)

	n.CommandChannel <- *telegram

	reply := <-n.ReplyChannel

	return string(reply.Message), nil
}
