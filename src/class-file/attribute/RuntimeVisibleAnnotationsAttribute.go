package attribute

import (
	"fmt"

	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type RuntimeVisibleAnnotationsAttribute struct {
	Annotations []Annotation
}

func (a *RuntimeVisibleAnnotationsAttribute) String() string {
	return fmt.Sprintf("RuntimeVisibleAnnotationsAttribute(Annotations: %s)", a.Annotations)
}

func (a *RuntimeVisibleAnnotationsAttribute) Name() AttributeName {
	return RuntimeVisibleAnnotationsAttributeName
}

func ParseAttributeRuntimeVisibleAnnotations(fieldPrefix string, attributeName *constantpool.Utf8Constant, attributeLength uint32, binaryReader *reader.BinaryReader, constantPool *constantpool.ConstantPool, index int) Attribute {
	annotations := make([]Annotation, attributeLength/2)
	for i := 0; i < int(attributeLength/2); i++ {
		annotations[i] = ParseAnnotation(fmt.Sprintf("%s.Annotations[%d]", fieldPrefix, i), binaryReader, constantPool)
	}
	return &RuntimeVisibleAnnotationsAttribute{
		Annotations: annotations,
	}
}
