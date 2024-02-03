package attribute

import (
	"fmt"

	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type LocalVariableTypeTableEntry struct {
	StartPc   uint16
	Length    uint16
	Name      *constantpool.Utf8Constant
	Signature *constantpool.Utf8Constant
	Index     uint16
}

func ParseLocalVariableTypeTableEntry(fieldPrefix string, binaryReader *reader.BinaryReader, constantPool *constantpool.ConstantPool) LocalVariableTypeTableEntry {
	return LocalVariableTypeTableEntry{
		StartPc:   binaryReader.ReadUint16(fmt.Sprintf("%s.StartPc", fieldPrefix)),
		Length:    binaryReader.ReadUint16(fmt.Sprintf("%s.Length", fieldPrefix)),
		Name:      constantPool.GetUtf8(binaryReader.ReadUint16(fmt.Sprintf("%s.Name", fieldPrefix))),
		Signature: constantPool.GetUtf8(binaryReader.ReadUint16(fmt.Sprintf("%s.Signature", fieldPrefix))),
		Index:     binaryReader.ReadUint16(fmt.Sprintf("%s.Index", fieldPrefix)),
	}
}

type LocalVariableTypeTableAttribute struct {
	Table []LocalVariableTypeTableEntry
}

func (l *LocalVariableTypeTableAttribute) String() string {
	return fmt.Sprintf("LocalVariableTypeTableAttribute(Table: %v)", l.Table)
}

func (l *LocalVariableTypeTableAttribute) Name() AttributeName {
	return LocalVariableTypeTableAttributeName
}

func ParseAttributeLocalVariableTypeTable(fieldPrefix string, attributeName *constantpool.Utf8Constant, attributeLength uint32, binaryReader *reader.BinaryReader, constantPool *constantpool.ConstantPool, index int) *LocalVariableTypeTableAttribute {
	localVariableTypeTableAttribute := &LocalVariableTypeTableAttribute{}
	localVariableTypeTableAttribute.Table = make([]LocalVariableTypeTableEntry, binaryReader.ReadUint16(fmt.Sprintf("%s.TableLength", fieldPrefix)))

	for i := 0; i < len(localVariableTypeTableAttribute.Table); i++ {
		localVariableTypeTableAttribute.Table[i] = ParseLocalVariableTypeTableEntry(fmt.Sprintf("%s.Table[%d]", fieldPrefix, i), binaryReader, constantPool)
	}

	return localVariableTypeTableAttribute
}
