package constantpool

import (
	"fmt"

	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type NameAndTypeConstant struct {
	Name       ConstantReference[*Utf8Constant]
	Descriptor ConstantReference[*Utf8Constant]
}

func (constant *NameAndTypeConstant) ResolveReferences(cp *ConstantPool) {
	constant.Name.ResolveReferences(cp)
	constant.Descriptor.ResolveReferences(cp)
}

func (constant NameAndTypeConstant) String() string {
	return fmt.Sprintf("NameAndTypeConstant(Name: %s, Descriptor: %s)", constant.Name.String(), constant.Descriptor.String())
}

func (constant NameAndTypeConstant) Tag() ConstantPoolInfoTag {
	return CONSTANT_NameAndType
}

func ParseConstantNameAndType(binaryReader *reader.BinaryReader, index int) *NameAndTypeConstant {
	return &NameAndTypeConstant{
		Name:       ParseConstantReference[*Utf8Constant](binaryReader, index),
		Descriptor: ParseConstantReference[*Utf8Constant](binaryReader, index),
	}
}
