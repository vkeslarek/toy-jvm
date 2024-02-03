package types

type AccessFlagName = string

const (
	AccessFlagPublic       AccessFlagName = "ACC_PUBLIC"
	AccessFlagPrivate      AccessFlagName = "ACC_PRIVATE"
	AccessFlagProtected    AccessFlagName = "ACC_PROTECTED"
	AccessFlagStatic       AccessFlagName = "ACC_STATIC"
	AccessFlagFinal        AccessFlagName = "ACC_FINAL"
	AccessFlagSynchronized AccessFlagName = "ACC_SYNCHRONIZED"
	AccessFlagBridge       AccessFlagName = "ACC_BRIDGE"
	AccessFlagVarargs      AccessFlagName = "ACC_VARARGS"
	AccessFlagNative       AccessFlagName = "ACC_NATIVE"
	AccessFlagVolatile     AccessFlagName = "ACC_VOLATILE"
	AccessFlagTransient    AccessFlagName = "ACC_TRANSIENT"
	AccessFlagSuper        AccessFlagName = "ACC_SUPER"
	AccessFlagInterface    AccessFlagName = "ACC_INTERFACE"
	AccessFlagAbstract     AccessFlagName = "ACC_ABSTRACT"
	AccessFlagStrict       AccessFlagName = "ACC_STRICT"
	AccessFlagSynthetic    AccessFlagName = "ACC_SYNTHETIC"
	AccessFlagMandated     AccessFlagName = "ACC_MANDATED"
	AccessFlagAnnotation   AccessFlagName = "ACC_ANNOTATION"
	AccessFlagEnum         AccessFlagName = "ACC_ENUM"
	AccessFlagModule       AccessFlagName = "ACC_MODULE"
)
