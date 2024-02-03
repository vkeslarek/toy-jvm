package field

import (
	"fmt"

	"github.com/vkeslarek/toy-jvm/class-file/attribute"
	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
	"github.com/vkeslarek/toy-jvm/types"
)

var FieldAccessFlags = map[uint16]types.AccessFlagName{
	0x0001: types.AccessFlagPublic,
	0x0002: types.AccessFlagPrivate,
	0x0004: types.AccessFlagProtected,
	0x0008: types.AccessFlagStatic,
	0x0010: types.AccessFlagFinal,
	0x0040: types.AccessFlagVolatile,
	0x0080: types.AccessFlagTransient,
	0x1000: types.AccessFlagSynthetic,
	0x4000: types.AccessFlagEnum,
}

type Field struct {
	AccessFlags types.BitFlags[uint16]
	Name        *constantpool.Utf8Constant
	Descriptor  *constantpool.Utf8Constant
	Attributes  attribute.Attributes
}

func (f *Field) String() string {
	return fmt.Sprintf("Field(AccessFlags: %s, Name: %s, Descriptor: %s, Attributes: %s)", f.AccessFlags.String(), f.Name, f.Descriptor, f.Attributes)
}

func ParseField(binaryReader *reader.BinaryReader, constantPool *constantpool.ConstantPool, index int) *Field {
	return &Field{
		AccessFlags: types.NewBitFlags[uint16](binaryReader.ReadUint16(fmt.Sprintf("$.Fields[%d].AccessFlags", index)), FieldAccessFlags),
		Name:        constantPool.GetUtf8(binaryReader.ReadUint16(fmt.Sprintf("$.Fields[%d].Name", index))),
		Descriptor:  constantPool.GetUtf8(binaryReader.ReadUint16(fmt.Sprintf("$.Fields[%d].Descriptor", index))),
		Attributes:  attribute.ParseAttributes(fmt.Sprintf("$.Fields[%d]", index), binaryReader, constantPool),
	}
}
