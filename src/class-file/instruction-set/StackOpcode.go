package instructionset

import (
	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type PushPopOpcodeParser struct{}

func (p *PushPopOpcodeParser) Parse(fieldName string, reader *reader.CodeReader, opcodeValue uint8, opcodeName OpcodeName, cp *constantpool.ConstantPool) *Opcode {
	// TODO: Paser opcodes:
	/*
		PUSH:
			bipush
			sipush
		POP:
			pop
			pop2
		DUP:
			dup
			dup_x1
			dup_x2
			dup2
			dup2_x1
			dup2_x2
		CONST:
			dconst_0
			dconst_1
			fconst_0
			fconst_1
			fconst_2
			iconst_m1
			iconst_0
			iconst_1
			iconst_2
			iconst_3
			iconst_4
			iconst_5
			lconst_0
			lconst_1
		SWAP:
			swap
	*/
	return nil
}
