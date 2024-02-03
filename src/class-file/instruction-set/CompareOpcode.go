package instructionset

import (
	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
	"github.com/vkeslarek/toy-jvm/types"
)

type CompareOpcodeParser struct{}

type CompareOperation string

const (
	NoopCompareOperation    = CompareOperation("no-op")
	GreaterCompareOperation = CompareOperation("greater")
	LessCompareOperation    = CompareOperation("less")
)

type CompareOpcode struct {
	OpcodeValue uint8
	OpcodeName  OpcodeName
	Operation   CompareOperation
	DataType    types.DataType
}

func (c *CompareOpcode) Value() uint8 {
	return c.OpcodeValue
}

func (c *CompareOpcode) Name() OpcodeName {
	return c.OpcodeName
}

func (c *CompareOpcode) Operands() []uint8 {
	return []uint8{}
}

func (p *CompareOpcodeParser) Parse(fieldName string, reader *reader.CodeReader, opcodeValue uint8, opcodeName OpcodeName, cp *constantpool.ConstantPool) Opcode {
	switch opcodeName {
	case dcmpgOpcodeName:
		return &CompareOpcode{OpcodeValue: opcodeValue, OpcodeName: opcodeName, Operation: GreaterCompareOperation, DataType: types.DoubleDataType}
	case dcmplOpcodeName:
		return &CompareOpcode{OpcodeValue: opcodeValue, OpcodeName: opcodeName, Operation: LessCompareOperation, DataType: types.DoubleDataType}
	case fcmpgOpcodeName:
		return &CompareOpcode{OpcodeValue: opcodeValue, OpcodeName: opcodeName, Operation: GreaterCompareOperation, DataType: types.FloatDataType}
	case fcmplOpcodeName:
		return &CompareOpcode{OpcodeValue: opcodeValue, OpcodeName: opcodeName, Operation: LessCompareOperation, DataType: types.FloatDataType}
	case lcmpOpcodeName:
		return &CompareOpcode{OpcodeValue: opcodeValue, OpcodeName: opcodeName, Operation: NoopCompareOperation, DataType: types.LongDataType}
	default:
		return nil
	}
}
