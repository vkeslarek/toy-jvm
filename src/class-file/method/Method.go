package method

import (
	"fmt"

	"github.com/vkeslarek/toy-jvm/class-file/attribute"
	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
	"github.com/vkeslarek/toy-jvm/types"
)

var MethodAccessFlags = map[uint16]types.AccessFlagName{
	0x0001: types.AccessFlagPublic,
	0x0002: types.AccessFlagPrivate,
	0x0004: types.AccessFlagProtected,
	0x0008: types.AccessFlagStatic,
	0x0010: types.AccessFlagFinal,
	0x0020: types.AccessFlagSynchronized,
	0x0040: types.AccessFlagBridge,
	0x0080: types.AccessFlagVarargs,
	0x0100: types.AccessFlagNative,
	0x0400: types.AccessFlagAbstract,
	0x0800: types.AccessFlagStrict,
	0x1000: types.AccessFlagSynthetic,
}

type Method struct {
	AccessFlags types.BitFlags[uint16]
	Name        *constantpool.Utf8Constant
	Descriptor  *constantpool.Utf8Constant
	Attributes  attribute.Attributes
}

func (m *Method) String() string {
	return fmt.Sprintf("Method(AccessFlags: %s, Name: %s, Descriptor: %s, Attributes: %s)", m.AccessFlags.String(), m.Name, m.Descriptor, m.Attributes)
}

func ParseMethod(binaryReader *reader.BinaryReader, constantPool *constantpool.ConstantPool, index int) *Method {
	return &Method{
		AccessFlags: types.NewBitFlags[uint16](binaryReader.ReadUint16("$.AccessFlags"), MethodAccessFlags),
		Name:        constantPool.GetUtf8(binaryReader.ReadUint16("$.Name")),
		Descriptor:  constantPool.GetUtf8(binaryReader.ReadUint16("$.Descriptor")),
		Attributes:  attribute.ParseAttributes(fmt.Sprintf("$.Methods[%d]", index), binaryReader, constantPool),
	}
}
