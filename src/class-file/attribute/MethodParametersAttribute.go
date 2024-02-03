package attribute

import (
	"fmt"

	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
	"github.com/vkeslarek/toy-jvm/types"
)

var MethodParameterAccessFlags = map[uint16]string{
	0x0010: types.AccessFlagFinal,
	0x1000: types.AccessFlagSynthetic,
	0x8000: types.AccessFlagMandated,
}

type MethodParameter struct {
	Name        constantpool.Utf8Constant
	AccessFlags types.BitFlags[uint16]
}

func (mp *MethodParameter) String() string {
	return fmt.Sprintf("%s %s", mp.AccessFlags.String(), mp.Name.String())
}

func ParseMethodParameter(binaryReader *reader.BinaryReader, constantPool *constantpool.ConstantPool, index int) MethodParameter {
	return MethodParameter{
		Name:        *constantPool.GetUtf8(binaryReader.ReadUint16(fmt.Sprintf("$.MethodParameters[%d].Name", index))),
		AccessFlags: types.NewBitFlags[uint16](binaryReader.ReadUint16(fmt.Sprintf("$.MethodParameters[%d].AccessFlags", index)), MethodParameterAccessFlags),
	}
}

type MethodParametersAttribute struct {
	Parameters []MethodParameter
}

func (mpa *MethodParametersAttribute) String() string {
	return fmt.Sprintf("MethodParametersAttribute(Parameters: %s)", mpa.Parameters)
}

func (mpa *MethodParametersAttribute) Name() AttributeName {
	return MethodParametersAttributeName
}

func ParseAttributeMethodParameters(fieldPrefix string, attributeName *constantpool.Utf8Constant, attributeLength uint32, binaryReader *reader.BinaryReader, constantPool *constantpool.ConstantPool, index int) *MethodParametersAttribute {
	parametersCount := binaryReader.ReadUint8("MethodParametersAttribute.ParametersCount")
	parameters := make([]MethodParameter, parametersCount)

	for i := 0; i < int(parametersCount); i++ {
		parameters[i] = ParseMethodParameter(binaryReader, constantPool, i)
	}

	return &MethodParametersAttribute{
		Parameters: parameters,
	}
}
