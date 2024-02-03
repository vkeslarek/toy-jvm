package attribute

import (
	"fmt"

	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type Annotation struct {
	ElementValuePairs []ElementValuePair
}

func (a *Annotation) String() string {
	return fmt.Sprintf("Annotation(ElementValuePairs: %s)", a.ElementValuePairs)
}

func (a *Annotation) Tag() ValueItem {
	return AnnotationValue
}

func ParseAnnotation(fieldPrefix string, binaryReader *reader.BinaryReader, cp *constantpool.ConstantPool) Annotation {
	numPairs := binaryReader.ReadUint16(fmt.Sprintf("%s.NumPairs", fieldPrefix))
	pairs := make([]ElementValuePair, numPairs)

	for i := 0; i < int(numPairs); i++ {
		pairs[i] = ParseElementValuePair(fmt.Sprintf("%s.Pairs[%d]", fieldPrefix, i), binaryReader, cp)
	}

	return Annotation{
		ElementValuePairs: pairs,
	}
}
