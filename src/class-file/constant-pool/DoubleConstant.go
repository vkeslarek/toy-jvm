package constantpool

import (
	"fmt"
	"math"

	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type DoubleConstant struct {
	Value uint64
}

func (d *DoubleConstant) ResolveReferences(cp *ConstantPool) {
	// Nothing to do here
}

func (d *DoubleConstant) String() string {
	return fmt.Sprintf("DoubleConstant(Value: %f)", math.Float64frombits(d.Value))
}

func (*DoubleConstant) Tag() ConstantPoolInfoTag {
	return CONSTANT_Double
}

func ParseConstantDouble(binaryReader *reader.BinaryReader, index int) *DoubleConstant {
	highBytes := binaryReader.ReadUint32(fmt.Sprintf("$.ConstantPool[%d](DoubleConstant).HighBytes", index))
	lowBytes := binaryReader.ReadUint32(fmt.Sprintf("$.ConstantPool[%d](DoubleConstant).LowBytes", index))

	return &DoubleConstant{
		Value: uint64(highBytes)<<32 | uint64(lowBytes),
	}
}
