package errorhandler

type SilentExecution interface {
	SetError(ParserError)
	GetError() ParserError
}

func RunSilent[T any](silentExecution SilentExecution, fieldName string, fn func() (T, error)) T {
	val, err := fn()

	if err != nil {
		silentExecution.SetError(Wrap(fieldName, err))
		// panic(Wrap(fieldName, err))
	}

	return val
}

func RunSilentVoid[T any](silentExecution SilentExecution, fieldName string, fn func() error) {
	if err := fn(); err != nil {
		silentExecution.SetError(Wrap(fieldName, err))
	}
}
