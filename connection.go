package nxt

import "io"

type Connection interface {
	DevicePort() string
	Open() (io.ReadWriteCloser, error)
}
