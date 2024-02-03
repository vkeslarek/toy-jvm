package constantpool

import (
	"fmt"

	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type InterfaceMethodrefConstant struct {
	Class       ConstantReference[*ClassConstant]
	NameAndType ConstantReference[*NameAndTypeConstant]
}

func (constant *InterfaceMethodrefConstant) ResolveReferences(cp *ConstantPool) {
	constant.Class.ResolveReferences(cp)
	constant.NameAndType.ResolveReferences(cp)
}

func (constant InterfaceMethodrefConstant) String() string {
	return fmt.Sprintf("InterfaceMethodrefConstant(Class: %s, NameAndType: %s)", constant.Class.String(), constant.NameAndType.String())
}

func (constant InterfaceMethodrefConstant) Tag() ConstantPoolInfoTag {
	return CONSTANT_InterfaceMethodref
}

func ParseConstantInterfaceMethodref(binaryReader *reader.BinaryReader, index int) *InterfaceMethodrefConstant {
	return &InterfaceMethodrefConstant{
		Class:       ParseConstantReference[*ClassConstant](binaryReader, index),
		NameAndType: ParseConstantReference[*NameAndTypeConstant](binaryReader, index),
	}
}
