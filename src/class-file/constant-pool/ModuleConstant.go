package constantpool

import (
	"fmt"

	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type ModuleConstant struct {
	Name ConstantReference[*Utf8Constant]
}

func (constant *ModuleConstant) ResolveReferences(cp *ConstantPool) {
	constant.Name.ResolveReferences(cp)
}

func (constant ModuleConstant) String() string {
	return fmt.Sprintf("ModuleConstant(Name: %s)", constant.Name.String())
}

func (constant ModuleConstant) Tag() ConstantPoolInfoTag {
	return CONSTANT_Module
}

func ParseConstantModule(binaryReader *reader.BinaryReader, index int) *ModuleConstant {
	return &ModuleConstant{
		Name: ParseConstantReference[*Utf8Constant](binaryReader, index),
	}
}
