package constantpool

import (
	"bytes"
	"fmt"

	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type ConstantPool struct {
	constantPoolMap map[int]Constant
}

func (cp *ConstantPool) ResolveReferences() {
	for j := 1; j < len(cp.constantPoolMap); j++ {
		cp.constantPoolMap[j].ResolveReferences(cp)
	}
}

func (cp *ConstantPool) Get(index int) Constant {
	if index == 0 || index > len(cp.constantPoolMap) {
		return nil
	}

	return cp.constantPoolMap[index]
}

func (cp *ConstantPool) GetClass(index uint16) *ClassConstant {
	constant, ok := cp.Get(int(index)).(*ClassConstant)
	if !ok {
		return nil
	}

	return constant
}

func (cp *ConstantPool) GetUtf8(index uint16) *Utf8Constant {
	constant, ok := cp.Get(int(index)).(*Utf8Constant)
	if !ok {
		return nil
	}

	return constant
}

func (cp *ConstantPool) GetNameAndType(index uint16) *NameAndTypeConstant {
	constant, ok := cp.Get(int(index)).(*NameAndTypeConstant)
	if !ok {
		return nil
	}

	return constant
}

func (cp *ConstantPool) GetMethodHandle(index uint16) *MethodHandleConstant {
	constant, ok := cp.Get(int(index)).(*MethodHandleConstant)
	if !ok {
		return nil
	}

	return constant
}

func (cp ConstantPool) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("ConstantPool: {\n")
	for i := 1; i < len(cp.constantPoolMap); i++ {
		buffer.WriteString(fmt.Sprintf("\t%d: %s\n", i, cp.constantPoolMap[i]))
	}
	buffer.WriteString("}")

	return buffer.String()
}

func Parse(binaryReader *reader.BinaryReader) ConstantPool {
	size := binaryReader.ReadUint16("$.ConstantPoolCount")
	constantPoolMap := make(map[int]Constant, int(size-1))

	for i := 1; i < int(size); i++ {
		constantPoolInf := ParseConstantPoolInfo(binaryReader, i)
		constantPoolMap[i] = constantPoolInf.Info
	}

	return ConstantPool{
		constantPoolMap: constantPoolMap,
	}
}
