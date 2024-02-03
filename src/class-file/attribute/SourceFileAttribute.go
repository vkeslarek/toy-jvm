package attribute

import (
	"fmt"

	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type SourceFileAttribute struct {
	SourceFile *constantpool.Utf8Constant
}

func (a SourceFileAttribute) Name() AttributeName {
	return SourceFileAttributeName
}

func (a SourceFileAttribute) String() string {
	return fmt.Sprintf("SourceFileAttribute(SourceFile: %s)", a.SourceFile)
}

func ParseAttributeSourceFile(fieldPrefix string, attributeName *constantpool.Utf8Constant, attributeLength uint32, binaryReader *reader.BinaryReader, constantPool *constantpool.ConstantPool, index int) *SourceFileAttribute {
	return &SourceFileAttribute{
		SourceFile: constantPool.GetUtf8(binaryReader.ReadUint16("SourceFile")),
	}
}
