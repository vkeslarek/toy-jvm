package instructionset

import (
	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type ArrayLengthOpcodeParser struct{}

type ArrayLenthOpcode struct {
	OpcodeValue uint8
	OpcodeName  OpcodeName
}

func (a *ArrayLenthOpcode) Value() uint8 {
	return a.OpcodeValue
}

func (a *ArrayLenthOpcode) Name() OpcodeName {
	return a.OpcodeName
}

func (a *ArrayLenthOpcode) Operands() []uint8 {
	return []uint8{}
}

func (p *ArrayLengthOpcodeParser) Parse(fieldName string, reader *reader.CodeReader, opcodeValue uint8, opcodeName OpcodeName, cp *constantpool.ConstantPool) Opcode {
	return &ArrayLenthOpcode{OpcodeValue: opcodeValue, OpcodeName: opcodeName}
}
