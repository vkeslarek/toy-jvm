package attribute

import (
	"fmt"

	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type LocalVariableTableEntry struct {
	StartPc    uint16
	Length     uint16
	Name       *constantpool.Utf8Constant
	Descriptor *constantpool.Utf8Constant
	Index      uint16
}

func ParseLocalVariableTableEntry(fieldPrefix string, binaryReader *reader.BinaryReader, constantPool *constantpool.ConstantPool) LocalVariableTableEntry {
	return LocalVariableTableEntry{
		StartPc:    binaryReader.ReadUint16(fmt.Sprintf("%s.StartPc", fieldPrefix)),
		Length:     binaryReader.ReadUint16(fmt.Sprintf("%s.Length", fieldPrefix)),
		Name:       constantPool.GetUtf8(binaryReader.ReadUint16(fmt.Sprintf("%s.Name", fieldPrefix))),
		Descriptor: constantPool.GetUtf8(binaryReader.ReadUint16(fmt.Sprintf("%s.Descriptor", fieldPrefix))),
		Index:      binaryReader.ReadUint16(fmt.Sprintf("%s.Index", fieldPrefix)),
	}
}

type LocalVariableTableAttribute struct {
	Table []LocalVariableTableEntry
}

func (l *LocalVariableTableAttribute) String() string {
	return fmt.Sprintf("LocalVariableTableAttribute(Table: %v)", l.Table)
}

func (*LocalVariableTableAttribute) Name() AttributeName {
	return LocalVariableTableAttributeName
}

func ParseAttributeLocalVariableTable(fieldPrefix string, attributeName *constantpool.Utf8Constant, attributeLength uint32, binaryReader *reader.BinaryReader, constantPool *constantpool.ConstantPool, index int) *LocalVariableTableAttribute {
	localVariableTableLength := binaryReader.ReadUint16("LocalVariableTableLength")
	localVariableTable := make([]LocalVariableTableEntry, int(localVariableTableLength))

	for i := 0; i < int(localVariableTableLength); i++ {
		localVariableTable[i] = ParseLocalVariableTableEntry(fmt.Sprintf("LocalVariableTableEntry[%d]", i), binaryReader, constantPool)
	}

	return &LocalVariableTableAttribute{
		Table: localVariableTable,
	}
}
