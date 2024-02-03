package interfaces

import (
	"fmt"

	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type Interfaces []Interface

func ParseInterfaces(binaryReader *reader.BinaryReader, cp *constantpool.ConstantPool) Interfaces {
	size := binaryReader.ReadUint16("$.InterfacesCount")
	interfaces := Interfaces(make([]Interface, size))

	for i := 0; i < int(size); i++ {
		interfaceIndex := binaryReader.ReadUint16(fmt.Sprintf("$.Interfaces[%d]", i))
		interfaces[i] = cp.GetClass(interfaceIndex)
	}

	return interfaces
}
