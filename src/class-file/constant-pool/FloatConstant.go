package constantpool

import (
	"fmt"
	"math"

	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type FloatConstant struct {
	Value uint32
}

func (i *FloatConstant) ResolveReferences(cp *ConstantPool) {
	// Nothing to do here
}

func (i *FloatConstant) String() string {
	return fmt.Sprintf("FloatConstant(Value: %f)", math.Float32frombits(i.Value))
}

func (i *FloatConstant) Tag() ConstantPoolInfoTag {
	return CONSTANT_Float
}

func ParseConstantFloat(binaryReader *reader.BinaryReader, index int) *FloatConstant {
	return &FloatConstant{
		Value: binaryReader.ReadUint32(fmt.Sprintf("$.ConstantPool[%d](FloatConstant).Bytes", index)),
	}
}
