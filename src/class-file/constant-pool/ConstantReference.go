package constantpool

import (
	"fmt"

	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type ConstantReference[T Constant] struct {
	Index     uint16
	Resolved  bool
	Exists    bool
	Reference T
}

func (cr *ConstantReference[T]) ResolveReferences(cp *ConstantPool) {
	if !cr.Resolved {
		constant := cp.Get(int(cr.Index))
		if constant != nil {
			cr.Reference = constant.(T)
			cr.Exists = true
		}

		cr.Resolved = true
	}
}

func (cr *ConstantReference[T]) String() string {
	if !cr.Resolved {
		return fmt.Sprintf("UnresolvedConstantReference(Type: %T, Index: %d)", *new(T), cr.Index)
	}

	if !cr.Exists {
		return fmt.Sprintf("MissingConstantReference(Type: %T, Index: %d)", *new(T), cr.Index)
	}

	return cr.Reference.String()
}

func ParseConstantReference[T Constant](binaryReader *reader.BinaryReader, index int) ConstantReference[T] {
	return ConstantReference[T]{
		Index:    binaryReader.ReadUint16(fmt.Sprintf("$.ConstantPool[%d](ConstantReferenceOf(%T)).Index", index, *new(T))),
		Resolved: false,
		Exists:   false,
	}
}
