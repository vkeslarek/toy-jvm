package constantpool

import (
	"fmt"

	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type IntegerConstant struct {
	Value uint32
}

func (i *IntegerConstant) ResolveReferences(cp *ConstantPool) {
	// Nothing to do here
}

func (i *IntegerConstant) String() string {
	return fmt.Sprintf("IntegerConstant(Value: %d)", i.Value)
}

func (i *IntegerConstant) Tag() ConstantPoolInfoTag {
	return CONSTANT_Integer
}

func ParseConstantInteger(binaryReader *reader.BinaryReader, index int) *IntegerConstant {
	return &IntegerConstant{
		Value: binaryReader.ReadUint32(fmt.Sprintf("$.ConstantPool[%d](IntegerConstant).Bytes", index)),
	}
}
