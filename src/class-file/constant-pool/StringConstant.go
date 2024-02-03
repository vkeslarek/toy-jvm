package constantpool

import (
	"fmt"

	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type StringConstant struct {
	Value ConstantReference[*Utf8Constant]
}

func (constant *StringConstant) ResolveReferences(cp *ConstantPool) {
	constant.Value.ResolveReferences(cp)
}

func (constant StringConstant) String() string {
	return fmt.Sprintf("StringConstant(String: %s)", constant.Value.String())
}

func (constant StringConstant) Tag() ConstantPoolInfoTag {
	return CONSTANT_String
}

func ParseConstantString(binaryReader *reader.BinaryReader, index int) *StringConstant {
	return &StringConstant{
		Value: ParseConstantReference[*Utf8Constant](binaryReader, index),
	}
}
