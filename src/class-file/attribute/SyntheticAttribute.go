package attribute

import (
	"fmt"

	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type SyntheticAttribute struct {
}

func (*SyntheticAttribute) Name() AttributeName {
	return SyntheticAttributeName
}

func (s *SyntheticAttribute) String() string {
	return fmt.Sprintf("SyntheticAttribute()")
}

func ParseAttributeSynthetic(fieldPrefix string, attributeName *constantpool.Utf8Constant, attributeLength uint32, binaryReader *reader.BinaryReader, constantPool *constantpool.ConstantPool, index int) *SyntheticAttribute {
	return &SyntheticAttribute{}
}
