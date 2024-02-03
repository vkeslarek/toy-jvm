package instructionset

import (
	constantpool "github.com/vkeslarek/toy-jvm/class-file/constant-pool"
	"github.com/vkeslarek/toy-jvm/class-file/reader"
)

type OpcodeName string

const (
	aaloadOpcodeName          OpcodeName = "aaload"
	aastoreOpcodeName         OpcodeName = "aastore"
	aconst_nullOpcodeName     OpcodeName = "aconst_null"
	aloadOpcodeName           OpcodeName = "aload"
	aload_0OpcodeName         OpcodeName = "aload_0"
	aload_1OpcodeName         OpcodeName = "aload_1"
	aload_2OpcodeName         OpcodeName = "aload_2"
	aload_3OpcodeName         OpcodeName = "aload_3"
	anewarrayOpcodeName       OpcodeName = "anewarray"
	areturnOpcodeName         OpcodeName = "areturn"
	arraylengthOpcodeName     OpcodeName = "arraylength"
	astoreOpcodeName          OpcodeName = "astore"
	astore_0OpcodeName        OpcodeName = "astore_0"
	astore_1OpcodeName        OpcodeName = "astore_1"
	astore_2OpcodeName        OpcodeName = "astore_2"
	astore_3OpcodeName        OpcodeName = "astore_3"
	athrowOpcodeName          OpcodeName = "athrow"
	baloadOpcodeName          OpcodeName = "baload"
	bastoreOpcodeName         OpcodeName = "bastore"
	bipushOpcodeName          OpcodeName = "bipush"
	caloadOpcodeName          OpcodeName = "caload"
	castoreOpcodeName         OpcodeName = "castore"
	checkcastOpcodeName       OpcodeName = "checkcast"
	d2fOpcodeName             OpcodeName = "d2f"
	d2iOpcodeName             OpcodeName = "d2i"
	d2lOpcodeName             OpcodeName = "d2l"
	daddOpcodeName            OpcodeName = "dadd"
	daloadOpcodeName          OpcodeName = "daload"
	dastoreOpcodeName         OpcodeName = "dastore"
	dcmpgOpcodeName           OpcodeName = "dcmpg"
	dcmplOpcodeName           OpcodeName = "dcmpl"
	dconst_0OpcodeName        OpcodeName = "dconst_0"
	dconst_1OpcodeName        OpcodeName = "dconst_1"
	ddivOpcodeName            OpcodeName = "ddiv"
	dloadOpcodeName           OpcodeName = "dload"
	dload_0OpcodeName         OpcodeName = "dload_0"
	dload_1OpcodeName         OpcodeName = "dload_1"
	dload_2OpcodeName         OpcodeName = "dload_2"
	dload_3OpcodeName         OpcodeName = "dload_3"
	dmulOpcodeName            OpcodeName = "dmul"
	dnegOpcodeName            OpcodeName = "dneg"
	dremOpcodeName            OpcodeName = "drem"
	dreturnOpcodeName         OpcodeName = "dreturn"
	dstoreOpcodeName          OpcodeName = "dstore"
	dstore_0OpcodeName        OpcodeName = "dstore_0"
	dstore_1OpcodeName        OpcodeName = "dstore_1"
	dstore_2OpcodeName        OpcodeName = "dstore_2"
	dstore_3OpcodeName        OpcodeName = "dstore_3"
	dsubOpcodeName            OpcodeName = "dsub"
	dupOpcodeName             OpcodeName = "dup"
	dup_x1OpcodeName          OpcodeName = "dup_x1"
	dup_x2OpcodeName          OpcodeName = "dup_x2"
	dup2OpcodeName            OpcodeName = "dup2"
	dup2_x1OpcodeName         OpcodeName = "dup2_x1"
	dup2_x2OpcodeName         OpcodeName = "dup2_x2"
	f2dOpcodeName             OpcodeName = "f2d"
	f2iOpcodeName             OpcodeName = "f2i"
	f2lOpcodeName             OpcodeName = "f2l"
	faddOpcodeName            OpcodeName = "fadd"
	faloadOpcodeName          OpcodeName = "faload"
	fastoreOpcodeName         OpcodeName = "fastore"
	fcmpgOpcodeName           OpcodeName = "fcmpg"
	fcmplOpcodeName           OpcodeName = "fcmpl"
	fconst_0OpcodeName        OpcodeName = "fconst_0"
	fconst_1OpcodeName        OpcodeName = "fconst_1"
	fconst_2OpcodeName        OpcodeName = "fconst_2"
	fdivOpcodeName            OpcodeName = "fdiv"
	floadOpcodeName           OpcodeName = "fload"
	fload_0OpcodeName         OpcodeName = "fload_0"
	fload_1OpcodeName         OpcodeName = "fload_1"
	fload_2OpcodeName         OpcodeName = "fload_2"
	fload_3OpcodeName         OpcodeName = "fload_3"
	fmulOpcodeName            OpcodeName = "fmul"
	fnegOpcodeName            OpcodeName = "fneg"
	fremOpcodeName            OpcodeName = "frem"
	freturnOpcodeName         OpcodeName = "freturn"
	fstoreOpcodeName          OpcodeName = "fstore"
	fstore_0OpcodeName        OpcodeName = "fstore_0"
	fstore_1OpcodeName        OpcodeName = "fstore_1"
	fstore_2OpcodeName        OpcodeName = "fstore_2"
	fstore_3OpcodeName        OpcodeName = "fstore_3"
	fsubOpcodeName            OpcodeName = "fsub"
	getfieldOpcodeName        OpcodeName = "getfield"
	getstaticOpcodeName       OpcodeName = "getstatic"
	gotoOpcodeName            OpcodeName = "goto"
	goto_wOpcodeName          OpcodeName = "goto_w"
	i2bOpcodeName             OpcodeName = "i2b"
	i2cOpcodeName             OpcodeName = "i2c"
	i2dOpcodeName             OpcodeName = "i2d"
	i2fOpcodeName             OpcodeName = "i2f"
	i2lOpcodeName             OpcodeName = "i2l"
	i2sOpcodeName             OpcodeName = "i2s"
	iaddOpcodeName            OpcodeName = "iadd"
	ialoadOpcodeName          OpcodeName = "iaload"
	iandOpcodeName            OpcodeName = "iand"
	iastoreOpcodeName         OpcodeName = "iastore"
	iconst_m1OpcodeName       OpcodeName = "iconst_m1"
	iconst_0OpcodeName        OpcodeName = "iconst_0"
	iconst_1OpcodeName        OpcodeName = "iconst_1"
	iconst_2OpcodeName        OpcodeName = "iconst_2"
	iconst_3OpcodeName        OpcodeName = "iconst_3"
	iconst_4OpcodeName        OpcodeName = "iconst_4"
	iconst_5OpcodeName        OpcodeName = "iconst_5"
	idivOpcodeName            OpcodeName = "idiv"
	if_acmpeqOpcodeName       OpcodeName = "if_acmpeq"
	if_acmpneOpcodeName       OpcodeName = "if_acmpne"
	if_icmpeqOpcodeName       OpcodeName = "if_icmpeq"
	if_icmpneOpcodeName       OpcodeName = "if_icmpne"
	if_icmpltOpcodeName       OpcodeName = "if_icmplt"
	if_icmpgeOpcodeName       OpcodeName = "if_icmpge"
	if_icmpgtOpcodeName       OpcodeName = "if_icmpgt"
	if_icmpleOpcodeName       OpcodeName = "if_icmple"
	ifeqOpcodeName            OpcodeName = "ifeq"
	ifneOpcodeName            OpcodeName = "ifne"
	ifltOpcodeName            OpcodeName = "iflt"
	ifgeOpcodeName            OpcodeName = "ifge"
	ifgtOpcodeName            OpcodeName = "ifgt"
	ifleOpcodeName            OpcodeName = "ifle"
	ifnonnullOpcodeName       OpcodeName = "ifnonnull"
	ifnullOpcodeName          OpcodeName = "ifnull"
	iincOpcodeName            OpcodeName = "iinc"
	iloadOpcodeName           OpcodeName = "iload"
	iload_0OpcodeName         OpcodeName = "iload_0"
	iload_1OpcodeName         OpcodeName = "iload_1"
	iload_2OpcodeName         OpcodeName = "iload_2"
	iload_3OpcodeName         OpcodeName = "iload_3"
	imulOpcodeName            OpcodeName = "imul"
	inegOpcodeName            OpcodeName = "ineg"
	instanceofOpcodeName      OpcodeName = "instanceof"
	invokedynamicOpcodeName   OpcodeName = "invokedynamic"
	invokeinterfaceOpcodeName OpcodeName = "invokeinterface"
	invokespecialOpcodeName   OpcodeName = "invokespecial"
	invokestaticOpcodeName    OpcodeName = "invokestatic"
	invokevirtualOpcodeName   OpcodeName = "invokevirtual"
	iorOpcodeName             OpcodeName = "ior"
	iremOpcodeName            OpcodeName = "irem"
	ireturnOpcodeName         OpcodeName = "ireturn"
	ishlOpcodeName            OpcodeName = "ishl"
	ishrOpcodeName            OpcodeName = "ishr"
	istoreOpcodeName          OpcodeName = "istore"
	istore_0OpcodeName        OpcodeName = "istore_0"
	istore_1OpcodeName        OpcodeName = "istore_1"
	istore_2OpcodeName        OpcodeName = "istore_2"
	istore_3OpcodeName        OpcodeName = "istore_3"
	isubOpcodeName            OpcodeName = "isub"
	iushrOpcodeName           OpcodeName = "iushr"
	ixorOpcodeName            OpcodeName = "ixor"
	jsrOpcodeName             OpcodeName = "jsr"
	jsr_wOpcodeName           OpcodeName = "jsr_w"
	l2dOpcodeName             OpcodeName = "l2d"
	l2fOpcodeName             OpcodeName = "l2f"
	l2iOpcodeName             OpcodeName = "l2i"
	laddOpcodeName            OpcodeName = "ladd"
	laloadOpcodeName          OpcodeName = "laload"
	landOpcodeName            OpcodeName = "land"
	lastoreOpcodeName         OpcodeName = "lastore"
	lcmpOpcodeName            OpcodeName = "lcmp"
	lconst_0OpcodeName        OpcodeName = "lconst_0"
	lconst_1OpcodeName        OpcodeName = "lconst_1"
	ldcOpcodeName             OpcodeName = "ldc"
	ldc_wOpcodeName           OpcodeName = "ldc_w"
	ldc2_wOpcodeName          OpcodeName = "ldc2_w"
	ldivOpcodeName            OpcodeName = "ldiv"
	lloadOpcodeName           OpcodeName = "lload"
	lload_0OpcodeName         OpcodeName = "lload_0"
	lload_1OpcodeName         OpcodeName = "lload_1"
	lload_2OpcodeName         OpcodeName = "lload_2"
	lload_3OpcodeName         OpcodeName = "lload_3"
	lmulOpcodeName            OpcodeName = "lmul"
	lnegOpcodeName            OpcodeName = "lneg"
	lookupswitchOpcodeName    OpcodeName = "lookupswitch"
	lorOpcodeName             OpcodeName = "lor"
	lremOpcodeName            OpcodeName = "lrem"
	lreturnOpcodeName         OpcodeName = "lreturn"
	lshlOpcodeName            OpcodeName = "lshl"
	lshrOpcodeName            OpcodeName = "lshr"
	lstoreOpcodeName          OpcodeName = "lstore"
	lstore_0OpcodeName        OpcodeName = "lstore_0"
	lstore_1OpcodeName        OpcodeName = "lstore_1"
	lstore_2OpcodeName        OpcodeName = "lstore_2"
	lstore_3OpcodeName        OpcodeName = "lstore_3"
	lsubOpcodeName            OpcodeName = "lsub"
	lushrOpcodeName           OpcodeName = "lushr"
	lxorOpcodeName            OpcodeName = "lxor"
	monitorenterOpcodeName    OpcodeName = "monitorenter"
	monitorexitOpcodeName     OpcodeName = "monitorexit"
	multianewarrayOpcodeName  OpcodeName = "multianewarray"
	newOpcodeName             OpcodeName = "new"
	newarrayOpcodeName        OpcodeName = "newarray"
	nopOpcodeName             OpcodeName = "nop"
	popOpcodeName             OpcodeName = "pop"
	pop2OpcodeName            OpcodeName = "pop2"
	putfieldOpcodeName        OpcodeName = "putfield"
	putstaticOpcodeName       OpcodeName = "putstatic"
	retOpcodeName             OpcodeName = "ret"
	returnOpcodeName          OpcodeName = "return"
	saloadOpcodeName          OpcodeName = "saload"
	sastoreOpcodeName         OpcodeName = "sastore"
	sipushOpcodeName          OpcodeName = "sipush"
	swapOpcodeName            OpcodeName = "swap"
	tableswitchOpcodeName     OpcodeName = "tableswitch"
	wideOpcodeName            OpcodeName = "wide"
)

