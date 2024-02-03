package instructionset

import (
	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type LoadOpcodeParser struct{}

func (*LoadOpcodeParser) Parse(fieldName string, reader *reader.CodeReader, opcodeValue uint8, opcodeName OpcodeName, cp *constantpool.ConstantPool) *Opcode {
	// TODO: Paser opcodes:
	/*
		aaload
		aload
		aload_0
		aload_1
		aload_2
		aload_3
		baload
		saload
		lload
		lload_0
		lload_1
		lload_2
		lload_3
		caload
		daload
		dload
		dload_0
		dload_1
		dload_2
		dload_3
		faload
		fload
		fload_0
		fload_1
		fload_2
		fload_3
		iaload
		iload
		iload_0
		iload_1
		iload_2
		iload_3
		laload
	*/
	return nil
}
