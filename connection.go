package nxt

import "io"

type Connection interface {
	io.ReadWriteCloser
	DevicePort() string
	Open() error
}

// TODO: Find a better place for these
func calculateLSB(number int) byte {
	return byte(number & 0xff)
}

func calculateMSB(number int) byte {
	return byte((number >> 8) & 0xff)
}

func calculateIntFromLSBAndMSB(lsb byte, msb byte) int {
	return (int(msb) << 8) + int(lsb)
}
