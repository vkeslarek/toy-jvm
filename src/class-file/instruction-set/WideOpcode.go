package instructionset

import (
	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type WideOpcodeParser struct{}

func (w *WideOpcodeParser) Parse(fieldName string, reader *reader.CodeReader, opcodeValue uint8, opcodeName OpcodeName, cp *constantpool.ConstantPool) *Opcode {
	// TODO: Paser opcodes:
	/*
		wide
	*/
	return nil
}
