package attribute

import (
	"fmt"

	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type Attributes []Attribute

func ParseAttributes(fieldPrefix string, binaryReader *reader.BinaryReader, constantPool *constantpool.ConstantPool) Attributes {
	attributesCount := binaryReader.ReadUint16(fmt.Sprintf("%s.AttributesCount", fieldPrefix))

	attributes := make([]Attribute, attributesCount)
	for i := 0; i < int(attributesCount); i++ {
		fieldPrefixWithAttributes := fmt.Sprintf("%s.Attributes[%d]", fieldPrefix, i)
		attributes[i] = ParseAttributeInfo(fieldPrefixWithAttributes, binaryReader, constantPool, i).Info
	}

	return attributes
}
