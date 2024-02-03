package attribute

import (
	"fmt"

	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type BootstrapMethod struct {
	BootstrapMetodRef *constantpool.MethodHandleConstant
	Arguments         []constantpool.Constant
}

func ParseBootstrapMethod(binaryReader *reader.BinaryReader, cp *constantpool.ConstantPool) BootstrapMethod {
	methodRef := cp.GetMethodHandle(binaryReader.ReadUint16("BootstrapMethod.BootstrapMetodRef"))
	numBootstrapArguments := binaryReader.ReadUint16("BootstrapMethod.BootstrapArguments")
	bootstrapArguments := make([]constantpool.Constant, numBootstrapArguments)

	for i := 0; i < int(numBootstrapArguments); i++ {
		bootstrapArguments[i] = cp.Get(int(binaryReader.ReadUint16("BootstrapMethod.BootstrapArguments")))
	}

	return BootstrapMethod{
		BootstrapMetodRef: methodRef,
		Arguments:         bootstrapArguments,
	}
}

func (b *BootstrapMethod) String() string {
	return fmt.Sprintf("BootstrapMethod(BootstrapMetodRef: %s, Arguments: %s)", b.BootstrapMetodRef, b.Arguments)
}

type BootstrapMethodsAttribute struct {
	BootstrapMethods []BootstrapMethod
}

func (b *BootstrapMethodsAttribute) String() string {
	return fmt.Sprintf("BootstrapMethodsAttribute(BootstrapMethods: %s)", b.BootstrapMethods)
}

func (b *BootstrapMethodsAttribute) Name() AttributeName {
	return BootstrapMethodsAttributeName
}

func ParseAttributeBootstrapMethods(fieldPrefix string, attributeName *constantpool.Utf8Constant, attributeLength uint32, binaryReader *reader.BinaryReader, constantPool *constantpool.ConstantPool, index int) *BootstrapMethodsAttribute {
	numBootstrapMethods := binaryReader.ReadUint16("BootstrapMethodsAttribute.BootstrapMethods")
	bootstrapMethods := make([]BootstrapMethod, numBootstrapMethods)

	for i := 0; i < int(numBootstrapMethods); i++ {
		bootstrapMethods[i] = ParseBootstrapMethod(binaryReader, constantPool)
	}

	return &BootstrapMethodsAttribute{
		BootstrapMethods: bootstrapMethods,
	}
}
