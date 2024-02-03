package attribute

import (
	"fmt"

	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type EnclosingMethodAttribute struct {
	Class  *constantpool.ClassConstant
	Method *constantpool.NameAndTypeConstant
}

func (a *EnclosingMethodAttribute) Name() AttributeName {
	return EnclosingMethodAttributeName
}

func (a *EnclosingMethodAttribute) String() string {
	return fmt.Sprintf("EnclosingMethodAttribute(Class: %s, Method: %s)", a.Class, a.Method)
}

func ParseAttributeEnclosingMethod(fieldPrefix string, attributeName *constantpool.Utf8Constant, attributeLength uint32, binaryReader *reader.BinaryReader, constantPool *constantpool.ConstantPool, index int) *EnclosingMethodAttribute {
	return &EnclosingMethodAttribute{
		Class:  constantPool.GetClass(binaryReader.ReadUint16(fmt.Sprintf("%s.Class", fieldPrefix))),
		Method: constantPool.GetNameAndType(binaryReader.ReadUint16(fmt.Sprintf("%s.Method", fieldPrefix))),
	}
}
