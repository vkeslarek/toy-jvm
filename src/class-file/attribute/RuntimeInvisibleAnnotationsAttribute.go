package attribute

import (
	"fmt"

	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type RuntimeInvisibleAnnotationsAttribute struct {
	Annotations []Annotation
}

func (a *RuntimeInvisibleAnnotationsAttribute) String() string {
	return fmt.Sprintf("RuntimeInvisibleAnnotationsAttribute(Annotations: %v)", a.Annotations)
}

func (a *RuntimeInvisibleAnnotationsAttribute) Name() AttributeName {
	return RuntimeInvisibleAnnotationsAttributeName
}

func ParseAttributeRuntimeInvisibleAnnotations(fieldPrefix string, attributeName *constantpool.Utf8Constant, attributeLength uint32, binaryReader *reader.BinaryReader, constantPool *constantpool.ConstantPool, index int) Attribute {
	annotations := make([]Annotation, attributeLength/2)
	for i := 0; i < int(attributeLength/2); i++ {
		annotations[i] = ParseAnnotation(fmt.Sprintf("Annotations[%d]", i), binaryReader, constantPool)
	}
	return &RuntimeInvisibleAnnotationsAttribute{
		Annotations: annotations,
	}
}
