package attribute

import (
	"fmt"

	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
	"github.com/vkeslarek/toy-jvm/class-file/version"
)

type AttributePermittedLocation string
type AttributePermittedLocations []AttributePermittedLocation

const (
	PermittedLocation_ClassFile           = AttributePermittedLocation("ClassFile")
	PermittedLocation_FieldInfo           = AttributePermittedLocation("field_info")
	PermittedLocation_MethodInfo          = AttributePermittedLocation("method_info")
	PermittedLocation_Code                = AttributePermittedLocation("Code")
	PermittedLocation_RecordComponentInfo = AttributePermittedLocation("record_component_info")
)

type AttributeName struct {
	Name               string
	MinimalJavaVersion version.JavaVersion
	PermittedLocations AttributePermittedLocations
}

var (
	ConstantValueAttributeName                        = AttributeName{"ConstantValue", version.JavaVersion1_0_2, []AttributePermittedLocation{PermittedLocation_FieldInfo}}
	CodeAttributeName                                 = AttributeName{"Code", version.JavaVersion1_0_2, []AttributePermittedLocation{PermittedLocation_MethodInfo}}
	StackMapTableAttributeName                        = AttributeName{"StackMapTable", version.JavaVersion6, []AttributePermittedLocation{PermittedLocation_Code}}
	ExceptionsAttributeName                           = AttributeName{"Exceptions", version.JavaVersion1_0_2, []AttributePermittedLocation{PermittedLocation_MethodInfo}}
	InnerClassesAttributeName                         = AttributeName{"InnerClasses", version.JavaVersion1_1, []AttributePermittedLocation{PermittedLocation_ClassFile}}
	EnclosingMethodAttributeName                      = AttributeName{"EnclosingMethod", version.JavaVersion5_0, []AttributePermittedLocation{PermittedLocation_ClassFile}}
	SyntheticAttributeName                            = AttributeName{"Synthetic", version.JavaVersion1_1, []AttributePermittedLocation{PermittedLocation_ClassFile, PermittedLocation_FieldInfo, PermittedLocation_MethodInfo}}
	SignatureAttributeName                            = AttributeName{"Signature", version.JavaVersion5_0, []AttributePermittedLocation{PermittedLocation_ClassFile, PermittedLocation_FieldInfo, PermittedLocation_MethodInfo, PermittedLocation_RecordComponentInfo}}
	SourceFileAttributeName                           = AttributeName{"SourceFile", version.JavaVersion1_0_2, []AttributePermittedLocation{PermittedLocation_ClassFile}}
	SourceDebugExtensionAttributeName                 = AttributeName{"SourceDebugExtension", version.JavaVersion5_0, []AttributePermittedLocation{PermittedLocation_ClassFile}}
	LineNumberTableAttributeName                      = AttributeName{"LineNumberTable", version.JavaVersion1_0_2, []AttributePermittedLocation{PermittedLocation_Code}}
	LocalVariableTableAttributeName                   = AttributeName{"LocalVariableTable", version.JavaVersion1_0_2, []AttributePermittedLocation{PermittedLocation_Code}}
	LocalVariableTypeTableAttributeName               = AttributeName{"LocalVariableTypeTable", version.JavaVersion5_0, []AttributePermittedLocation{PermittedLocation_Code}}
	DeprecatedAttributeName                           = AttributeName{"Deprecated", version.JavaVersion1_1, []AttributePermittedLocation{PermittedLocation_ClassFile, PermittedLocation_FieldInfo, PermittedLocation_MethodInfo}}
	RuntimeVisibleAnnotationsAttributeName            = AttributeName{"RuntimeVisibleAnnotations", version.JavaVersion5_0, []AttributePermittedLocation{PermittedLocation_ClassFile, PermittedLocation_FieldInfo, PermittedLocation_MethodInfo, PermittedLocation_RecordComponentInfo}}
	RuntimeInvisibleAnnotationsAttributeName          = AttributeName{"RuntimeInvisibleAnnotations", version.JavaVersion5_0, []AttributePermittedLocation{PermittedLocation_ClassFile, PermittedLocation_FieldInfo, PermittedLocation_MethodInfo, PermittedLocation_RecordComponentInfo}}
	RuntimeVisibleParameterAnnotationsAttributeName   = AttributeName{"RuntimeVisibleParameterAnnotations", version.JavaVersion5_0, []AttributePermittedLocation{PermittedLocation_MethodInfo}}
	RuntimeInvisibleParameterAnnotationsAttributeName = AttributeName{"RuntimeInvisibleParameterAnnotations", version.JavaVersion5_0, []AttributePermittedLocation{PermittedLocation_MethodInfo}}
	RuntimeVisibleTypeAnnotationsAttributeName        = AttributeName{"RuntimeVisibleTypeAnnotations", version.JavaVersion8, []AttributePermittedLocation{PermittedLocation_ClassFile, PermittedLocation_FieldInfo, PermittedLocation_MethodInfo, PermittedLocation_Code, PermittedLocation_RecordComponentInfo}}
	RuntimeInvisibleTypeAnnotationsAttributeName      = AttributeName{"RuntimeInvisibleTypeAnnotations", version.JavaVersion8, []AttributePermittedLocation{PermittedLocation_ClassFile, PermittedLocation_FieldInfo, PermittedLocation_MethodInfo, PermittedLocation_Code, PermittedLocation_RecordComponentInfo}}
	AnnotationDefaultAttributeName                    = AttributeName{"AnnotationDefault", version.JavaVersion5_0, []AttributePermittedLocation{PermittedLocation_MethodInfo}}
	BootstrapMethodsAttributeName                     = AttributeName{"BootstrapMethods", version.JavaVersion7, []AttributePermittedLocation{PermittedLocation_ClassFile}}
	MethodParametersAttributeName                     = AttributeName{"MethodParameters", version.JavaVersion8, []AttributePermittedLocation{PermittedLocation_MethodInfo}}
	ModuleAttributeName                               = AttributeName{"Module", version.JavaVersion9, []AttributePermittedLocation{PermittedLocation_ClassFile}}
	ModulePackagesAttributeName                       = AttributeName{"ModulePackages", version.JavaVersion9, []AttributePermittedLocation{PermittedLocation_ClassFile}}
	ModuleMainClassAttributeName                      = AttributeName{"ModuleMainClass", version.JavaVersion9, []AttributePermittedLocation{PermittedLocation_ClassFile}}
	NestHostAttributeName                             = AttributeName{"NestHost", version.JavaVersion11, []AttributePermittedLocation{PermittedLocation_ClassFile}}
	NestMembersAttributeName                          = AttributeName{"NestMembers", version.JavaVersion11, []AttributePermittedLocation{PermittedLocation_ClassFile}}
	RecordAttributeName                               = AttributeName{"Record", version.JavaVersion16, []AttributePermittedLocation{PermittedLocation_ClassFile}}
	PermittedSubclassesAttributeName                  = AttributeName{"PermittedSubclasses", version.JavaVersion17, []AttributePermittedLocation{PermittedLocation_ClassFile}}
)

