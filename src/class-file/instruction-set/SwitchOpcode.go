package instructionset

import (
	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type SwithOpcodeParser struct{}

func (p SwithOpcodeParser) Parse(fieldName string, reader *reader.CodeReader, opcodeValue uint8, opcodeName OpcodeName, cp *constantpool.ConstantPool) *Opcode {

	// TODO: Paser opcodes:
	/*
		tableswitch
		lookupswitch
	*/
	return nil
}
