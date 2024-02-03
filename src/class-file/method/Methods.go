package method

import (
	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type Methods []*Method

func ParseMethods(binaryReader *reader.BinaryReader, constantPool *constantpool.ConstantPool) Methods {
	size := binaryReader.ReadUint16("$.MethodCount")
	methods := make(Methods, size)
	for i := 0; i < int(size); i++ {
		methods[i] = ParseMethod(binaryReader, constantPool, i)
	}
	return methods
}
