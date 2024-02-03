package constantpool

import (
	"fmt"

	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type DynamicConstant struct {
	BootstrapMethodAttrIndex uint16
	NameAndType              ConstantReference[*NameAndTypeConstant]
}

func (constant *DynamicConstant) ResolveReferences(cp *ConstantPool) {
	constant.NameAndType.ResolveReferences(cp)
}

func (constant DynamicConstant) String() string {
	return fmt.Sprintf("DynamicConstant(BootstrapMethodAttrIndex: %d, NameAndType: %s)", constant.BootstrapMethodAttrIndex, constant.NameAndType.String())
}

func (constant DynamicConstant) Tag() ConstantPoolInfoTag {
	return CONSTANT_Dynamic
}

func ParseConstantDynamic(reader *reader.BinaryReader, index int) *DynamicConstant {
	return &DynamicConstant{
		BootstrapMethodAttrIndex: reader.ReadUint16(fmt.Sprintf("$.ConstantPool[%d](DynamicConstant).BootstrapMethodAttrIndex", index)),
		NameAndType:              ParseConstantReference[*NameAndTypeConstant](reader, index),
	}
}
