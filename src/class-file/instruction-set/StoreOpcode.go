package instructionset

import (
	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type StoreOpcodeParser struct{}

func (*StoreOpcodeParser) Parse(fieldName string, reader *reader.CodeReader, opcodeValue uint8, opcodeName OpcodeName, cp *constantpool.ConstantPool) *Opcode {
	// TODO: Paser opcodes:
	/*
		astore
		astore_0
		astore_1
		astore_2
		astore_3
		bastore
		castore
		dastore
		dstore
		dstore_0
		dstore_1
		dstore_2
		dstore_3
		fastore
		fstore
		fstore_0
		fstore_1
		fstore_2
		fstore_3
		iastore
		istore
		istore_0
		istore_1
		istore_2
		istore_3
		lastore
		lstore
		lstore_0
		lstore_1
		lstore_2
		lstore_3
		sastore
	*/
	return nil
}
