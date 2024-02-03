package attribute

import (
	"fmt"

	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type RuntimeInvisibleTypeAnnotationsAttribute struct {
	Annotations []TypeAnnotation
}

func (a *RuntimeInvisibleTypeAnnotationsAttribute) String() string {
	return fmt.Sprintf("RuntimeInvisibleTypeAnnotationsAttribute(Annotations: %s)", a.Annotations)
}

func (a *RuntimeInvisibleTypeAnnotationsAttribute) Name() AttributeName {
	return RuntimeInvisibleTypeAnnotationsAttributeName
}

func ParseAttributeRuntimeInvisibleTypeAnnotations(fieldPrefix string, attributeName *constantpool.Utf8Constant, attributeLength uint32, binaryReader *reader.BinaryReader, constantPool *constantpool.ConstantPool, index int) Attribute {
	return &RuntimeInvisibleTypeAnnotationsAttribute{
		Annotations: ParseTypeAnnotations(binaryReader, constantPool),
	}
}
