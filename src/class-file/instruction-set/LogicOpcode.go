package instructionset

import (
	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type LogicOpcodeParser struct{}

func (LogicOpcodeParser) Parse(fieldName string, reader *reader.CodeReader, opcodeValue uint8, opcodeName OpcodeName, cp *constantpool.ConstantPool) *Opcode {
	// TODO: Paser opcodes:
	/*
		AND:
			iand
			land
		OR:
			ior
			lor
		XOR:
			ixor
			lxor
		SHL:
			ishl
			lshl
		SHR:
			iushr
			lushr
			ishr
			lshr
	*/
	return nil
}
