package nxt

import "fmt"

// StartProgram creates a Command to start a program with the filename passed in.
// NOTE: If the filename on the device does NOT have a file extension but just
// has a name, you should add ".rxe" to the end of the filename.
// To wait for the reply, pass in a replyChannel; to not wait, pass in nil
// for the replyChannel.
func StartProgram(filename string, replyChannel chan *ReplyTelegram) *Command {
	file := append([]byte(filename), 0) // null-terminated string

	return NewDirectCommand(0x00, file, replyChannel)
}

// StopProgram creates a Command to stop a running program.
// To wait for the reply, pass in a replyChannel; to not wait, pass in nil
// for the replyChannel.
func StopProgram(replyChannel chan *ReplyTelegram) *Command {
	return NewDirectCommand(0x01, nil, replyChannel)
}

// GetCurrentProgramName creates a Command to get the currently running program name.
// The filename will be passed on the message of the reply to this command.
func GetCurrentProgramName(replyChannel chan *ReplyTelegram) *Command {
	return NewDirectCommand(0x11, nil, replyChannel)
}

// GetCurrentProgramNameReply represent the reply to the GetCurrentProgramName call.
// The name of the running program can be accessed via the Filename member.
type GetCurrentProgramNameReply struct {
	*ReplyTelegram
	Filename string
}

// ParseGetCurrentProgramNameReply takes a raw ReplyTelegram and returns a
// GetCurrentProgramNameReply.
func ParseGetCurrentProgramNameReply(reply *ReplyTelegram) *GetCurrentProgramNameReply {
	return &GetCurrentProgramNameReply{
		ReplyTelegram: reply,
		Filename:      string(reply.Message),
	}
}

// StartProgram starts a program on the NXT with the given filename.
// This call is asynchronous and does not wait for a reply.  To wait
// for a reply to see if the call is successful, use StartProgramSync.
func (n NXT) StartProgram(filename string) {
	n.CommandChannel <- StartProgram(filename, nil)
}

// StartProgramSync starts a program on the NXT with the given filename.
// This call is snchronous waits for a reply.  If there was a problem
// starting the program, it will return a non-nil error.
func (n NXT) StartProgramSync(filename string) (*ReplyTelegram, error) {
	reply := make(chan *ReplyTelegram)

	n.CommandChannel <- StartProgram(filename, reply)
	startProgramReply := <-reply

	if !startProgramReply.IsSuccess() {
		return startProgramReply, fmt.Errorf("%v: \"%s\"", startProgramReply.Status, filename)
	}

	return startProgramReply, nil

}

// StopProgram stops the currently running program on the NXT.
// This call is asynchronous and does not wait for a reply.  To wait
// for a reply to see if the call is successful, use StopProgramSync.
func (n NXT) StopProgram() {
	n.CommandChannel <- StopProgram(nil)
}

// StopProgramSync stops the currently running program on the NXT.
// This call is snchronous waits for a reply.  If there was a problem
// stopping the program (usually because no program is running),
// it will return a non-nil error.
func (n NXT) StopProgramSync() (*ReplyTelegram, error) {
	reply := make(chan *ReplyTelegram)

	n.CommandChannel <- StopProgram(reply)
	stopProgramReply := <-reply

	if !stopProgramReply.IsSuccess() {
		return stopProgramReply, fmt.Errorf("%v", stopProgramReply.Status)
	}

	return stopProgramReply, nil

}

// GetCurrentProgramName gets the currently running program on the NXT.
// If there was a problem getting the program name
// (usually because no program is running), it will return a non-nil error.
func (n NXT) GetCurrentProgramName() (string, error) {
	reply := make(chan *ReplyTelegram)

	n.CommandChannel <- GetCurrentProgramName(reply)
	programNameReply := ParseGetCurrentProgramNameReply(<-reply)

	if !programNameReply.IsSuccess() {
		return "", fmt.Errorf("%v", programNameReply.Status)
	}

	return programNameReply.Filename, nil
}
