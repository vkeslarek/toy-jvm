package constantpool

import (
	"github.com/vkeslarek/toy-jvm/class-file/reader"
	"github.com/vkeslarek/toy-jvm/class-file/version"
)

type ConstantPoolInfoTag struct {
	Value          uint8
	MinimalVersion version.JavaVersion
	LoadableSince  version.JavaVersion
}

func TagFromValue(value uint8) ConstantPoolInfoTag {
	switch value {
	case 1:
		return CONSTANT_Utf8
	case 3:
		return CONSTANT_Integer
	case 4:
		return CONSTANT_Float
	case 5:
		return CONSTANT_Long
	case 6:
		return CONSTANT_Double
	case 7:
		return CONSTANT_Class
	case 8:
		return CONSTANT_String
	case 9:
		return CONSTANT_Fieldref
	case 10:
		return CONSTANT_Methodref
	case 11:
		return CONSTANT_InterfaceMethodref
	case 12:
		return CONSTANT_NameAndType
	case 15:
		return CONSTANT_MethodHandle
	case 16:
		return CONSTANT_MethodType
	case 17:
		return CONSTANT_Dynamic
	case 18:
		return CONSTANT_InvokeDynamic
	case 19:
		return CONSTANT_Module
	case 20:
		return CONSTANT_Package
	default:
		return CONSTANT_Undefined
	}
}

var (
	CONSTANT_Undefined          ConstantPoolInfoTag = ConstantPoolInfoTag{0, version.JavaVersionUndefined, version.JavaVersionUndefined}
	CONSTANT_Utf8               ConstantPoolInfoTag = ConstantPoolInfoTag{1, version.JavaVersion1_0_2, version.JavaVersionUndefined}
	CONSTANT_Integer            ConstantPoolInfoTag = ConstantPoolInfoTag{3, version.JavaVersion1_0_2, version.JavaVersion1_0_2}
	CONSTANT_Float              ConstantPoolInfoTag = ConstantPoolInfoTag{4, version.JavaVersion1_0_2, version.JavaVersion1_0_2}
	CONSTANT_Long               ConstantPoolInfoTag = ConstantPoolInfoTag{5, version.JavaVersion1_0_2, version.JavaVersion1_0_2}
	CONSTANT_Double             ConstantPoolInfoTag = ConstantPoolInfoTag{6, version.JavaVersion1_0_2, version.JavaVersion1_0_2}
	CONSTANT_Class              ConstantPoolInfoTag = ConstantPoolInfoTag{7, version.JavaVersion1_0_2, version.JavaVersion5_0}
	CONSTANT_String             ConstantPoolInfoTag = ConstantPoolInfoTag{8, version.JavaVersion1_0_2, version.JavaVersion1_0_2}
	CONSTANT_Fieldref           ConstantPoolInfoTag = ConstantPoolInfoTag{9, version.JavaVersion1_0_2, version.JavaVersionUndefined}
	CONSTANT_Methodref          ConstantPoolInfoTag = ConstantPoolInfoTag{10, version.JavaVersion1_0_2, version.JavaVersionUndefined}
	CONSTANT_InterfaceMethodref ConstantPoolInfoTag = ConstantPoolInfoTag{11, version.JavaVersion1_0_2, version.JavaVersionUndefined}
	CONSTANT_NameAndType        ConstantPoolInfoTag = ConstantPoolInfoTag{12, version.JavaVersion1_0_2, version.JavaVersionUndefined}
	CONSTANT_MethodHandle       ConstantPoolInfoTag = ConstantPoolInfoTag{15, version.JavaVersion7, version.JavaVersion7}
	CONSTANT_MethodType         ConstantPoolInfoTag = ConstantPoolInfoTag{16, version.JavaVersion7, version.JavaVersion7}
	CONSTANT_Dynamic            ConstantPoolInfoTag = ConstantPoolInfoTag{17, version.JavaVersion11, version.JavaVersion11}
	CONSTANT_InvokeDynamic      ConstantPoolInfoTag = ConstantPoolInfoTag{18, version.JavaVersion7, version.JavaVersionUndefined}
	CONSTANT_Module             ConstantPoolInfoTag = ConstantPoolInfoTag{19, version.JavaVersion9, version.JavaVersionUndefined}
	CONSTANT_Package            ConstantPoolInfoTag = ConstantPoolInfoTag{20, version.JavaVersion9, version.JavaVersionUndefined}
)

type Constant interface {
	Tag() ConstantPoolInfoTag
	String() string
	ResolveReferences(cp *ConstantPool)
}

type ConstantPoolInfo struct {
	Tag  ConstantPoolInfoTag
	Info Constant
}

func ParseConstantPoolInfo(binaryReader *reader.BinaryReader, index int) *ConstantPoolInfo {
	tag := TagFromValue(binaryReader.ReadUint8("tag"))
	var constant Constant
	switch tag {
	case CONSTANT_Utf8:
		constant = ParseConstantUtf8(binaryReader, index)
	case CONSTANT_Integer:
		constant = ParseConstantInteger(binaryReader, index)
	case CONSTANT_Float:
		constant = ParseConstantFloat(binaryReader, index)
	case CONSTANT_Long:
		constant = ParseConstantLong(binaryReader, index)
	case CONSTANT_Double:
		constant = ParseConstantDouble(binaryReader, index)
	case CONSTANT_Class:
		constant = ParseConstantClass(binaryReader, index)
	case CONSTANT_String:
		constant = ParseConstantString(binaryReader, index)
	case CONSTANT_Fieldref:
		constant = ParseConstantFieldref(binaryReader, index)
	case CONSTANT_Methodref:
		constant = ParseConstantMethodref(binaryReader, index)
	case CONSTANT_InterfaceMethodref:
		constant = ParseConstantInterfaceMethodref(binaryReader, index)
	case CONSTANT_NameAndType:
		constant = ParseConstantNameAndType(binaryReader, index)
	case CONSTANT_MethodHandle:
		constant = ParseConstantMethodHandle(binaryReader, index)
	case CONSTANT_MethodType:
		constant = ParseConstantMethodType(binaryReader, index)
	case CONSTANT_Dynamic:
		constant = ParseConstantDynamic(binaryReader, index)
	case CONSTANT_InvokeDynamic:
		constant = ParseConstantInvokeDynamic(binaryReader, index)
	case CONSTANT_Module:
		constant = ParseConstantModule(binaryReader, index)
	case CONSTANT_Package:
		constant = ParseConstantPackage(binaryReader, index)
	default:
		constant = nil
	}

	return &ConstantPoolInfo{
		Tag:  tag,
		Info: constant,
	}
}
