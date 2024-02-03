package attribute

import (
	"fmt"

	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type StackFrameType uint8

const (
	FrameTypeReserved           StackFrameType = iota
	FrameTypeSame               StackFrameType = iota
	FrameTypeSameLocals         StackFrameType = iota
	FrameTypeSameLocalsExtended StackFrameType = iota
	FrameTypeChop               StackFrameType = iota
	FrameTypeSameExtended       StackFrameType = iota
	FrameTypeAppend             StackFrameType = iota
	FrameTypeFull               StackFrameType = iota
)

func StackFrameTypeFromValue(value uint8) StackFrameType {
	if value >= 0 && value <= 63 {
		return FrameTypeSame
	}

	if value >= 64 && value <= 127 {
		return FrameTypeSameLocals
	}

	if value == 247 {
		return FrameTypeSameLocalsExtended
	}

	if value >= 248 && value <= 250 {
		return FrameTypeChop
	}

	if value == 251 {
		return FrameTypeSameExtended
	}

	if value >= 252 && value <= 254 {
		return FrameTypeAppend
	}

	if value == 255 {
		return FrameTypeFull
	}

	return FrameTypeReserved
}

type StackFrame interface {
	Type() StackFrameType
	String() string
}

type SameFrame struct{}

func (s *SameFrame) String() string {
	return fmt.Sprintf("SameFrame()")
}

func (*SameFrame) Type() StackFrameType {
	return FrameTypeSame
}

func ParseSameFrame(fieldPrefix string, frameTypeValue uint8, frameLength uint8, binaryReader *reader.BinaryReader, cp *constantpool.ConstantPool) StackFrame {
	return &SameFrame{}
}

type SameLocalsFrame struct {
	VerificationType VerificationType
}

func (s *SameLocalsFrame) String() string {
	return fmt.Sprintf("SameLocalsFrame(VerificationType: %v)", s.VerificationType)
}

func (*SameLocalsFrame) Type() StackFrameType {
	return FrameTypeSameLocals
}

func ParseSameLocalsFrame(fieldPrefix string, frameTypeValue uint8, frameLength uint8, binaryReader *reader.BinaryReader, cp *constantpool.ConstantPool) StackFrame {
	return &SameLocalsFrame{
		VerificationType: ParseVerificationType(fmt.Sprintf("%s.VerificationType", fieldPrefix), binaryReader, cp),
	}
}

type SameLocalsExtendedFrame struct {
	Offset            uint16
	VerificationTypes VerificationType
}

func (s *SameLocalsExtendedFrame) String() string {
	return fmt.Sprintf("SameLocalsExtendedFrame(Offset: %v, VerificationType: %v)", s.Offset, s.VerificationTypes)
}

func (*SameLocalsExtendedFrame) Type() StackFrameType {
	return FrameTypeSameLocalsExtended
}

func ParseSameLocalsExtendedFrame(fieldPrefix string, frameTypeValue uint8, frameLength uint8, binaryReader *reader.BinaryReader, cp *constantpool.ConstantPool) StackFrame {
	offset := binaryReader.ReadUint16(fmt.Sprintf("%s.Offset", fieldPrefix))
	return &SameLocalsExtendedFrame{
		Offset:            offset,
		VerificationTypes: ParseVerificationType(fmt.Sprintf("%s.VerificationType", fieldPrefix), binaryReader, cp),
	}
}

type ChopFrame struct {
	Offset uint16
}

func (c *ChopFrame) String() string {
	return fmt.Sprintf("ChopFrame(Offset: %v)", c.Offset)
}

func (*ChopFrame) Type() StackFrameType {
	return FrameTypeChop
}

func ParseChopFrame(fieldPrefix string, frameTypeValue uint8, frameLength uint8, binaryReader *reader.BinaryReader, cp *constantpool.ConstantPool) StackFrame {
	return &ChopFrame{
		Offset: binaryReader.ReadUint16(fmt.Sprintf("%s.Offset", fieldPrefix)),
	}
}

type SameExtendedFrame struct {
	Offset uint16
}

func (s *SameExtendedFrame) String() string {
	return fmt.Sprintf("SameExtendedFrame(Offset: %v)", s.Offset)
}

func (*SameExtendedFrame) Type() StackFrameType {
	return FrameTypeSameExtended
}

func ParseSameExtendedFrame(fieldPrefix string, frameTypeValue uint8, frameLength uint8, binaryReader *reader.BinaryReader, cp *constantpool.ConstantPool) StackFrame {
	return &SameExtendedFrame{
		Offset: binaryReader.ReadUint16(fmt.Sprintf("%s.Offset", fieldPrefix)),
	}
}

type AppendFrame struct {
	Offset            uint16
	VerificationTypes []VerificationType
}

func (a *AppendFrame) String() string {
	return fmt.Sprintf("AppendFrame(Offset: %v, VerificationTypes: %v)", a.Offset, a.VerificationTypes)
}

func (*AppendFrame) Type() StackFrameType {
	return FrameTypeAppend
}

func ParseAppendFrame(fieldPrefix string, frameTypeValue uint8, frameLength uint8, binaryReader *reader.BinaryReader, cp *constantpool.ConstantPool) StackFrame {
	offset := binaryReader.ReadUint16(fmt.Sprintf("%s.Offset", fieldPrefix))
	verificationTypes := make([]VerificationType, frameTypeValue-251)

	for i := 0; i < int(frameTypeValue-251); i++ {
		verificationTypes[i] = ParseVerificationType(fmt.Sprintf("%s.VerificationType[%d]", fieldPrefix, i), binaryReader, cp)
	}

	return &AppendFrame{
		Offset:            offset,
		VerificationTypes: verificationTypes,
	}
}

type FullFrame struct {
	Offset uint16
	Locals []VerificationType
	Stack  []VerificationType
}

func (f *FullFrame) String() string {
	return fmt.Sprintf("FullFrame(Offset: %v, Locals: %v, Stack: %v)", f.Offset, f.Locals, f.Stack)
}

func (*FullFrame) Type() StackFrameType {
	return FrameTypeFull
}

func ParseFullFrame(fieldPrefix string, frameTypeValue uint8, frameLength uint8, binaryReader *reader.BinaryReader, cp *constantpool.ConstantPool) StackFrame {
	offset := binaryReader.ReadUint16(fmt.Sprintf("%s.Offset", fieldPrefix))

	localCount := binaryReader.ReadUint16(fmt.Sprintf("%s.LocalCount", fieldPrefix))
	locals := make([]VerificationType, localCount)
	for i := 0; i < int(localCount); i++ {
		locals[i] = ParseVerificationType(fmt.Sprintf("%s.Local[%d]", fieldPrefix, i), binaryReader, cp)
	}

	stackCount := binaryReader.ReadUint16(fmt.Sprintf("%s.StackCount", fieldPrefix))
	stack := make([]VerificationType, stackCount)
	for i := 0; i < int(stackCount); i++ {
		stack[i] = ParseVerificationType(fmt.Sprintf("%s.Stack[%d]", fieldPrefix, i), binaryReader, cp)
	}

	return &FullFrame{
		Offset: offset,
		Locals: locals,
		Stack:  stack,
	}
}

func ParseStackFrame(fieldPrefix string, frameType StackFrameType, frameTypeValue uint8, frameLength uint8, binaryReader *reader.BinaryReader, cp *constantpool.ConstantPool) StackFrame {
	switch frameType {
	case FrameTypeSame:
		return ParseSameFrame(fieldPrefix, frameTypeValue, frameLength, binaryReader, cp)
	case FrameTypeSameLocals:
		return ParseSameLocalsFrame(fieldPrefix, frameTypeValue, frameLength, binaryReader, cp)
	case FrameTypeSameLocalsExtended:
		return ParseSameLocalsExtendedFrame(fieldPrefix, frameTypeValue, frameLength, binaryReader, cp)
	case FrameTypeChop:
		return ParseChopFrame(fieldPrefix, frameTypeValue, frameLength, binaryReader, cp)
	case FrameTypeSameExtended:
		return ParseSameExtendedFrame(fieldPrefix, frameTypeValue, frameLength, binaryReader, cp)
	case FrameTypeAppend:
		return ParseAppendFrame(fieldPrefix, frameTypeValue, frameLength, binaryReader, cp)
	case FrameTypeFull:
		return ParseFullFrame(fieldPrefix, frameTypeValue, frameLength, binaryReader, cp)
	}

	return nil
}

type StackMapFrame struct {
	Value uint8
	Frame StackFrame
}

func (s *StackMapFrame) String() string {
	return fmt.Sprintf("StackMapFrame(Value: %v, Frame: %v)", s.Value, s.Frame)
}

func ParseStackMapFrame(fieldPrefix string, binaryReader *reader.BinaryReader, cp *constantpool.ConstantPool) StackMapFrame {
	frameTypeValue := binaryReader.ReadUint8(fmt.Sprintf("%s.FrameType", fieldPrefix))
	frameType := StackFrameTypeFromValue(frameTypeValue)
	return StackMapFrame{
		Value: frameTypeValue,
		Frame: ParseStackFrame(fmt.Sprint("%s.Frame", fieldPrefix), frameType, frameTypeValue, binaryReader.ReadUint8(fmt.Sprintf("%s.FrameType")), binaryReader, cp),
	}
}

type StackMapTableAttribute struct {
	Entries []StackMapFrame
}

func (s *StackMapTableAttribute) Name() AttributeName {
	return StackMapTableAttributeName
}

func (s *StackMapTableAttribute) String() string {
	return fmt.Sprintf("StackMapTableAttribute(Entries: %v)", s.Entries)
}

func ParseStackMapTableAttribute(fieldPrefix string, attributeName *constantpool.Utf8Constant, attributeLength uint32, binaryReader *reader.BinaryReader, constantPool *constantpool.ConstantPool, index int) *StackMapTableAttribute {
	numberOfEntries := binaryReader.ReadUint16(fmt.Sprintf("%s.NumberOfEntries", fieldPrefix))
	entries := make([]StackMapFrame, numberOfEntries)
	for i := 0; i < int(numberOfEntries); i++ {
		entries[i] = ParseStackMapFrame(fmt.Sprintf("%s.Entries[%d]", fieldPrefix, i), binaryReader, constantPool)
	}

	return &StackMapTableAttribute{
		Entries: entries,
	}
}
