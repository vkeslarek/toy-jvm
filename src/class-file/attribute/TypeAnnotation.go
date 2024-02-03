package attribute

import (
	"fmt"

	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type TypeAnnotationTarget string

const (
	TypeAnnotationTarget_Unknown         = "unknown_target"
	TypeAnnotationTarget_Parameter       = "type_parameter_target"
	TypeAnnotationTarget_Supertype       = "supertype_target"
	TypeAnnotationTarget_Bound           = "type_parameter_bound_target"
	TypeAnnotationTarget_Empty           = "empty_target"
	TypeAnnotationTarget_FormalParameter = "formal_parameter_target"
	TypeAnnotationTarget_Throws          = "throws_target"
	TypeAnnotationTarget_LocalVariable   = "localvar_target"
	TypeAnnotationTarget_Catch           = "catch_parameter_target"
	TypeAnnotationTarget_Offset          = "catch_target"
	TypeAnnotationTarget_TypeArgument    = "type_argument_target"
)

func TypeAnnotationTargetFromValue(value uint8) TypeAnnotationTarget {
	switch value {
	case 0x00:
	case 0x01:
		return TypeAnnotationTarget_Parameter
	case 0x10:
		return TypeAnnotationTarget_Supertype
	case 0x11:
	case 0x12:
		return TypeAnnotationTarget_Bound
	case 0x13:
	case 0x14:
	case 0x15:
		return TypeAnnotationTarget_Empty
	case 0x16:
		return TypeAnnotationTarget_FormalParameter
	case 0x17:
		return TypeAnnotationTarget_Throws
	case 0x40:
	case 0x41:
		return TypeAnnotationTarget_LocalVariable
	case 0x42:
		return TypeAnnotationTarget_Catch
	case 0x43:
	case 0x44:
	case 0x45:
	case 0x46:
		return TypeAnnotationTarget_Offset
	case 0x47:
	case 0x48:
	case 0x49:
	case 0x4A:
	case 0x4B:
		return TypeAnnotationTarget_TypeArgument
	}

	return TypeAnnotationTarget_Unknown
}

type TypeAnnotationLocation string

const (
	TypeAnnotationLocation_ClassFile           = "classfile"
	TypeAnnotationLocation_Code                = "code"
	TypeAnnotationLocation_MethodInfo          = "method_info"
	TypeAnnotationLocation_FieldInfo           = "field_info"
	TypeAnnotationLocation_RecordComponentInfo = "record_component_info"
)

func TypeAnnotationLocationsFromValue(value uint8) []TypeAnnotationLocation {
	switch value {
	case 0x00:
	case 0x10:
	case 0x11:
		return []TypeAnnotationLocation{TypeAnnotationLocation_Code}
	case 0x01:
	case 0x12:
	case 0x14:
	case 0x15:
	case 0x16:
	case 0x17:
		return []TypeAnnotationLocation{TypeAnnotationLocation_MethodInfo}
	case 0x13:
		return []TypeAnnotationLocation{TypeAnnotationLocation_FieldInfo, TypeAnnotationLocation_RecordComponentInfo}
	case 0x40:
	case 0x41:
	case 0x42:
	case 0x43:
	case 0x44:
	case 0x45:
	case 0x46:
	case 0x47:
	case 0x48:
	case 0x49:
	case 0x4A:
	case 0x4B:
		return []TypeAnnotationLocation{TypeAnnotationLocation_Code}
	}

	return []TypeAnnotationLocation{}
}

type TargetInfo interface {
	Target() TypeAnnotationTarget
	String() string
}

func ParseTargetInfo(targetValue uint8, binaryReader *reader.BinaryReader) TargetInfo {
	switch TypeAnnotationTargetFromValue(targetValue) {
	case TypeAnnotationTarget_Parameter:
		return ParseTypeParameterTarget(binaryReader)
	case TypeAnnotationTarget_Supertype:
		return ParseSuperTypeTarget(binaryReader)
	case TypeAnnotationTarget_Bound:
		return ParseBoundTarget(binaryReader)
	case TypeAnnotationTarget_Empty:
		return ParseEmptyTarget()
	case TypeAnnotationTarget_FormalParameter:
		return ParseFormalParameterTarget(binaryReader)
	case TypeAnnotationTarget_Throws:
		return ParseThrowsTarget(binaryReader)
	case TypeAnnotationTarget_LocalVariable:
		return ParseLocalVarTarget(binaryReader)
	case TypeAnnotationTarget_Catch:
		return ParseCatchTarget(binaryReader)
	case TypeAnnotationTarget_Offset:
		return ParseOffsetTarget(binaryReader)
	case TypeAnnotationTarget_TypeArgument:
		return ParseTypeArgumentTarget(binaryReader)
	default:
		return nil
	}
}

type TypeParameterTarget struct {
	Index uint8
}

func (t *TypeParameterTarget) Target() TypeAnnotationTarget {
	return TypeAnnotationTarget_Parameter
}

func (t *TypeParameterTarget) String() string {
	return fmt.Sprintf("TypeParameterTarget(Index: %d)", t.Index)
}

func ParseTypeParameterTarget(binaryReader *reader.BinaryReader) *TypeParameterTarget {
	return &TypeParameterTarget{
		Index: binaryReader.ReadUint8("TypeParameterTarget.Index"),
	}
}

type SuperTypeTarget struct {
	Index uint16
}

func (t *SuperTypeTarget) Target() TypeAnnotationTarget {
	return TypeAnnotationTarget_Supertype
}

func (t *SuperTypeTarget) String() string {
	return fmt.Sprintf("SuperTypeTarget(Index: %d)", t.Index)
}

func ParseSuperTypeTarget(binaryReader *reader.BinaryReader) *SuperTypeTarget {
	return &SuperTypeTarget{
		Index: binaryReader.ReadUint16("SuperTypeTarget.Index"),
	}
}

type BoundTarget struct {
	Index      uint8
	BoundIndex uint8
}

func (t *BoundTarget) Target() TypeAnnotationTarget {
	return TypeAnnotationTarget_Bound
}

func (t *BoundTarget) String() string {
	return fmt.Sprintf("BoundTarget(Index: %d, BoundIndex: %d)", t.Index, t.BoundIndex)
}

func ParseBoundTarget(binaryReader *reader.BinaryReader) *BoundTarget {
	return &BoundTarget{
		Index:      binaryReader.ReadUint8("BoundTarget.Index"),
		BoundIndex: binaryReader.ReadUint8("BoundTarget.BoundIndex"),
	}
}

type EmptyTarget struct{}

func (t *EmptyTarget) Target() TypeAnnotationTarget {
	return TypeAnnotationTarget_Empty
}

func (t *EmptyTarget) String() string {
	return "EmptyTarget()"
}

func ParseEmptyTarget() *EmptyTarget {
	return &EmptyTarget{}
}

type FormalParameterTarget struct {
	Index uint8
}

func (t *FormalParameterTarget) Target() TypeAnnotationTarget {
	return TypeAnnotationTarget_FormalParameter
}

func (t *FormalParameterTarget) String() string {
	return fmt.Sprintf("FormalParameterTarget(Index: %d)", t.Index)
}

func ParseFormalParameterTarget(binaryReader *reader.BinaryReader) *FormalParameterTarget {
	return &FormalParameterTarget{
		Index: binaryReader.ReadUint8("FormalParameterTarget.Index"),
	}
}

type ThrowsTarget struct {
	Index uint16
}

func (t *ThrowsTarget) Target() TypeAnnotationTarget {
	return TypeAnnotationTarget_Throws
}

func (t *ThrowsTarget) String() string {
	return fmt.Sprintf("ThrowsTarget(Index: %d)", t.Index)
}

func ParseThrowsTarget(binaryReader *reader.BinaryReader) *ThrowsTarget {
	return &ThrowsTarget{
		Index: binaryReader.ReadUint16("ThrowsTarget.Index"),
	}
}

type LocalVar struct {
	StartPc uint16
	Length  uint16
	Index   uint16
}

func ParseLocalVar(binaryReader *reader.BinaryReader) *LocalVar {
	return &LocalVar{
		StartPc: binaryReader.ReadUint16("LocalVar.StartPc"),
		Length:  binaryReader.ReadUint16("LocalVar.Length"),
		Index:   binaryReader.ReadUint16("LocalVar.Index"),
	}
}

type LocalVarTarget struct {
	LocalVars []LocalVar
}

func (t *LocalVarTarget) Target() TypeAnnotationTarget {
	return TypeAnnotationTarget_LocalVariable
}

func (t *LocalVarTarget) String() string {
	return fmt.Sprintf("LocalVarTarget(LocalVars: %v)", t.LocalVars)
}

func ParseLocalVarTarget(binaryReader *reader.BinaryReader) *LocalVarTarget {
	localVarsCount := binaryReader.ReadUint16("LocalVarTarget.LocalVarsCount")
	localVars := make([]LocalVar, localVarsCount)
	for i := 0; i < int(localVarsCount); i++ {
		localVars[i] = *ParseLocalVar(binaryReader)
	}
	return &LocalVarTarget{
		LocalVars: localVars,
	}
}

type CatchTarget struct {
	Index uint16
}

func (t *CatchTarget) Target() TypeAnnotationTarget {
	return TypeAnnotationTarget_Catch
}

func (t *CatchTarget) String() string {
	return fmt.Sprintf("CatchTarget(Index: %d)", t.Index)
}

func ParseCatchTarget(binaryReader *reader.BinaryReader) *CatchTarget {
	return &CatchTarget{
		Index: binaryReader.ReadUint16("CatchTarget.Index"),
	}
}

type OffsetTarget struct {
	Offset uint16
}

func (t *OffsetTarget) Target() TypeAnnotationTarget {
	return TypeAnnotationTarget_Offset
}

func (t *OffsetTarget) String() string {
	return fmt.Sprintf("OffsetTarget(Offset: %d)", t.Offset)
}

func ParseOffsetTarget(binaryReader *reader.BinaryReader) *OffsetTarget {
	return &OffsetTarget{
		Offset: binaryReader.ReadUint16("OffsetTarget.Offset"),
	}
}

type TypeArgumentTarget struct {
	Offset uint16
	Index  uint8
}

func (t *TypeArgumentTarget) Target() TypeAnnotationTarget {
	return TypeAnnotationTarget_TypeArgument
}

func (t *TypeArgumentTarget) String() string {
	return fmt.Sprintf("TypeArgumentTarget(Offset: %d, Index: %d)", t.Offset, t.Index)
}

func ParseTypeArgumentTarget(binaryReader *reader.BinaryReader) *TypeArgumentTarget {
	return &TypeArgumentTarget{
		Offset: binaryReader.ReadUint16("TypeArgumentTarget.Offset"),
		Index:  binaryReader.ReadUint8("TypeArgumentTarget.Index"),
	}
}

type Path struct {
	TypePathKind      uint8
	TypeArgumentIndex uint8
}

func ParsePath(binaryReader *reader.BinaryReader) *Path {
	return &Path{
		TypePathKind:      binaryReader.ReadUint8("Path.TypePathKind"),
		TypeArgumentIndex: binaryReader.ReadUint8("Path.TypeArgumentIndex"),
	}
}

type TypePath struct {
	Path []Path
}

func ParseTypePath(binaryReader *reader.BinaryReader) *TypePath {
	pathCount := binaryReader.ReadUint8("TypePath.PathCount")
	path := make([]Path, pathCount)
	for i := 0; i < int(pathCount); i++ {
		path[i] = *ParsePath(binaryReader)
	}
	return &TypePath{
		Path: path,
	}
}

type TypeAnnotation struct {
	TargetValue       uint8
	TargetInfo        TargetInfo
	TypePath          *TypePath
	TypeIndex         uint16
	ElementValuePairs []ElementValuePair
}

func ParseTypeAnnotation(binaryReader *reader.BinaryReader, cp *constantpool.ConstantPool) *TypeAnnotation {
	targetValue := binaryReader.ReadUint8("TypeAnnotation.TargetValue")
	return &TypeAnnotation{
		TargetValue:       targetValue,
		TargetInfo:        ParseTargetInfo(targetValue, binaryReader),
		TypePath:          ParseTypePath(binaryReader),
		TypeIndex:         binaryReader.ReadUint16("TypeAnnotation.TypeIndex"),
		ElementValuePairs: ParseElementValuePairs(binaryReader, cp),
	}
}

func ParseTypeAnnotations(binaryReader *reader.BinaryReader, cp *constantpool.ConstantPool) []TypeAnnotation {
	numTypeAnnotations := binaryReader.ReadUint16("TypeAnnotations.NumTypeAnnotations")
	typeAnnotations := make([]TypeAnnotation, numTypeAnnotations)
	for i := 0; i < int(numTypeAnnotations); i++ {
		typeAnnotations[i] = *ParseTypeAnnotation(binaryReader, cp)
	}
	return typeAnnotations
}
