package attribute

import (
	"fmt"

	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type DeprecatedAttribute struct {
}

func (d *DeprecatedAttribute) String() string {
	return fmt.Sprintf("DeprecatedAttribute()")
}

func (*DeprecatedAttribute) Name() AttributeName {
	return DeprecatedAttributeName
}

func ParseAttributeDeprecated(fieldPrefix string, attributeName *constantpool.Utf8Constant, attributeLength uint32, binaryReader *reader.BinaryReader, constantPool *constantpool.ConstantPool, index int) Attribute {
	return &DeprecatedAttribute{}
}
