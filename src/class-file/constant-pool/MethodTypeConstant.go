package constantpool

import (
	"fmt"

	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type MethodTypeConstant struct {
	Descriptor ConstantReference[*Utf8Constant]
}

func (constant *MethodTypeConstant) ResolveReferences(cp *ConstantPool) {
	constant.Descriptor.ResolveReferences(cp)
}

func (constant MethodTypeConstant) String() string {
	return fmt.Sprintf("MethodTypeConstant(Descriptor: %s)", constant.Descriptor.String())
}

func (constant MethodTypeConstant) Tag() ConstantPoolInfoTag {
	return CONSTANT_MethodType
}

func ParseConstantMethodType(reader *reader.BinaryReader, index int) *MethodTypeConstant {
	return &MethodTypeConstant{
		Descriptor: ParseConstantReference[*Utf8Constant](reader, index),
	}
}
