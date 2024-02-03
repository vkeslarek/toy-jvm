package constantpool

import (
	"fmt"

	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type InvokeDynamicConstant struct {
	BootstrapMethodAttrIndex uint16
	NameAndType              ConstantReference[*NameAndTypeConstant]
}

func (constant *InvokeDynamicConstant) ResolveReferences(cp *ConstantPool) {
	constant.NameAndType.ResolveReferences(cp)
}

func (constant InvokeDynamicConstant) String() string {
	return fmt.Sprintf("InvokeDynamicConstant(BootstrapMethodAttrIndex: %d, NameAndType: %s)", constant.BootstrapMethodAttrIndex, constant.NameAndType.String())
}

func (constant InvokeDynamicConstant) Tag() ConstantPoolInfoTag {
	return CONSTANT_InvokeDynamic
}

func ParseConstantInvokeDynamic(reader *reader.BinaryReader, index int) *InvokeDynamicConstant {
	return &InvokeDynamicConstant{
		BootstrapMethodAttrIndex: reader.ReadUint16(fmt.Sprintf("$.ConstantPool[%d](InvokeDynamicConstant).BootstrapMethodAttrIndex", index)),
		NameAndType:              ParseConstantReference[*NameAndTypeConstant](reader, index),
	}
}
