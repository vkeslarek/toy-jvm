package attribute

import (
	"fmt"

	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type ParameterAnnotation struct {
	Annotations []Annotation
}

func ParseParameterAnnotation(fieldPrefix string, binaryReader *reader.BinaryReader, cp *constantpool.ConstantPool) ParameterAnnotation {
	numAnnotations := binaryReader.ReadUint16(fmt.Sprintf("%s.NumParameters", fieldPrefix))
	anotations := make([]Annotation, numAnnotations)

	for i := 0; i < int(numAnnotations); i++ {
		anotations[i] = ParseAnnotation(fmt.Sprintf("%s.Annotations[%d]", fieldPrefix, i), binaryReader, cp)
	}

	return ParameterAnnotation{
		Annotations: anotations,
	}
}