type OpcodeTable map[uint8]OpcodeName

var opcodeTable = OpcodeTable{
	0x32: aaloadOpcodeName,
	0x53: aastoreOpcodeName,
	0x01: aconst_nullOpcodeName,
	0x19: aloadOpcodeName,
	0x2A: aload_0OpcodeName,
	0x2B: aload_1OpcodeName,
	0x2C: aload_2OpcodeName,
	0x2D: aload_3OpcodeName,
	0xBD: anewarrayOpcodeName,
	0xB0: areturnOpcodeName,
	0xBE: arraylengthOpcodeName,
	0x3A: astoreOpcodeName,
	0x4B: astore_0OpcodeName,
	0x4C: astore_1OpcodeName,
	0x4D: astore_2OpcodeName,
	0x4E: astore_3OpcodeName,
	0xBF: athrowOpcodeName,
	0x33: baloadOpcodeName,
	0x54: bastoreOpcodeName,
	0x10: bipushOpcodeName,
	0x34: caloadOpcodeName,
	0x55: castoreOpcodeName,
	0xC0: checkcastOpcodeName,
	0x90: d2fOpcodeName,
	0x8E: d2iOpcodeName,
	0x8F: d2lOpcodeName,
	0x63: daddOpcodeName,
	0x31: daloadOpcodeName,
	0x52: dastoreOpcodeName,
	0x98: dcmpgOpcodeName,
	0x97: dcmplOpcodeName,
	0x0E: dconst_0OpcodeName,
	0x0F: dconst_1OpcodeName,
	0x6F: ddivOpcodeName,
	0x18: dloadOpcodeName,
	0x26: dload_0OpcodeName,
	0x27: dload_1OpcodeName,
	0x28: dload_2OpcodeName,
	0x29: dload_3OpcodeName,
	0x6B: dmulOpcodeName,
	0x77: dnegOpcodeName,
	0x73: dremOpcodeName,
	0xAF: dreturnOpcodeName,
	0x39: dstoreOpcodeName,
	0x47: dstore_0OpcodeName,
	0x48: dstore_1OpcodeName,
	0x49: dstore_2OpcodeName,
	0x4A: dstore_3OpcodeName,
	0x67: dsubOpcodeName,
	0x59: dupOpcodeName,
	0x5A: dup_x1OpcodeName,
	0x5B: dup_x2OpcodeName,
	0x5C: dup2OpcodeName,
	0x5D: dup2_x1OpcodeName,
	0x5E: dup2_x2OpcodeName,
	0x8D: f2dOpcodeName,
	0x8B: f2iOpcodeName,
	0x8C: f2lOpcodeName,
	0x62: faddOpcodeName,
	0x30: faloadOpcodeName,
	0x51: fastoreOpcodeName,
	0x96: fcmpgOpcodeName,
	0x95: fcmplOpcodeName,
	0x0B: fconst_0OpcodeName,
	0x0C: fconst_1OpcodeName,
	0x0D: fconst_2OpcodeName,
	0x6E: fdivOpcodeName,
	0x17: floadOpcodeName,
	0x22: fload_0OpcodeName,
	0x23: fload_1OpcodeName,
	0x24: fload_2OpcodeName,
	0x25: fload_3OpcodeName,
	0x6A: fmulOpcodeName,
	0x76: fnegOpcodeName,
	0x72: fremOpcodeName,
	0xAE: freturnOpcodeName,
	0x38: fstoreOpcodeName,
	0x43: fstore_0OpcodeName,
	0x44: fstore_1OpcodeName,
	0x45: fstore_2OpcodeName,
	0x46: fstore_3OpcodeName,
	0x66: fsubOpcodeName,
	0xB4: getfieldOpcodeName,
	0xB2: getstaticOpcodeName,
	0xA7: gotoOpcodeName,
	0xC8: goto_wOpcodeName,
	0x91: i2bOpcodeName,
	0x92: i2cOpcodeName,
	0x87: i2dOpcodeName,
	0x86: i2fOpcodeName,
	0x85: i2lOpcodeName,
	0x93: i2sOpcodeName,
	0x60: iaddOpcodeName,
	0x2E: ialoadOpcodeName,
	0x7E: iandOpcodeName,
	0x4F: iastoreOpcodeName,
	0x02: iconst_m1OpcodeName,
	0x03: iconst_0OpcodeName,
	0x04: iconst_1OpcodeName,
	0x05: iconst_2OpcodeName,
	0x06: iconst_3OpcodeName,
	0x07: iconst_4OpcodeName,
	0x08: iconst_5OpcodeName,
	0x6C: idivOpcodeName,
	0xA5: if_acmpeqOpcodeName,
	0xA6: if_acmpneOpcodeName,
	0x9F: if_icmpeqOpcodeName,
	0xA0: if_icmpneOpcodeName,
	0xA1: if_icmpltOpcodeName,
	0xA2: if_icmpgeOpcodeName,
	0xA3: if_icmpgtOpcodeName,
	0xA4: if_icmpleOpcodeName,
	0x99: ifeqOpcodeName,
	0x9A: ifneOpcodeName,
	0x9B: ifltOpcodeName,
	0x9C: ifgeOpcodeName,
	0x9D: ifgtOpcodeName,
	0x9E: ifleOpcodeName,
	0xC7: ifnonnullOpcodeName,
	0xC6: ifnullOpcodeName,
	0x84: iincOpcodeName,
	0x15: iloadOpcodeName,
	0x1A: iload_0OpcodeName,
	0x1B: iload_1OpcodeName,
	0x1C: iload_2OpcodeName,
	0x1D: iload_3OpcodeName,
	0x68: imulOpcodeName,
	0x74: inegOpcodeName,
	0xC1: instanceofOpcodeName,
	0xBA: invokedynamicOpcodeName,
	0xB9: invokeinterfaceOpcodeName,
	0xB7: invokespecialOpcodeName,
	0xB8: invokestaticOpcodeName,
	0xB6: invokevirtualOpcodeName,
	0x80: iorOpcodeName,
	0x70: iremOpcodeName,
	0xAC: ireturnOpcodeName,
	0x78: ishlOpcodeName,
	0x7A: ishrOpcodeName,
	0x36: istoreOpcodeName,
	0x3B: istore_0OpcodeName,
	0x3C: istore_1OpcodeName,
	0x3D: istore_2OpcodeName,
	0x3E: istore_3OpcodeName,
	0x64: isubOpcodeName,
	0x7C: iushrOpcodeName,
	0x82: ixorOpcodeName,
	0xA8: jsrOpcodeName,
	0xC9: jsr_wOpcodeName,
	0x8A: l2dOpcodeName,
	0x89: l2fOpcodeName,
	0x88: l2iOpcodeName,
	0x61: laddOpcodeName,
	0x2F: laloadOpcodeName,
	0x7F: landOpcodeName,
	0x50: lastoreOpcodeName,
	0x94: lcmpOpcodeName,
	0x09: lconst_0OpcodeName,
	0x0A: lconst_1OpcodeName,
	0x12: ldcOpcodeName,
	0x13: ldc_wOpcodeName,
	0x14: ldc2_wOpcodeName,
	0x6D: ldivOpcodeName,
	0x16: lloadOpcodeName,
	0x1E: lload_0OpcodeName,
	0x1F: lload_1OpcodeName,
	0x20: lload_2OpcodeName,
	0x21: lload_3OpcodeName,
	0x69: lmulOpcodeName,
	0x75: lnegOpcodeName,
	0xAB: lookupswitchOpcodeName,
	0x81: lorOpcodeName,
	0x71: lremOpcodeName,
	0xAD: lreturnOpcodeName,
	0x79: lshlOpcodeName,
	0x7B: lshrOpcodeName,
	0x37: lstoreOpcodeName,
	0x3F: lstore_0OpcodeName,
	0x40: lstore_1OpcodeName,
	0x41: lstore_2OpcodeName,
	0x42: lstore_3OpcodeName,
	0x65: lsubOpcodeName,
	0x7D: lushrOpcodeName,
	0x83: lxorOpcodeName,
	0xC2: monitorenterOpcodeName,
	0xC3: monitorexitOpcodeName,
	0xC5: multianewarrayOpcodeName,
	0xBB: newOpcodeName,
	0xBC: newarrayOpcodeName,
	0x00: nopOpcodeName,
	0x57: popOpcodeName,
	0x58: pop2OpcodeName,
	0xB5: putfieldOpcodeName,
	0xB3: putstaticOpcodeName,
	0xA9: retOpcodeName,
	0xB1: returnOpcodeName,
	0x35: saloadOpcodeName,
	0x56: sastoreOpcodeName,
	0x11: sipushOpcodeName,
	0x5F: swapOpcodeName,
	0xAA: tableswitchOpcodeName,
	0xC4: wideOpcodeName,
}