var AttributeNames = map[string]AttributeName{
	"ConstantValue":                        ConstantValueAttributeName,
	"Code":                                 CodeAttributeName,
	"StackMapTable":                        StackMapTableAttributeName,
	"Exceptions":                           ExceptionsAttributeName,
	"InnerClasses":                         InnerClassesAttributeName,
	"EnclosingMethod":                      EnclosingMethodAttributeName,
	"Synthetic":                            SyntheticAttributeName,
	"Signature":                            SignatureAttributeName,
	"SourceFile":                           SourceFileAttributeName,
	"SourceDebugExtension":                 SourceDebugExtensionAttributeName,
	"LineNumberTable":                      LineNumberTableAttributeName,
	"LocalVariableTable":                   LocalVariableTableAttributeName,
	"LocalVariableTypeTable":               LocalVariableTypeTableAttributeName,
	"Deprecated":                           DeprecatedAttributeName,
	"RuntimeVisibleAnnotations":            RuntimeVisibleAnnotationsAttributeName,
	"RuntimeInvisibleAnnotations":          RuntimeInvisibleAnnotationsAttributeName,
	"RuntimeVisibleParameterAnnotations":   RuntimeVisibleParameterAnnotationsAttributeName,
	"RuntimeInvisibleParameterAnnotations": RuntimeInvisibleParameterAnnotationsAttributeName,
	"RuntimeVisibleTypeAnnotations":        RuntimeVisibleTypeAnnotationsAttributeName,
	"RuntimeInvisibleTypeAnnotations":      RuntimeInvisibleTypeAnnotationsAttributeName,
	"AnnotationDefault":                    AnnotationDefaultAttributeName,
	"BootstrapMethods":                     BootstrapMethodsAttributeName,
	"MethodParameters":                     MethodParametersAttributeName,
	"Module":                               ModuleAttributeName,
	"ModulePackages":                       ModulePackagesAttributeName,
	"ModuleMainClass":                      ModuleMainClassAttributeName,
	"NestHost":                             NestHostAttributeName,
	"NestMembers":                          NestMembersAttributeName,
	"Record":                               RecordAttributeName,
	"PermittedSubclasses":                  PermittedSubclassesAttributeName,
}

