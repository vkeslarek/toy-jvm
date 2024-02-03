package attribute

import (
	"fmt"

	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type ExceptionsAttribute struct {
	ExceptionIndexTable []*constantpool.ClassConstant
}

func (attribute *ExceptionsAttribute) Name() AttributeName {
	return ExceptionsAttributeName
}

func (attribute *ExceptionsAttribute) String() string {
	return fmt.Sprintf("ExceptionsAttribute{ExceptionIndexTable: %v}", attribute.ExceptionIndexTable)
}

func ParseAttributeExceptions(fieldPrefix string, attributeName *constantpool.Utf8Constant, attributeLength uint32, binaryReader *reader.BinaryReader, cp *constantpool.ConstantPool, index int) *ExceptionsAttribute {
	exceptionIndexTableLength := binaryReader.ReadUint16(fmt.Sprintf("%s.ExceptionIndexTableLength", fieldPrefix))
	exceptionIndexTable := make([]*constantpool.ClassConstant, exceptionIndexTableLength)
	for i := 0; i < int(exceptionIndexTableLength); i++ {
		exceptionIndexTable[i] = cp.GetClass(binaryReader.ReadUint16(fmt.Sprintf("%s.ExceptionIndexTable[%d]", fieldPrefix, i)))
	}
	return &ExceptionsAttribute{
		ExceptionIndexTable: exceptionIndexTable,
	}
}
