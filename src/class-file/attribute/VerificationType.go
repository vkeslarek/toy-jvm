package attribute

import (
	"fmt"

	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type VerificationTypeTag uint8

const (
	VerificationTypeTagTop               = VerificationTypeTag(0)
	VerificationTypeTagInteger           = VerificationTypeTag(1)
	VerificationTypeTagFloat             = VerificationTypeTag(2)
	VerificationTypeTagDouble            = VerificationTypeTag(3)
	VerificationTypeTagLong              = VerificationTypeTag(4)
	VerificationTypeTagNull              = VerificationTypeTag(5)
	VerificationTypeTagUninitializedThis = VerificationTypeTag(6)
	VerificationTypeTagObject            = VerificationTypeTag(7)
	VerificationTypeTagUninitialized     = VerificationTypeTag(8)
)

type VerificationType interface {
	Tag() VerificationTypeTag
	String() string
}

type TopVerificationType struct{}

func ParseTopVerificationType(fieldPrefix string, binaryReader *reader.BinaryReader, cp *constantpool.ConstantPool) *TopVerificationType {
	return &TopVerificationType{}
}

func (*TopVerificationType) Tag() VerificationTypeTag {
	return VerificationTypeTagTop
}

func (*TopVerificationType) String() string {
	return "TopVerificationType()"
}

type IntegerVerificationType struct{}

func ParseIntegerVerificationType(fieldPrefix string, binaryReader *reader.BinaryReader, cp *constantpool.ConstantPool) *IntegerVerificationType {
	return &IntegerVerificationType{}
}

func (*IntegerVerificationType) Tag() VerificationTypeTag {
	return VerificationTypeTagInteger
}

func (*IntegerVerificationType) String() string {
	return "IntegerVerificationType()"
}

type FloatVerificationType struct{}

func ParseFloatVerificationType(fieldPrefix string, binaryReader *reader.BinaryReader, cp *constantpool.ConstantPool) *FloatVerificationType {
	return &FloatVerificationType{}
}

func (*FloatVerificationType) Tag() VerificationTypeTag {
	return VerificationTypeTagFloat
}

func (*FloatVerificationType) String() string {
	return "FloatVerificationType()"
}

type DoubleVerificationType struct{}

func ParseDoubleVerificationType(fieldPrefix string, binaryReader *reader.BinaryReader, cp *constantpool.ConstantPool) *DoubleVerificationType {
	return &DoubleVerificationType{}
}

func (*DoubleVerificationType) Tag() VerificationTypeTag {
	return VerificationTypeTagDouble
}

func (*DoubleVerificationType) String() string {
	return "DoubleVerificationType()"
}

type LongVerificationType struct{}

func ParseLongVerificationType(fieldPrefix string, binaryReader *reader.BinaryReader, cp *constantpool.ConstantPool) *LongVerificationType {
	return &LongVerificationType{}
}

func (*LongVerificationType) Tag() VerificationTypeTag {
	return VerificationTypeTagLong
}

func (*LongVerificationType) String() string {
	return "LongVerificationType()"
}

type NullVerificationType struct{}

func ParseNullVerificationType(fieldPrefix string, binaryReader *reader.BinaryReader, cp *constantpool.ConstantPool) *NullVerificationType {
	return &NullVerificationType{}
}

func (*NullVerificationType) Tag() VerificationTypeTag {
	return VerificationTypeTagNull
}

func (*NullVerificationType) String() string {
	return "NullVerificationType()"
}

type UninitializedThisVerificationType struct{}

func ParseUninitializedThisVerificationType(fieldPrefix string, binaryReader *reader.BinaryReader, cp *constantpool.ConstantPool) *UninitializedThisVerificationType {
	return &UninitializedThisVerificationType{}
}

func (*UninitializedThisVerificationType) Tag() VerificationTypeTag {
	return VerificationTypeTagUninitializedThis
}

func (*UninitializedThisVerificationType) String() string {
	return "UninitializedThisVerificationType()"
}

type ObjectVerificationType struct {
	Object *constantpool.ClassConstant
}

func ParseObjectVerificationType(fieldPrefix string, binaryReader *reader.BinaryReader, cp *constantpool.ConstantPool) *ObjectVerificationType {
	return &ObjectVerificationType{
		Object: cp.GetClass(binaryReader.ReadUint16(fmt.Sprintf("%s.Object", fieldPrefix))),
	}
}

func (object *ObjectVerificationType) Tag() VerificationTypeTag {
	return VerificationTypeTagObject
}

func (object *ObjectVerificationType) String() string {
	return fmt.Sprintf("ObjectVerificationType(Object: %v)", object.Object)
}

type UninitializedVerificationType struct {
	Offset uint16
}

func ParseUninitializedVerificationType(fieldPrefix string, binaryReader *reader.BinaryReader, cp *constantpool.ConstantPool) *UninitializedVerificationType {
	return &UninitializedVerificationType{
		Offset: binaryReader.ReadUint16(fmt.Sprintf("%s.Offset", fieldPrefix)),
	}
}

func (*UninitializedVerificationType) Tag() VerificationTypeTag {
	return VerificationTypeTagUninitialized
}

func (*UninitializedVerificationType) String() string {
	return fmt.Sprintf("UninitializedVerificationType(Offset: %v)", 0)
}

func ParseVerificationType(fieldPrefix string, binaryReader *reader.BinaryReader, cp *constantpool.ConstantPool) VerificationType {
	verificationTypeTag := binaryReader.ReadUint8(fmt.Sprintf("%s.Tag", fieldPrefix))

	switch VerificationTypeTag(verificationTypeTag) {
	case VerificationTypeTagTop:
		return ParseTopVerificationType(fieldPrefix, binaryReader, cp)
	case VerificationTypeTagInteger:
		return ParseIntegerVerificationType(fieldPrefix, binaryReader, cp)
	case VerificationTypeTagFloat:
		return ParseFloatVerificationType(fieldPrefix, binaryReader, cp)
	case VerificationTypeTagDouble:
		return ParseDoubleVerificationType(fieldPrefix, binaryReader, cp)
	case VerificationTypeTagLong:
		return ParseLongVerificationType(fieldPrefix, binaryReader, cp)
	case VerificationTypeTagNull:
		return ParseNullVerificationType(fieldPrefix, binaryReader, cp)
	case VerificationTypeTagUninitializedThis:
		return ParseUninitializedThisVerificationType(fieldPrefix, binaryReader, cp)
	case VerificationTypeTagObject:
		return ParseObjectVerificationType(fieldPrefix, binaryReader, cp)
	case VerificationTypeTagUninitialized:
		return ParseUninitializedVerificationType(fieldPrefix, binaryReader, cp)
	default:
		return nil
	}
}
