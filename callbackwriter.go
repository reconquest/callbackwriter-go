package main

import (
	"io"
)

type Writer struct {
	backend io.WriteCloser

	write func([]byte)
	close func()
}

func NewWriter(
	backend io.WriteCloser,
	write func([]byte),
	close func(),
) (*Writer, error) {
	return &Writer{
		backend: backend,
		write:   write,
		close:   close,
	}
}

func (writer *Writer) Write(data []byte) (int, error) {
	if writer.write {
		writer.write(data)
	}

	return writer.backend.Write(data)
}

func (writer *Writer) Closer() error {
	if writer.close {
		writer.close()
	}

	return writer.backend.Close()
}
