package attribute

import (
	"fmt"

	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type LineNumberTableEntry struct {
	StartPc    uint16
	LineNumber uint16
}

func ParseLineNumberTableEntry(fieldPrefix string, binaryReader *reader.BinaryReader) LineNumberTableEntry {
	return LineNumberTableEntry{
		StartPc:    binaryReader.ReadUint16(fmt.Sprintf("%s.StartPc", fieldPrefix)),
		LineNumber: binaryReader.ReadUint16(fmt.Sprintf("%s.LineNumber", fieldPrefix)),
	}
}

type LineNumberTableAttribute struct {
	LineNumberTable []LineNumberTableEntry
}

func (l *LineNumberTableAttribute) Name() AttributeName {
	return LineNumberTableAttributeName
}

func (l *LineNumberTableAttribute) String() string {
	return fmt.Sprintf("LineNumberTableAttribute(LineNumberTable: %v)", l.LineNumberTable)
}

func ParseAttributeLineNumberTable(fieldPrefix string, attributeName *constantpool.Utf8Constant, attributeLength uint32, binaryReader *reader.BinaryReader, constantPool *constantpool.ConstantPool, index int) *LineNumberTableAttribute {
	numberOfEntries := binaryReader.ReadUint16(fmt.Sprintf("%s.NumberOfEntries", fieldPrefix))
	entries := make([]LineNumberTableEntry, numberOfEntries)
	for i := 0; i < int(numberOfEntries); i++ {
		entries[i] = ParseLineNumberTableEntry(fmt.Sprintf("%s.Entries[%d]", fieldPrefix, i), binaryReader)
	}

	return &LineNumberTableAttribute{LineNumberTable: entries}
}
