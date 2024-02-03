package constantpool

import (
	"fmt"

	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type Utf8Constant struct {
	Value string
}

func (constant *Utf8Constant) ResolveReferences(cp *ConstantPool) {
	// Nothing to do here
}

func (constant Utf8Constant) String() string {
	return fmt.Sprintf("Utf8Constant(Value: %s)", constant.Value)
}

func (constant Utf8Constant) Tag() ConstantPoolInfoTag {
	return CONSTANT_Utf8
}

func ParseConstantUtf8(binaryReader *reader.BinaryReader, index int) *Utf8Constant {
	size := binaryReader.ReadUint16(fmt.Sprintf("$.ConstantPool[%d](Utf8Constant).Lenth", index))

	return &Utf8Constant{
		Value: binaryReader.ReadString(fmt.Sprintf("$.ConstantPool[%d](Utf8Constant).Bytes", index), int(size)),
	}
}
