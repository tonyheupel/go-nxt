package nxt

import "io"

type Connection interface {
	io.ReadWriteCloser
	DevicePort() string
	Open() error
}
