package constantpool

import (
	"fmt"

	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type ReferenceKind uint8

const (
	REF_getField         ReferenceKind = 1
	REF_getStatic        ReferenceKind = 2
	REF_putField         ReferenceKind = 3
	REF_putStatic        ReferenceKind = 4
	REF_invokeVirtual    ReferenceKind = 5
	REF_invokeStatic     ReferenceKind = 6
	REF_invokeSpecial    ReferenceKind = 7
	REF_newInvokeSpecial ReferenceKind = 8
	REF_invokeInterface  ReferenceKind = 9
)

type MethodHandleConstant struct {
	ReferenceKind ReferenceKind
	Reference     ConstantReference[Constant]
}

func (m *MethodHandleConstant) ResolveReferences(cp *ConstantPool) {
	m.Reference.ResolveReferences(cp)
}

func (m MethodHandleConstant) String() string {
	return fmt.Sprintf("MethodHandleConstant(ReferenceKind: %d, Reference: %s)", m.ReferenceKind, m.Reference.String())
}

func (m MethodHandleConstant) Tag() ConstantPoolInfoTag {
	return CONSTANT_MethodHandle
}

func ParseConstantMethodHandle(binaryReader *reader.BinaryReader, index int) *MethodHandleConstant {
	return &MethodHandleConstant{
		ReferenceKind: ReferenceKind(binaryReader.ReadUint8(fmt.Sprintf("$.ConstantPool[%d](MethodHandleConstant).ReferenceKind", index))),
		Reference:     ParseConstantReference[Constant](binaryReader, index),
	}
}
