package field

import (
	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type Fields []*Field

func ParseFields(binaryReader *reader.BinaryReader, constantPool *constantpool.ConstantPool) Fields {
	fieldCount := binaryReader.ReadUint16("$.FieldsCount")
	fields := make([]*Field, fieldCount)
	for i := 0; i < int(fieldCount); i++ {
		fields[i] = ParseField(binaryReader, constantPool, i)
	}
	return fields
}
