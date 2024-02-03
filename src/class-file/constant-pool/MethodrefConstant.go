package constantpool

import (
	"fmt"

	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type MethodrefConstant struct {
	Class       ConstantReference[*ClassConstant]
	NameAndType ConstantReference[*NameAndTypeConstant]
}

func (constant *MethodrefConstant) ResolveReferences(cp *ConstantPool) {
	constant.Class.ResolveReferences(cp)
	constant.NameAndType.ResolveReferences(cp)
}

func (constant MethodrefConstant) String() string {
	return fmt.Sprintf("MethodrefConstant(Class: %s, NameAndType: %s)", constant.Class.String(), constant.NameAndType.String())
}

func (constant MethodrefConstant) Tag() ConstantPoolInfoTag {
	return CONSTANT_Methodref
}

func ParseConstantMethodref(binaryReader *reader.BinaryReader, index int) *MethodrefConstant {
	return &MethodrefConstant{
		Class:       ParseConstantReference[*ClassConstant](binaryReader, index),
		NameAndType: ParseConstantReference[*NameAndTypeConstant](binaryReader, index),
	}
}
