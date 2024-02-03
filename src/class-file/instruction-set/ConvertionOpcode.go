package instructionset

import (
	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type ConvertionOpcode struct{}

func (*ConvertionOpcode) Parse(fieldName string, reader *reader.CodeReader, opcodeValue uint8, opcodeName OpcodeName, cp *constantpool.ConstantPool) *Opcode {
	// TODO: Paser opcodes:
	/*
		d2f
		d2i
		d2l
		f2d
		f2i
		f2l
		i2b
		i2c
		i2d
		i2f
		i2l
		i2s
		l2d
		l2f
		l2i
	*/
	return nil
}
