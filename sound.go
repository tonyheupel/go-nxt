package nxt

import "fmt"

// PlaySoundFile creates a Command to play a sound file given a filename and
// whether it should loop or not.
// NOTE: The NXT will so the filename without a file extension.  If playing
// that file does not work, try adding ".rso" as the extension for sound files.
// To wait for the reply, pass in a replyChannel; to not wait, pass in nil
// for the replyChannel.
func PlaySoundFile(filename string, loop bool, replyChannel chan *ReplyTelegram) *Command {
	var loopBytes []byte
	if loop {
		loopBytes = []byte{0xff}
	} else {
		loopBytes = []byte{0x00}
	}
	fileBytes := append([]byte(filename), 0) // null-terminated string
	message := append(loopBytes, fileBytes...)

	return NewDirectCommand(0x02, message, replyChannel)
}

// PlayTone plays a tone with the Hz specified in frequency and for the
// duration of duration milliseconds.
// To wait for the reply, pass in a replyChannel; to not wait, pass in nil
// for the replyChannel.
func PlayTone(frequency int, duration int, replyChannel chan *ReplyTelegram) *Command {
	frequencyBytes := []byte{calculateLSB(frequency), calculateMSB(frequency)}
	durationBytes := []byte{calculateLSB(duration), calculateMSB(duration)}
	message := append(frequencyBytes, durationBytes...)

	return NewDirectCommand(0x03, message, replyChannel)
}

// StopSoundPlayback tells the NXT to stop playing whatever sound
// is currently playing.
func StopSoundPlayback(replyChannel chan *ReplyTelegram) *Command {
	return NewDirectCommand(0x0C, nil, replyChannel)
}

// PlaySoundFile plays the sound file with the given filename and
// loops the file if true is passed in for loop; does not loop if
// false is passed in.
// NOTE: The NXT will so the filename without a file extension.  If playing
// that file does not work, try adding ".rso" as the extension for sound files.
// This call is asynchronous and does not wait for a reply.  To wait
// for a reply to see if the call is successful, use PlayToneSync.
func (n NXT) PlaySoundFile(filename string, loop bool) {
	n.CommandChannel <- PlaySoundFile(filename, loop, nil)
}

// PlaySoundFileSync plays the sound file with the given filename and
// loops the file if true is passed in for loop; does not loop if
// false is passed in.
// NOTE: The NXT will so the filename without a file extension.  If playing
// that file does not work, try adding ".rso" as the extension for sound files.
// This call is snchronous waits for a reply.  If there was a problem
// starting the program, it will return a non-nil error.
func (n NXT) PlaySoundFileSync(filename string, loop bool) (*ReplyTelegram, error) {
	reply := make(chan *ReplyTelegram)

	n.CommandChannel <- PlaySoundFile(filename, loop, reply)
	playSoundFileReply := <-reply

	if !playSoundFileReply.IsSuccess() {
		return playSoundFileReply, fmt.Errorf("%v: \"%s\"", playSoundFileReply.Status, filename)
	}

	return playSoundFileReply, nil

}

// PlayTone plays a tone with the Hz specified in frequency and for the
// duration of duration milliseconds.
// This call is asynchronous and does not wait for a reply.  To wait
// for a reply to see if the call is successful, use PlayToneSync
func (n NXT) PlayTone(frequency int, duration int) {
	n.CommandChannel <- PlayTone(frequency, duration, nil)
}

// PlayToneSync plays a tone with the Hz specified in frequency and for the
// duration of duration milliseconds.
// This call is snchronous waits for a reply.  If there was a problem,
// it will return a non-nil error.
func (n NXT) PlayToneSync(frequency int, duration int) (*ReplyTelegram, error) {
	reply := make(chan *ReplyTelegram)

	n.CommandChannel <- PlayTone(frequency, duration, reply)
	playToneReply := <-reply

	if !playToneReply.IsSuccess() {
		return playToneReply, fmt.Errorf("%v", playToneReply.Status)
	}

	return playToneReply, nil
}

// StopSoundPlayback tells the NXT to stop playing whatever sound
// is currently playing.
// This call is asynchronous and does not wait for a reply.  To wait
// for a reply to see if the call is successful, use PlayToneSync
func (n NXT) StopSoundPlayback() {
	n.CommandChannel <- StopSoundPlayback(nil)
}

// StopSoundPlaybackSync tells the NXT to stop playing whatever sound
// is currently playing.
// This call is snchronous waits for a reply.  If there was a problem,
// it will return a non-nil error.
func (n NXT) StopSoundPlaybackSync() (*ReplyTelegram, error) {
	reply := make(chan *ReplyTelegram)

	n.CommandChannel <- StopSoundPlayback(reply)
	stopSoundPlaybackReply := <-reply

	if !stopSoundPlaybackReply.IsSuccess() {
		return stopSoundPlaybackReply, fmt.Errorf("%v", stopSoundPlaybackReply.Status)
	}

	return stopSoundPlaybackReply, nil
}
