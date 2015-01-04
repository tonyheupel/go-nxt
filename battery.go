package nxt

func (n NXT) GetBatteryLevelMillivolts() (int, error) {
	telegram := NewDirectCommand(true, 0x0B, nil)

	_, err := n.connection.Write(telegram.Bytes())

	if err != nil {
		return 0, nil
	}

	reply := getReplyFromReader(n.connection)

	return calculateIntFromLSBAndMSB(reply.Message[0], reply.Message[1]), nil

}
