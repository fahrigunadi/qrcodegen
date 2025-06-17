package utils

import "io"

type NopWriteCloser struct {
	io.Writer
}

func (n NopWriteCloser) Close() error {
	return nil
}
