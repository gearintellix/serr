package serr

type (
	// ErrPayload type
	//
	// Deprecated, use ErrorPayload instead!
	ErrPayload = ErrorPayload

	// ErrLevel type
	//
	// Deprecated, use ErrorLevel instead!
	ErrLevel = ErrorLevel
)

const (
	// ErrLevelFatal constant for fatal error level
	//
	// Deprecated, use ErrorLevelFatal instead!
	ErrLevelFatal = ErrorLevelFatal

	// ErrLevelWarn constant for warning error level
	//
	// Deprecated, use ErrorLevelWarn instead!
	ErrLevelWarn = ErrorLevelWarn

	// ErrLevelInfo constant for info error level
	//
	// Deprecated, use ErrorLevelInfo instead!
	ErrLevelInfo = ErrorLevelInfo

	// ErrKeyNothing constant for empty error key
	//
	// Deprecated, use ErrorKeyNothing instead!
	ErrKeyNothing = ErrorKeyNothing

	// ErrKeyUnexpected constant for unexpected error key
	//
	// Deprecated, use ErrorKeyUnexpected instead!
	ErrKeyUnexpected = ErrorKeyUnexpected

	// ErrKeyExpected constant for expected error key
	//
	// Deprecated, use ErrorKeyExpected instead!
	ErrKeyExpected = ErrorKeyExpected

	// ErrCodeNothing constant for empty error code
	//
	// Deprecated, use ErrorCodeNothing instead!
	ErrCodeNothing = ErrorCodeNothing
)

// RegisterThisAsRoot to registering here (current file path) as a root path.
//
// Deprecated, use RegisterHereAsRootPath instead!
func RegisterThisAsRoot(callerSkip int, pathSkip int) SErr {
	return RegisterHereAsRootPath(callerSkip, pathSkip)
}
