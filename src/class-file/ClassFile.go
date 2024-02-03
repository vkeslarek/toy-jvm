package classfile

import (
	"bytes"
	"fmt"

	"github.com/vkeslarek/toy-jvm/class-file/attribute"
	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/field"
	"github.com/vkeslarek/toy-jvm/class-file/interfaces"
	"github.com/vkeslarek/toy-jvm/class-file/method"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
	"github.com/vkeslarek/toy-jvm/class-file/version"
	"github.com/vkeslarek/toy-jvm/types"
)

var ClassAccessFlags = map[uint16]types.AccessFlagName{
	0x0001: types.AccessFlagPublic,
	0x0010: types.AccessFlagFinal,
	0x0020: types.AccessFlagSuper,
	0x0200: types.AccessFlagInterface,
	0x0400: types.AccessFlagAbstract,
	0x1000: types.AccessFlagSynthetic,
	0x2000: types.AccessFlagAnnotation,
	0x4000: types.AccessFlagEnum,
	0x8000: types.AccessFlagModule,
}

type ClassFile struct {
	Magic        uint32
	Version      version.JavaVersion
	ConstantPool constantpool.ConstantPool
	AccessFlags  types.BitFlags[uint16]
	ThisClass    *constantpool.ClassConstant
	SuperClass   *constantpool.ClassConstant
	Interfaces   interfaces.Interfaces
	Fields       field.Fields
	Methods      method.Methods
	Attributes   attribute.Attributes
}

func (cf ClassFile) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("ClassFile: {\n")
	buffer.WriteString(fmt.Sprintf("\tMagic: 0x%X\n", cf.Magic))
	buffer.WriteString(fmt.Sprintf("\tVersion: %s\n", cf.Version))
	buffer.WriteString(fmt.Sprintf("\tConstantPool: %s\n", cf.ConstantPool))
	buffer.WriteString(fmt.Sprintf("\tAccessFlags: %s\n", cf.AccessFlags))
	buffer.WriteString(fmt.Sprintf("\tThisClass: %s\n", cf.ThisClass))
	buffer.WriteString(fmt.Sprintf("\tSuperClass: %s\n", cf.SuperClass))
	buffer.WriteString(fmt.Sprintf("\tInterfaces: %s\n", cf.Interfaces))
	buffer.WriteString(fmt.Sprintf("\tFields: %s\n", cf.Fields))
	buffer.WriteString(fmt.Sprintf("\tMethods: %s\n", cf.Methods))
	buffer.WriteString(fmt.Sprintf("\tAttributes: %s\n", cf.Attributes))

	return buffer.String()
}

func ParseClassFile(binaryReader *reader.BinaryReader) ClassFile {
	classFile := ClassFile{}

	// Header
	classFile.Magic = binaryReader.ReadUint32("$.Magic")
	classFile.Version = version.ParseJavaVersion(binaryReader)

	// Constant Pool
	classFile.ConstantPool = constantpool.Parse(binaryReader)
	classFile.ConstantPool.ResolveReferences()

	// Class info
	classFile.AccessFlags = types.NewBitFlags(binaryReader.ReadUint16("$.AccessFlags"), ClassAccessFlags)
	classFile.ThisClass = classFile.ConstantPool.GetClass(binaryReader.ReadUint16("$.ThisClass"))
	classFile.SuperClass = classFile.ConstantPool.GetClass(binaryReader.ReadUint16("$.SuperClass"))
	classFile.Interfaces = interfaces.ParseInterfaces(binaryReader, &classFile.ConstantPool)
	classFile.Fields = field.ParseFields(binaryReader, &classFile.ConstantPool)
	classFile.Methods = method.ParseMethods(binaryReader, &classFile.ConstantPool)
	classFile.Attributes = attribute.ParseAttributes("$", binaryReader, &classFile.ConstantPool)

	return classFile
}
