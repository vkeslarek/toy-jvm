package reader

import (
	"bytes"
	"encoding/binary"
	"io"

	errorhandler "github.com/vkeslarek/toy-jvm/error-handler"
)

type CodeReader struct {
	reader    io.Reader
	byteOrder binary.ByteOrder
	offset    uint32
	err       errorhandler.ParserError
}

func (b *CodeReader) SetError(err errorhandler.ParserError) {
	b.err = err
}

func (b *CodeReader) GetError() errorhandler.ParserError {
	return b.err
}

func NewCodeReader(code []byte, byteOrder binary.ByteOrder) *CodeReader {
	return &CodeReader{
		reader:    bytes.NewReader(code),
		offset:    0,
		byteOrder: byteOrder,
	}
}
