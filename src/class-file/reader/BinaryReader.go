package reader

import (
	"encoding/binary"
	"io"

	errorhandler "github.com/vkeslarek/toy-jvm/error-handler"
)

type BinaryReader struct {
	reader    io.Reader
	byteOrder binary.ByteOrder
	err       errorhandler.ParserError
}

func (b *BinaryReader) SetError(err errorhandler.ParserError) {
	b.err = err
}

func (b *BinaryReader) GetError() errorhandler.ParserError {
	return b.err
}

func NewBinaryReader(reader io.Reader, byteOrder binary.ByteOrder) *BinaryReader {
	return &BinaryReader{
		reader:    reader,
		byteOrder: byteOrder,
	}
}

func (b *BinaryReader) ReadUint8(fieldName string) uint8 {
	return errorhandler.RunSilent[uint8](b, fieldName, func() (uint8, error) {
		var val uint8
		if err := binary.Read(b.reader, b.byteOrder, &val); err != nil {
			return 0, err
		}

		return val, nil
	})
}

func (b *BinaryReader) ReadUint16(fieldName string) uint16 {
	return errorhandler.RunSilent[uint16](b, fieldName, func() (uint16, error) {
		var val uint16
		if err := binary.Read(b.reader, b.byteOrder, &val); err != nil {
			return 0, err
		}

		return val, nil
	})
}

func (b *BinaryReader) ReadUint32(fieldname string) uint32 {
	return errorhandler.RunSilent[uint32](b, fieldname, func() (uint32, error) {
		var val uint32
		if err := binary.Read(b.reader, b.byteOrder, &val); err != nil {
			return 0, err
		}

		return val, nil
	})
}

func (b *BinaryReader) ReadUint64(fieldname string) uint64 {
	return errorhandler.RunSilent[uint64](b, fieldname, func() (uint64, error) {
		var val uint64
		if err := binary.Read(b.reader, b.byteOrder, &val); err != nil {
			return 0, err
		}

		return val, nil
	})
}

func (b *BinaryReader) ReadStruct(fieldName string, structType interface{}) interface{} {
	return errorhandler.RunSilent[interface{}](b, fieldName, func() (interface{}, error) {
		return structType, binary.Read(b.reader, b.byteOrder, structType)
	})
}

func (b *BinaryReader) ReadString(fieldName string, length int) string {
	return errorhandler.RunSilent[string](b, fieldName, func() (string, error) {
		byteBuffer := make([]byte, length)

		if err := binary.Read(b.reader, b.byteOrder, &byteBuffer); err != nil {
			return "", err
		}

		return string(byteBuffer), nil
	})
}

func (b *BinaryReader) ReadBytes(fieldName string, length int) []byte {
	return errorhandler.RunSilent[[]byte](b, fieldName, func() ([]byte, error) {
		byteBuffer := make([]byte, length)

		if err := binary.Read(b.reader, b.byteOrder, &byteBuffer); err != nil {
			return nil, err
		}

		return byteBuffer, nil
	})
}
