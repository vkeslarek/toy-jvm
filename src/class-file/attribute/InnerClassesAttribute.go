package attribute

import (
	"fmt"

	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
	"github.com/vkeslarek/toy-jvm/types"
)

var InnerClassAccessFlags = map[uint16]types.AccessFlagName{
	0x0001: types.AccessFlagPublic,
	0x0002: types.AccessFlagPrivate,
	0x0004: types.AccessFlagProtected,
	0x0008: types.AccessFlagStatic,
	0x0010: types.AccessFlagFinal,
	0x0200: types.AccessFlagInterface,
	0x0400: types.AccessFlagAbstract,
	0x1000: types.AccessFlagSynthetic,
	0x2000: types.AccessFlagAnnotation,
	0x4000: types.AccessFlagEnum,
}

type InnerClass struct {
	InnerClassInfo        *constantpool.ClassConstant
	OuterClassInfo        *constantpool.ClassConstant
	InnerNameInfo         *constantpool.Utf8Constant
	InnerClassAccessFlags types.BitFlags[uint16]
}

func (*InnerClass) String() string {
	return fmt.Sprintf("InnerClass(InnerClassInfo: %s, OuterClassInfo: %s, InnerNameInfo: %s, InnerClassAccessFlags: %s)")
}

func ParseInnerClass(fieldPrefix string, binaryReader *reader.BinaryReader, constantPool *constantpool.ConstantPool) InnerClass {
	return InnerClass{
		InnerClassInfo:        constantPool.GetClass(binaryReader.ReadUint16(fmt.Sprintf("%s.InnerClassInfo", fieldPrefix))),
		OuterClassInfo:        constantPool.GetClass(binaryReader.ReadUint16(fmt.Sprintf("%s.OuterClassInfo", fieldPrefix))),
		InnerNameInfo:         constantPool.GetUtf8(binaryReader.ReadUint16(fmt.Sprintf("%s.InnerNameInfo", fieldPrefix))),
		InnerClassAccessFlags: types.NewBitFlags[uint16](binaryReader.ReadUint16(fmt.Sprintf("%s.InnerClassAccessFlags", fieldPrefix)), InnerClassAccessFlags),
	}
}

type InnerClassesAttribute struct {
	InnerClasses []InnerClass
}

func (*InnerClassesAttribute) Name() AttributeName {
	return InnerClassesAttributeName
}

func (i *InnerClassesAttribute) String() string {
	return fmt.Sprintf("InnerClassesAttribute(InnerClasses: %v)", i.InnerClasses)
}

func ParseAttributeInnerClasses(fieldPrefix string, attributeName *constantpool.Utf8Constant, attributeLength uint32, binaryReader *reader.BinaryReader, constantPool *constantpool.ConstantPool, index int) Attribute {
	numberOfClasses := binaryReader.ReadUint16(fmt.Sprintf("%s.NumberOfClasses", fieldPrefix))
	innerClasses := make([]InnerClass, numberOfClasses)
	for i := 0; i < int(numberOfClasses); i++ {
		innerClasses[i] = ParseInnerClass(fmt.Sprintf("%s.InnerClasses[%d]", fieldPrefix, i), binaryReader, constantPool)
	}

	return &InnerClassesAttribute{
		InnerClasses: innerClasses,
	}
}
