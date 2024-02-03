package attribute

import (
	"fmt"

	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type RuntimeVisibleTypeAnnotationsAttribute struct {
	Annotations []TypeAnnotation
}

func (a *RuntimeVisibleTypeAnnotationsAttribute) String() string {
	return fmt.Sprintf("RuntimeVisibleTypeAnnotationsAttribute(Annotations: %s)", a.Annotations)
}

func (a *RuntimeVisibleTypeAnnotationsAttribute) Name() AttributeName {
	return RuntimeVisibleTypeAnnotationsAttributeName
}

func ParseAttributeRuntimeVisibleTypeAnnotations(fieldPrefix string, attributeName *constantpool.Utf8Constant, attributeLength uint32, binaryReader *reader.BinaryReader, constantPool *constantpool.ConstantPool, index int) Attribute {
	return &RuntimeVisibleTypeAnnotationsAttribute{
		Annotations: ParseTypeAnnotations(binaryReader, constantPool),
	}
}
