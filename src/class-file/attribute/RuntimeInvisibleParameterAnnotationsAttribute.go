package attribute

import (
	"fmt"

	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type RuntimeInvisibleParameterAnnotationsAttribute struct {
	ParameterAnnotations []ParameterAnnotation
}

func (a *RuntimeInvisibleParameterAnnotationsAttribute) String() string {
	return fmt.Sprintf("RuntimeInvisibleParameterAnnotationsAttribute(Annotations: %v)", a.ParameterAnnotations)
}

func (a *RuntimeInvisibleParameterAnnotationsAttribute) Name() AttributeName {
	return RuntimeVisibleParameterAnnotationsAttributeName
}

func ParseAttributeRuntimeInvisibleParameterAnnotations(fieldPrefix string, attributeName *constantpool.Utf8Constant, attributeLength uint32, binaryReader *reader.BinaryReader, constantPool *constantpool.ConstantPool, index int) Attribute {
	numParameters := binaryReader.ReadUint16(fmt.Sprintf("%s.NumParameters", fieldPrefix))
	parameterAnnotations := make([]ParameterAnnotation, numParameters)

	for i := 0; i < int(numParameters); i++ {
		parameterAnnotations[i] = ParseParameterAnnotation(fmt.Sprintf("%s.ParameterAnnotations[%d]", fieldPrefix, i), binaryReader, constantPool)
	}

	return &RuntimeInvisibleParameterAnnotationsAttribute{
		ParameterAnnotations: parameterAnnotations,
	}
}
