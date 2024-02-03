package attribute

import (
	"fmt"

	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type ValueItem string

const (
	ConstValueIndex ValueItem = "const_value_index"
	EnumConstValue  ValueItem = "enum_const_value"
	ClassInfoIndex  ValueItem = "class_info_index"
	AnnotationValue ValueItem = "annotation_value"
	ArrayValue      ValueItem = "array_value"
)

type ConstantType string

const (
	ConstantInteger ConstantType = "CONSTANT_Integer"
	ConstantFloat   ConstantType = "CONSTANT_Float"
	ConstantLong    ConstantType = "CONSTANT_Long"
	ConstantDouble  ConstantType = "CONSTANT_Double"
	ConstantUtf8    ConstantType = "CONSTANT_Utf8"
)

type ElementValueTag struct {
	Type         string
	ValueItem    ValueItem
	ConstantType ConstantType
}

var (
	ElementValueTagByte             ElementValueTag = ElementValueTag{Type: "byte", ValueItem: ConstValueIndex, ConstantType: ConstantInteger}
	ElementValueTagChar             ElementValueTag = ElementValueTag{Type: "char", ValueItem: ConstValueIndex, ConstantType: ConstantInteger}
	ElementValueDouble              ElementValueTag = ElementValueTag{Type: "double", ValueItem: ConstValueIndex, ConstantType: ConstantDouble}
	ElementValueFloat               ElementValueTag = ElementValueTag{Type: "float", ValueItem: ConstValueIndex, ConstantType: ConstantFloat}
	ElementValueInt                 ElementValueTag = ElementValueTag{Type: "int", ValueItem: ConstValueIndex, ConstantType: ConstantInteger}
	ElementValueLong                ElementValueTag = ElementValueTag{Type: "long", ValueItem: ConstValueIndex, ConstantType: ConstantLong}
	ElementValueShort               ElementValueTag = ElementValueTag{Type: "short", ValueItem: ConstValueIndex, ConstantType: ConstantInteger}
	ElementValueBoolean             ElementValueTag = ElementValueTag{Type: "boolean", ValueItem: ConstValueIndex, ConstantType: ConstantInteger}
	ElementValueString              ElementValueTag = ElementValueTag{Type: "String", ValueItem: ConstValueIndex, ConstantType: ConstantUtf8}
	ElementValueEnum                ElementValueTag = ElementValueTag{Type: "Enum class", ValueItem: EnumConstValue}
	ElementValueClass               ElementValueTag = ElementValueTag{Type: "Class", ValueItem: ClassInfoIndex}
	ElementValueAnnotationInterface ElementValueTag = ElementValueTag{Type: "Annotation interface", ValueItem: AnnotationValue}
	ElementValueArray               ElementValueTag = ElementValueTag{Type: "Array type", ValueItem: ArrayValue}
)

var ElementValueTags = map[uint8]ElementValueTag{
	'B': ElementValueTagByte,
	'C': ElementValueTagChar,
	'D': ElementValueDouble,
	'F': ElementValueFloat,
	'I': ElementValueInt,
	'J': ElementValueLong,
	'S': ElementValueShort,
	'Z': ElementValueBoolean,
	's': ElementValueString,
	'e': ElementValueEnum,
	'c': ElementValueClass,
	'@': ElementValueAnnotationInterface,
	'[': ElementValueArray,
}

type ElementValue struct {
	Tag   uint8
	Value ElementValueItem
}

func (e *ElementValue) String() string {
	return fmt.Sprintf("ElementValue(Tag: %d, Value: %s)", e.Tag, e.Value)
}

func ParseElementValue(fieldPrefix string, binaryReader *reader.BinaryReader, cp *constantpool.ConstantPool) ElementValue {
	tag := binaryReader.ReadUint8(fmt.Sprintf("%s.Tag", fieldPrefix))
	return ElementValue{
		Tag:   tag,
		Value: ParseElementValueItem(ElementValueTags[tag].ValueItem, fmt.Sprintf("%s.Value", fieldPrefix), binaryReader, cp),
	}
}

type ElementValueItem interface {
	Tag() ValueItem
	String() string
}

func ParseElementValueItem(valueItem ValueItem, fieldPrefix string, binaryReader *reader.BinaryReader, cp *constantpool.ConstantPool) ElementValueItem {
	switch valueItem {
	case ConstValueIndex:
		return ParseConstValueIndexElementValue(fieldPrefix, binaryReader, cp)
	case EnumConstValue:
		return ParseEnumConstValueElementValue(fieldPrefix, binaryReader, cp)
	case ClassInfoIndex:
		return ParseClassInfoIndexElementValue(fieldPrefix, binaryReader, cp)
	case AnnotationValue:
		return ParseAnnotationValueElementValue(fieldPrefix, binaryReader, cp)
	case ArrayValue:
		return ParseArrayValueElementValue(fieldPrefix, binaryReader, cp)
	default:
		return nil
	}
}

type ConstValueIndexElementValueItem struct {
	Constant constantpool.Constant
}

func ParseConstValueIndexElementValue(fieldPrefix string, binaryReader *reader.BinaryReader, cp *constantpool.ConstantPool) *ConstValueIndexElementValueItem {
	return &ConstValueIndexElementValueItem{
		Constant: cp.Get(int(binaryReader.ReadUint16(fmt.Sprintf("%s.Constant", fieldPrefix)))),
	}
}

func (e *ConstValueIndexElementValueItem) String() string {
	return fmt.Sprintf("ConstValueIndexElementValue(Constant: %s)", e.Constant.String())
}

func (e *ConstValueIndexElementValueItem) Tag() ValueItem {
	return ConstValueIndex
}

type EnumConstValueElementValueItem struct {
	TypeName  *constantpool.Utf8Constant
	ConstName *constantpool.Utf8Constant
}

func ParseEnumConstValueElementValue(fieldPrefix string, binaryReader *reader.BinaryReader, cp *constantpool.ConstantPool) *EnumConstValueElementValueItem {
	return &EnumConstValueElementValueItem{
		TypeName:  cp.GetUtf8(binaryReader.ReadUint16(fmt.Sprintf("%s.TypeName", fieldPrefix))),
		ConstName: cp.GetUtf8(binaryReader.ReadUint16(fmt.Sprintf("%s.ConstName", fieldPrefix))),
	}
}

func (e *EnumConstValueElementValueItem) String() string {
	return fmt.Sprintf("EnumConstValueElementValue(TypeName: %s, ConstName: %s)", e.TypeName.String(), e.ConstName.String())
}

func (e *EnumConstValueElementValueItem) Tag() ValueItem {
	return EnumConstValue
}

type ClassInfoIndexElementValueItem struct {
	ClassInfo *constantpool.Utf8Constant
}

func ParseClassInfoIndexElementValue(fieldPrefix string, binaryReader *reader.BinaryReader, cp *constantpool.ConstantPool) *ClassInfoIndexElementValueItem {
	return &ClassInfoIndexElementValueItem{
		ClassInfo: cp.GetUtf8(binaryReader.ReadUint16(fmt.Sprintf("%s.ClassInfo", fieldPrefix))),
	}
}

func (e *ClassInfoIndexElementValueItem) String() string {
	return fmt.Sprintf("ClassInfoIndexElementValue(ClassInfo: %s)", e.ClassInfo.String())
}

func (e *ClassInfoIndexElementValueItem) Tag() ValueItem {
	return ClassInfoIndex
}

type AnnotationValueElementValueItem struct {
	Annotation Annotation
}

func ParseAnnotationValueElementValue(fieldPrefix string, binaryReader *reader.BinaryReader, cp *constantpool.ConstantPool) *AnnotationValueElementValueItem {
	return &AnnotationValueElementValueItem{
		Annotation: ParseAnnotation(fmt.Sprintf("%s.Annotation", fieldPrefix), binaryReader, cp),
	}
}

func (e *AnnotationValueElementValueItem) String() string {
	return fmt.Sprintf("AnnotationValueElementValue(Annotation: %s)", e.Annotation.String())
}

func (e *AnnotationValueElementValueItem) Tag() ValueItem {
	return AnnotationValue
}

type ArrayValueElementValueItem struct {
	Values []ElementValue
}

func ParseArrayValueElementValue(fieldPrefix string, binaryReader *reader.BinaryReader, cp *constantpool.ConstantPool) *ArrayValueElementValueItem {
	numValues := binaryReader.ReadUint16(fmt.Sprintf("%s.NumValues", fieldPrefix))
	values := make([]ElementValue, numValues)

	for i := 0; i < int(numValues); i++ {
		values[i] = ParseElementValue(fmt.Sprintf("%s.Values[%d]", fieldPrefix, i), binaryReader, cp)
	}

	return &ArrayValueElementValueItem{
		Values: values,
	}
}

func (e *ArrayValueElementValueItem) String() string {
	return fmt.Sprintf("ArrayValueElementValue(Values: %s)", e.Values)
}

func (e *ArrayValueElementValueItem) Tag() ValueItem {
	return ArrayValue
}

type ElementValuePair struct {
	Name  *constantpool.Utf8Constant
	Value ElementValue
}

func ParseElementValuePair(fieldPrefix string, binaryReader *reader.BinaryReader, cp *constantpool.ConstantPool) ElementValuePair {
	return ElementValuePair{
		Name:  cp.GetUtf8(binaryReader.ReadUint16(fmt.Sprintf("%s.Name", fieldPrefix))),
		Value: ParseElementValue(fmt.Sprintf("%s.Value", fieldPrefix), binaryReader, cp),
	}
}

func ParseElementValuePairs(binaryReader *reader.BinaryReader, cp *constantpool.ConstantPool) []ElementValuePair {
	numPairs := binaryReader.ReadUint16("ElementValuePairs.NumPairs")
	pairs := make([]ElementValuePair, numPairs)

	for i := 0; i < int(numPairs); i++ {
		pairs[i] = ParseElementValuePair(fmt.Sprintf("ElementValuePairs.Pairs[%d]", i), binaryReader, cp)
	}

	return pairs
}
