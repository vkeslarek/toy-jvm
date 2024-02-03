package instructionset

import (
	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
	"github.com/vkeslarek/toy-jvm/types"
)

type ArithmeticOpcodeParser struct{}

type OpcodeType string

const (
	AddArithmeticOpcodeType OpcodeType = "ADD"
	DivArithmeticOpcodeType OpcodeType = "DIV"
	IncArithmeticOpcodeType OpcodeType = "INC"
	MulArithmeticOpcodeType OpcodeType = "MUL"
	SubArithmeticOpcodeType OpcodeType = "SUB"
	NegArithmeticOpcodeType OpcodeType = "NEG"
	RemArithmeticOpcodeType OpcodeType = "REM"
)

type ArithmeticOpcode struct {
	OpcodeValue uint8
	OpcodeName  OpcodeName
	Type        OpcodeType
	DataType    types.DataType
}

func (a *ArithmeticOpcode) Value() uint8 {
	return a.OpcodeValue
}

func (a *ArithmeticOpcode) Name() OpcodeName {
	return a.OpcodeName
}

func (a *ArithmeticOpcode) Operands() []uint8 {
	return []uint8{}
}

func (parser *ArithmeticOpcodeParser) Parse(fieldName string, reader *reader.CodeReader, opcodeValue uint8, opcodeName OpcodeName, cp *constantpool.ConstantPool) Opcode {
	switch opcodeName {
	//ADD OPCODES
	case daddOpcodeName:
		return &ArithmeticOpcode{OpcodeValue: opcodeValue, OpcodeName: opcodeName, Type: AddArithmeticOpcodeType, DataType: types.DoubleDataType}
	case faddOpcodeName:
		return &ArithmeticOpcode{OpcodeValue: opcodeValue, OpcodeName: opcodeName, Type: AddArithmeticOpcodeType, DataType: types.FloatDataType}
	case iaddOpcodeName:
		return &ArithmeticOpcode{OpcodeValue: opcodeValue, OpcodeName: opcodeName, Type: AddArithmeticOpcodeType, DataType: types.IntegerDataType}
	case laddOpcodeName:
		return &ArithmeticOpcode{OpcodeValue: opcodeValue, OpcodeName: opcodeName, Type: AddArithmeticOpcodeType, DataType: types.LongDataType}

	// DIV OPCODES
	case ddivOpcodeName:
		return &ArithmeticOpcode{OpcodeValue: opcodeValue, OpcodeName: opcodeName, Type: DivArithmeticOpcodeType, DataType: types.DoubleDataType}
	case fdivOpcodeName:
		return &ArithmeticOpcode{OpcodeValue: opcodeValue, OpcodeName: opcodeName, Type: DivArithmeticOpcodeType, DataType: types.FloatDataType}
	case idivOpcodeName:
		return &ArithmeticOpcode{OpcodeValue: opcodeValue, OpcodeName: opcodeName, Type: DivArithmeticOpcodeType, DataType: types.IntegerDataType}
	case ldivOpcodeName:
		return &ArithmeticOpcode{OpcodeValue: opcodeValue, OpcodeName: opcodeName, Type: DivArithmeticOpcodeType, DataType: types.LongDataType}

	// INC OPCODES
	case iincOpcodeName:
		return &ArithmeticOpcode{OpcodeValue: opcodeValue, OpcodeName: opcodeName, Type: IncArithmeticOpcodeType, DataType: types.IntegerDataType}

	// MUL OPCODES
	case dmulOpcodeName:
		return &ArithmeticOpcode{OpcodeValue: opcodeValue, OpcodeName: opcodeName, Type: MulArithmeticOpcodeType, DataType: types.DoubleDataType}
	case fmulOpcodeName:
		return &ArithmeticOpcode{OpcodeValue: opcodeValue, OpcodeName: opcodeName, Type: MulArithmeticOpcodeType, DataType: types.FloatDataType}
	case imulOpcodeName:
		return &ArithmeticOpcode{OpcodeValue: opcodeValue, OpcodeName: opcodeName, Type: MulArithmeticOpcodeType, DataType: types.IntegerDataType}
	case lmulOpcodeName:
		return &ArithmeticOpcode{OpcodeValue: opcodeValue, OpcodeName: opcodeName, Type: MulArithmeticOpcodeType, DataType: types.LongDataType}

	// SUB OPCODES
	case dsubOpcodeName:
		return &ArithmeticOpcode{OpcodeValue: opcodeValue, OpcodeName: opcodeName, Type: SubArithmeticOpcodeType, DataType: types.DoubleDataType}
	case fsubOpcodeName:
		return &ArithmeticOpcode{OpcodeValue: opcodeValue, OpcodeName: opcodeName, Type: SubArithmeticOpcodeType, DataType: types.FloatDataType}
	case isubOpcodeName:
		return &ArithmeticOpcode{OpcodeValue: opcodeValue, OpcodeName: opcodeName, Type: SubArithmeticOpcodeType, DataType: types.IntegerDataType}
	case lsubOpcodeName:
		return &ArithmeticOpcode{OpcodeValue: opcodeValue, OpcodeName: opcodeName, Type: SubArithmeticOpcodeType, DataType: types.LongDataType}

	// NEG OPCODES
	case dnegOpcodeName:
		return &ArithmeticOpcode{OpcodeValue: opcodeValue, OpcodeName: opcodeName, Type: NegArithmeticOpcodeType, DataType: types.DoubleDataType}
	case fnegOpcodeName:
		return &ArithmeticOpcode{OpcodeValue: opcodeValue, OpcodeName: opcodeName, Type: NegArithmeticOpcodeType, DataType: types.FloatDataType}
	case inegOpcodeName:
		return &ArithmeticOpcode{OpcodeValue: opcodeValue, OpcodeName: opcodeName, Type: NegArithmeticOpcodeType, DataType: types.IntegerDataType}
	case lnegOpcodeName:
		return &ArithmeticOpcode{OpcodeValue: opcodeValue, OpcodeName: opcodeName, Type: NegArithmeticOpcodeType, DataType: types.LongDataType}

	// REM OPCODES
	case dremOpcodeName:
		return &ArithmeticOpcode{OpcodeValue: opcodeValue, OpcodeName: opcodeName, Type: RemArithmeticOpcodeType, DataType: types.DoubleDataType}
	case fremOpcodeName:
		return &ArithmeticOpcode{OpcodeValue: opcodeValue, OpcodeName: opcodeName, Type: RemArithmeticOpcodeType, DataType: types.FloatDataType}
	case iremOpcodeName:
		return &ArithmeticOpcode{OpcodeValue: opcodeValue, OpcodeName: opcodeName, Type: RemArithmeticOpcodeType, DataType: types.IntegerDataType}
	case lremOpcodeName:
		return &ArithmeticOpcode{OpcodeValue: opcodeValue, OpcodeName: opcodeName, Type: RemArithmeticOpcodeType, DataType: types.LongDataType}
	default:
		return nil
	}
}
