package constantpool

import (
	"fmt"

	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type LongConstant struct {
	Value uint64
}

func (l *LongConstant) ResolveReferences(cp *ConstantPool) {
	// Nothing to do here
}

func (l *LongConstant) String() string {
	return fmt.Sprintf("LongConstant(Value: %d)", l.Value)
}

func (*LongConstant) Tag() ConstantPoolInfoTag {
	return CONSTANT_Long
}

func ParseConstantLong(binaryReader *reader.BinaryReader, index int) *LongConstant {
	highBytes := binaryReader.ReadUint32(fmt.Sprintf("$.ConstantPool[%d](LongConstant).HighBytes", index))
	lowBytes := binaryReader.ReadUint32(fmt.Sprintf("$.ConstantPool[%d](LongConstant).LowBytes", index))

	return &LongConstant{
		Value: uint64(highBytes)<<32 | uint64(lowBytes),
	}
}
