package nxt

import "fmt"

func StartProgram(filename string, replyChannel chan *ReplyTelegram) *Command {
	file := append([]byte(filename), 0) // null-terminated string

	return NewDirectCommand(0x00, file, replyChannel)
}

func StopProgram(replyChannel chan *ReplyTelegram) *Command {
	return NewDirectCommand(0x01, nil, replyChannel)
}

func GetCurrentProgramName(replyChannel chan *ReplyTelegram) *Command {
	return NewDirectCommand(0x11, nil, replyChannel)
}

type GetCurrentProgramNameReply struct {
	*ReplyTelegram
	Filename string
}

func ParseGetCurrentProgramNameReply(reply *ReplyTelegram) *GetCurrentProgramNameReply {
	return &GetCurrentProgramNameReply{
		ReplyTelegram: reply,
		Filename:      string(reply.Message),
	}
}

func (n NXT) StartProgram(filename string) {
	n.CommandChannel <- StartProgram(filename, nil)
}

func (n NXT) StartProgramSync(filename string) (*ReplyTelegram, error) {
	reply := make(chan *ReplyTelegram)

	n.CommandChannel <- StartProgram(filename, reply)
	startProgramReply := <-reply

	if !startProgramReply.IsSuccess() {
		return startProgramReply, fmt.Errorf("Error trying to start program \"%s\": %v", filename, startProgramReply)
	}

	return startProgramReply, nil

}

func (n NXT) StopProgram() {
	n.CommandChannel <- StopProgram(nil)
}

func (n NXT) StopProgramSync() (*ReplyTelegram, error) {
	reply := make(chan *ReplyTelegram)

	n.CommandChannel <- StopProgram(reply)
	stopProgramReply := <-reply

	if !stopProgramReply.IsSuccess() {
		return stopProgramReply, fmt.Errorf("Error trying to stop program: %v", stopProgramReply)
	}

	return stopProgramReply, nil

}
func (n NXT) GetCurrentProgramName() (string, error) {
	reply := make(chan *ReplyTelegram)

	n.CommandChannel <- GetCurrentProgramName(reply)
	programNameReply := ParseGetCurrentProgramNameReply(<-reply)

	if !programNameReply.IsSuccess() {
		return "", fmt.Errorf("Error getting current program name:", programNameReply)
	}

	return programNameReply.Filename, nil
}
