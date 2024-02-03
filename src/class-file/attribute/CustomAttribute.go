package attribute

import (
	"fmt"

	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type CustomAttribute struct {
	NameConstant *constantpool.Utf8Constant
	Length       uint32
	Bytes        []uint8
}

func (attribute *CustomAttribute) Name() AttributeName {
	return AttributeName{
		Name: attribute.NameConstant.Value,
	}
}

func (attribute *CustomAttribute) String() string {
	return fmt.Sprintf("CustomAttribute(Name: %s, Length: %d, Bytes: %v)", attribute.Name(), attribute.Length, attribute.Bytes)
}

func ParseCustomAttribute(fieldPrefix string, attributeName *constantpool.Utf8Constant, attributeLength uint32, binaryReader *reader.BinaryReader, constantPool *constantpool.ConstantPool, index int) *CustomAttribute {
	return &CustomAttribute{
		NameConstant: attributeName,
		Length:       attributeLength,
		Bytes:        binaryReader.ReadBytes(fmt.Sprintf("%s.Attributes[%d].Bytes", fieldPrefix, index), int(attributeLength)),
	}
}
