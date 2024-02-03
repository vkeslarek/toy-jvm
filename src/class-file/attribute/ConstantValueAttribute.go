package attribute

import (
	"fmt"

	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type ConstantValueAttribute struct {
	ConstantValue constantpool.Constant
}

func (c *ConstantValueAttribute) Name() AttributeName {
	return ConstantValueAttributeName
}

func (c *ConstantValueAttribute) String() string {
	return fmt.Sprintf("ConstantValueAttribute(ConstantValue: %s)", c.ConstantValue)
}

func ParseAttributeConstantValue(fieldPrefix string, attributeName *constantpool.Utf8Constant, attributeLength uint32, binaryReader *reader.BinaryReader, constantPool *constantpool.ConstantPool, index int) *ConstantValueAttribute {
	return &ConstantValueAttribute{
		ConstantValue: constantPool.Get(int(binaryReader.ReadUint16("$.ConstantValue"))),
	}
}
