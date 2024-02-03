package instructionset

import (
	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type InvokeOpcodeParser struct{}

func (p *InvokeOpcodeParser) Parse(fieldName string, reader *reader.CodeReader, opcodeValue uint8, opcodeName OpcodeName, cp *constantpool.ConstantPool) *Opcode {
	// TODO: Paser opcodes:
	/*
		invokedynamic
		invokeinterface
		invokespecial
		invokestatic
		invokevirtual
	*/
	return nil
}