type OpcodeParser interface {
	Parse(fieldName string, reader *reader.CodeReader, opcodeValue uint8, opcodeName OpcodeName, cp *constantpool.ConstantPool) Opcode
}

type OpcodeParserTable (map[OpcodeName]OpcodeParser)

var arithmeticOpcodeParser = ArithmeticOpcodeParser{}
var arrayLengthOpcodeParser = ArrayLengthOpcodeParser{}
var compareOpcodeParser = CompareOpcodeParser{}

var opcodeParserTable = OpcodeParserTable{
	// ARITHMETIC OPCODES
	daddOpcodeName: &arithmeticOpcodeParser,
	faddOpcodeName: &arithmeticOpcodeParser,
	iaddOpcodeName: &arithmeticOpcodeParser,
	laddOpcodeName: &arithmeticOpcodeParser,
	ddivOpcodeName: &arithmeticOpcodeParser,
	fdivOpcodeName: &arithmeticOpcodeParser,
	idivOpcodeName: &arithmeticOpcodeParser,
	ldivOpcodeName: &arithmeticOpcodeParser,
	iincOpcodeName: &arithmeticOpcodeParser,
	dmulOpcodeName: &arithmeticOpcodeParser,
	fmulOpcodeName: &arithmeticOpcodeParser,
	imulOpcodeName: &arithmeticOpcodeParser,
	lmulOpcodeName: &arithmeticOpcodeParser,
	dsubOpcodeName: &arithmeticOpcodeParser,
	fsubOpcodeName: &arithmeticOpcodeParser,
	isubOpcodeName: &arithmeticOpcodeParser,
	lsubOpcodeName: &arithmeticOpcodeParser,
	dnegOpcodeName: &arithmeticOpcodeParser,
	fnegOpcodeName: &arithmeticOpcodeParser,
	inegOpcodeName: &arithmeticOpcodeParser,
	lnegOpcodeName: &arithmeticOpcodeParser,
	dremOpcodeName: &arithmeticOpcodeParser,
	fremOpcodeName: &arithmeticOpcodeParser,
	iremOpcodeName: &arithmeticOpcodeParser,
	lremOpcodeName: &arithmeticOpcodeParser,
	// ARRAY LENGTH
	arraylengthOpcodeName: &arrayLengthOpcodeParser,
	// COMPARE OPCODES
	dcmplOpcodeName: &compareOpcodeParser,
	dcmpgOpcodeName: &compareOpcodeParser,
	fcmplOpcodeName: &compareOpcodeParser,
	fcmpgOpcodeName: &compareOpcodeParser,
	lcmpOpcodeName:  &compareOpcodeParser,
	// TODO: other OPCODES
}

type Opcode interface {
	Value() uint8
	Name() OpcodeName
	Operands() []uint8
}

func ParseOpcode(fieldName string, reader *reader.CodeReader, cp *constantpool.ConstantPool) Opcode {
	opcodeValue := reader.ReadUint8(fieldName)
	opcodeName, ok := opcodeTable[opcodeValue]
	if !ok {
		return nil
	}

	parser, ok := opcodeParserTable[opcodeName]
	if !ok {
		return nil
	}

	return parser.Parse(fieldName, reader, opcodeValue, opcodeName, cp)
}
