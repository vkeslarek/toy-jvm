package instructionset

type OperandName string

const (
	OperandNameNone            = OperandName("")
	OperandNameArrayRef        = OperandName("arrayref")
	OperandNameIndex           = OperandName("index")
	OperandNameObjectReference = OperandName("ref")
)

type Operand interface {
	Name() OperandName
}

type OperandArrayRef uint16

func (o OperandArrayRef) Name() OperandName {
	return OperandNameArrayRef
}

type OperandIndex uint16

func (o OperandIndex) Name() OperandName {
	return OperandNameIndex
}
