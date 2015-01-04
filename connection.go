package nxt

import "io"

type Connection interface {
	io.ReadWriteCloser
	DevicePort() string
	Open() error
}

// TODO: Find a better place for these
func CalculateLSB(number int) byte {
	return byte(number & 0xff)
}

func CalculateMSB(number int) byte {
	return byte((number >> 8) & 0xff)
}

func CalculateIntFromLSBAndMSB(lsb byte, msb byte) int {
	return (int(msb) << 8) + int(lsb)
}
