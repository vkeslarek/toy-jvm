package attribute

import (
	"fmt"

	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type SourceDebugExtensionAttribute struct {
	SourceDebugExtension []byte
}

func (a *SourceDebugExtensionAttribute) Name() AttributeName {
	return SourceDebugExtensionAttributeName
}

func (a *SourceDebugExtensionAttribute) String() string {
	return fmt.Sprintf("SourceDebugExtensionAttribute(SourceDebugExtension: %v)", a.SourceDebugExtension)
}

func ParseAttributeSourceDebugExtension(fieldPrefix string, attributeName *constantpool.Utf8Constant, attributeLength uint32, binaryReader *reader.BinaryReader, constantPool *constantpool.ConstantPool, index int) *SourceDebugExtensionAttribute {
	return &SourceDebugExtensionAttribute{
		SourceDebugExtension: binaryReader.ReadBytes(fmt.Sprintf("%s.SourceDebugExtension", fieldPrefix), int(attributeLength)),
	}
}
