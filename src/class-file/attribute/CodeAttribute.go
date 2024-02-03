package attribute

import (
	"encoding/hex"
	"fmt"

	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type ExceptionTableEntry struct {
	StartPc   uint16
	EndPc     uint16
	HandlerPc uint16
	CatchType *constantpool.ClassConstant
}

func (e *ExceptionTableEntry) String() string {
	return fmt.Sprintf("ExceptionTableEntry(StartPc: %d, EndPc: %d, HandlerPc: %d, CatchType: %s)",
		e.StartPc, e.EndPc, e.HandlerPc, e.CatchType.String())
}

func ParseExceptionTableEntry(fieldPrefix string, binaryReader *reader.BinaryReader, cp *constantpool.ConstantPool) ExceptionTableEntry {
	return ExceptionTableEntry{
		StartPc:   binaryReader.ReadUint16(fmt.Sprintf("%s.StartPc", fieldPrefix)),
		EndPc:     binaryReader.ReadUint16(fmt.Sprintf("%s.EndPc", fieldPrefix)),
		HandlerPc: binaryReader.ReadUint16(fmt.Sprintf("%s.HandlerPc", fieldPrefix)),
		CatchType: cp.GetClass(binaryReader.ReadUint16(fmt.Sprintf("%s.CatchType", fieldPrefix))),
	}
}

func ParseExceptionTable(fieldPrefix string, binaryReader *reader.BinaryReader, cp *constantpool.ConstantPool) []ExceptionTableEntry {
	size := binaryReader.ReadUint16(fmt.Sprintf("%s.ExceptionTableCount", fieldPrefix))
	exceptionTable := make([]ExceptionTableEntry, size)

	for i := 0; i < int(size); i++ {
		exceptionTable[i] = ParseExceptionTableEntry(fmt.Sprintf("%s.ExceptionTable[%d]", fieldPrefix, i), binaryReader, cp)
	}

	return exceptionTable
}

type CodeAttribute struct {
	MaxStack       uint16
	MaxLocals      uint16
	Code           []byte
	ExceptionTable []ExceptionTableEntry
	Attributes     Attributes
}

func (c *CodeAttribute) Name() AttributeName {
	return CodeAttributeName
}

func (c *CodeAttribute) String() string {
	return fmt.Sprintf("CodeAttribute(MaxStack: %d, MaxLocals: %d, Code: %s, ExceptionTable: %s, Attributes: %s)",
		c.MaxStack, c.MaxLocals, hex.EncodeToString(c.Code), c.ExceptionTable, c.Attributes)
}

func ParseCodeAttribute(fieldPrefix string, attributeName *constantpool.Utf8Constant, attributeLength uint32, binaryReader *reader.BinaryReader, constantPool *constantpool.ConstantPool, index int) *CodeAttribute {
	return &CodeAttribute{
		MaxStack:       binaryReader.ReadUint16(fmt.Sprintf("%s.MaxStack", fieldPrefix)),
		MaxLocals:      binaryReader.ReadUint16(fmt.Sprintf("%s.MaxLocals", fieldPrefix)),
		Code:           binaryReader.ReadBytes(fmt.Sprintf("%s.Code", fieldPrefix), int(binaryReader.ReadUint32(fmt.Sprintf("%s.CodeLength", fieldPrefix)))),
		ExceptionTable: ParseExceptionTable(fmt.Sprintf("%s.ExceptionTable", fieldPrefix), binaryReader, constantPool),
		Attributes:     ParseAttributes(fieldPrefix, binaryReader, constantPool),
	}
}
