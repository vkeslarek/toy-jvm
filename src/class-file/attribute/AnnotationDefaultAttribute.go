package attribute

import (
	"fmt"

	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type AnnotationDefaultAttribute struct {
	DefaultValue ElementValue
}

func (a *AnnotationDefaultAttribute) String() string {
	return fmt.Sprintf("AnnotationDefaultAttribute(DefaultValue: %s)", a.DefaultValue)
}

func (a *AnnotationDefaultAttribute) Name() AttributeName {
	return AnnotationDefaultAttributeName
}

func ParseAttributeAnnotationDefault(fieldPrefix string, attributeName *constantpool.Utf8Constant, attributeLength uint32, binaryReader *reader.BinaryReader, constantPool *constantpool.ConstantPool, index int) *AnnotationDefaultAttribute {
	return &AnnotationDefaultAttribute{
		DefaultValue: ParseElementValue(fmt.Sprintf("%s.DefaultValue", fieldPrefix), binaryReader, constantPool),
	}
}
