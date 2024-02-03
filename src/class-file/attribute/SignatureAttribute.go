package attribute

import (
	"fmt"

	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type SignatureAttribute struct {
	Signature *constantpool.Utf8Constant
}

func (*SignatureAttribute) Name() AttributeName {
	return SignatureAttributeName
}

func (a *SignatureAttribute) String() string {
	return fmt.Sprintf("Signature: %s", a.Signature)
}

func ParseAttributeSignature(fieldPrefix string, attributeName *constantpool.Utf8Constant, attributeLength uint32, binaryReader *reader.BinaryReader, constantPool *constantpool.ConstantPool, index int) *SignatureAttribute {
	return &SignatureAttribute{
		Signature: constantPool.GetUtf8(binaryReader.ReadUint16("Signature")),
	}
}
