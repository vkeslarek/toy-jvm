package instructionset

import (
	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type ConditionalBranchOpcodeParser struct{}

func (ConditionalBranchOpcodeParser) Parse(fieldName string, reader *reader.CodeReader, opcodeValue uint8, opcodeName OpcodeName, cp *constantpool.ConstantPool) *Opcode {
	// TODO: Paser opcodes:
	/*
		if_acmpeq
		if_acmpne
		if_icmpeq
		if_icmpne
		if_icmplt
		if_icmpge
		if_icmpgt
		if_icmple
		ifeq
		ifne
		iflt
		ifge
		ifgt
		ifle
		ifnonnull
		ifnull
	*/
	return nil
}
