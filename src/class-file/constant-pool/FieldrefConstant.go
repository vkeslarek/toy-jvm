package constantpool

import (
	"fmt"

	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type FieldrefConstant struct {
	Class       ConstantReference[*ClassConstant]
	NameAndType ConstantReference[*NameAndTypeConstant]
}

func (constant *FieldrefConstant) ResolveReferences(cp *ConstantPool) {
	constant.Class.ResolveReferences(cp)
	constant.NameAndType.ResolveReferences(cp)
}

func (constant FieldrefConstant) String() string {
	return fmt.Sprintf("FieldrefConstant(Class: %s, NameAndType: %s)", constant.Class.String(), constant.NameAndType.String())
}

func (constant FieldrefConstant) Tag() ConstantPoolInfoTag {
	return CONSTANT_Fieldref
}

func ParseConstantFieldref(binaryReader *reader.BinaryReader, index int) *FieldrefConstant {
	return &FieldrefConstant{
		Class:       ParseConstantReference[*ClassConstant](binaryReader, index),
		NameAndType: ParseConstantReference[*NameAndTypeConstant](binaryReader, index),
	}
}
