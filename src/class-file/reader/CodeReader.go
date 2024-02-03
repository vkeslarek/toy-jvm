package reader

import (
	"bytes"
	"encoding/binary"
	"fmt"
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

func (b *CodeReader) ReadUint8(fieldName string) uint8 {
	return errorhandler.RunSilent[uint8](b, fmt.Sprintf("%s[%d]", fieldName, b.offset), func() (uint8, error) {
		var val uint8
		if err := binary.Read(b.reader, b.byteOrder, &val); err != nil {
			return 0, err
		}

		b.offset += 1
		return val, nil
	})
}
