package constantpool

import (
	"fmt"

	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type ClassConstant struct {
	Name ConstantReference[*Utf8Constant]
}

func (constant *ClassConstant) ResolveReferences(cp *ConstantPool) {
	constant.Name.ResolveReferences(cp)
}

func (constant ClassConstant) String() string {
	return fmt.Sprintf("ClassConstant(Name: %s)", constant.Name.String())
}

func (constant ClassConstant) Tag() ConstantPoolInfoTag {
	return CONSTANT_Class
}

func ParseConstantClass(binaryReader *reader.BinaryReader, index int) *ClassConstant {
	return &ClassConstant{
		Name: ParseConstantReference[*Utf8Constant](binaryReader, index),
	}
}