type Attribute interface {
	Name() AttributeName
	String() string
}

type AttributeInfo struct {
	Name   *constantpool.Utf8Constant
	Length uint32
	Info   Attribute
}

func ParseAttributeInfo(fieldPrefix string, binaryReader *reader.BinaryReader, constantPool *constantpool.ConstantPool, index int) AttributeInfo {
	attributeName := constantPool.GetUtf8(binaryReader.ReadUint16(fmt.Sprintf("%s.Name", fieldPrefix)))
	attributeLength := binaryReader.ReadUint32(fmt.Sprintf("%s.Length", fieldPrefix))
	return AttributeInfo{
		Name:   attributeName,
		Length: attributeLength,
		Info:   ParseAttribute(fmt.Sprintf("%s.Info(%s)", fieldPrefix, attributeName.Value), attributeName, attributeLength, binaryReader, constantPool, index),
	}
}

func ParseAttribute(fieldPrefix string, attributeName *constantpool.Utf8Constant, attributeLength uint32, binaryReader *reader.BinaryReader, constantPool *constantpool.ConstantPool, index int) Attribute {
	switch attributeName.Value {
	case ConstantValueAttributeName.Name:
		return ParseAttributeConstantValue(fieldPrefix, attributeName, attributeLength, binaryReader, constantPool, index)
	case CodeAttributeName.Name:
		return ParseCodeAttribute(fieldPrefix, attributeName, attributeLength, binaryReader, constantPool, index)
	case StackMapTableAttributeName.Name:
		return ParseStackMapTableAttribute(fieldPrefix, attributeName, attributeLength, binaryReader, constantPool, index)
	case ExceptionsAttributeName.Name:
		return ParseAttributeExceptions(fieldPrefix, attributeName, attributeLength, binaryReader, constantPool, index)
	case InnerClassesAttributeName.Name:
		return ParseAttributeInnerClasses(fieldPrefix, attributeName, attributeLength, binaryReader, constantPool, index)
	case EnclosingMethodAttributeName.Name:
		return ParseAttributeEnclosingMethod(fieldPrefix, attributeName, attributeLength, binaryReader, constantPool, index)
	case SyntheticAttributeName.Name:
		return ParseAttributeSynthetic(fieldPrefix, attributeName, attributeLength, binaryReader, constantPool, index)
	case SignatureAttributeName.Name:
		return ParseAttributeSignature(fieldPrefix, attributeName, attributeLength, binaryReader, constantPool, index)
	case SourceFileAttributeName.Name:
		return ParseAttributeSourceFile(fieldPrefix, attributeName, attributeLength, binaryReader, constantPool, index)
	case SourceDebugExtensionAttributeName.Name:
		return ParseAttributeSourceDebugExtension(fieldPrefix, attributeName, attributeLength, binaryReader, constantPool, index)
	case LineNumberTableAttributeName.Name:
		return ParseAttributeLineNumberTable(fieldPrefix, attributeName, attributeLength, binaryReader, constantPool, index)
	case LocalVariableTableAttributeName.Name:
		return ParseAttributeLocalVariableTable(fieldPrefix, attributeName, attributeLength, binaryReader, constantPool, index)
	case LocalVariableTypeTableAttributeName.Name:
		return ParseAttributeLocalVariableTypeTable(fieldPrefix, attributeName, attributeLength, binaryReader, constantPool, index)
	case DeprecatedAttributeName.Name:
		return ParseAttributeDeprecated(fieldPrefix, attributeName, attributeLength, binaryReader, constantPool, index)
	case RuntimeVisibleAnnotationsAttributeName.Name:
		return ParseAttributeRuntimeVisibleAnnotations(fieldPrefix, attributeName, attributeLength, binaryReader, constantPool, index)
	case RuntimeInvisibleAnnotationsAttributeName.Name:
		return ParseAttributeRuntimeInvisibleAnnotations(fieldPrefix, attributeName, attributeLength, binaryReader, constantPool, index)
	case RuntimeVisibleParameterAnnotationsAttributeName.Name:
		return ParseAttributeRuntimeVisibleParameterAnnotations(fieldPrefix, attributeName, attributeLength, binaryReader, constantPool, index)
	case RuntimeInvisibleParameterAnnotationsAttributeName.Name:
		return ParseAttributeRuntimeInvisibleParameterAnnotations(fieldPrefix, attributeName, attributeLength, binaryReader, constantPool, index)
	case RuntimeVisibleTypeAnnotationsAttributeName.Name:
		return ParseAttributeRuntimeVisibleTypeAnnotations(fieldPrefix, attributeName, attributeLength, binaryReader, constantPool, index)
	case RuntimeInvisibleTypeAnnotationsAttributeName.Name:
		return ParseAttributeRuntimeInvisibleTypeAnnotations(fieldPrefix, attributeName, attributeLength, binaryReader, constantPool, index)
	case AnnotationDefaultAttributeName.Name:
		return ParseAttributeAnnotationDefault(fieldPrefix, attributeName, attributeLength, binaryReader, constantPool, index)
	case BootstrapMethodsAttributeName.Name:
		return ParseAttributeBootstrapMethods(fieldPrefix, attributeName, attributeLength, binaryReader, constantPool, index)
	case MethodParametersAttributeName.Name:
		return ParseAttributeMethodParameters(fieldPrefix, attributeName, attributeLength, binaryReader, constantPool, index)
	// case ModuleAttributeName.Name:
	// 	return ParseAttributeModule(fieldPrefix, attributeName, attributeLength, binaryReader, constantPool, index)
	// case ModulePackagesAttributeName.Name:
	// 	return ParseAttributeModulePackages(fieldPrefix, attributeName, attributeLength, binaryReader, constantPool, index)
	// case ModuleMainClassAttributeName.Name:
	// 	return ParseAttributeModuleMainClass(fieldPrefix, attributeName, attributeLength, binaryReader, constantPool, index)
	// case NestHostAttributeName.Name:
	// 	return ParseAttributeNestHost(fieldPrefix, attributeName, attributeLength, binaryReader, constantPool, index)
	// case NestMembersAttributeName.Name:
	// 	return ParseAttributeNestMembers(fieldPrefix, attributeName, attributeLength, binaryReader, constantPool, index)
	// case RecordAttributeName.Name:
	// 	return ParseAttributeRecord(fieldPrefix, attributeName, attributeLength, binaryReader, constantPool, index)
	// case PermittedSubclassesAttributeName.Name:
	// 	return ParseAttributePermittedSubclasses(fieldPrefix, attributeName, attributeLength, binaryReader, constantPool, index)
	default:
		return ParseCustomAttribute(fieldPrefix, attributeName, attributeLength, binaryReader, constantPool, index)
	}
}
