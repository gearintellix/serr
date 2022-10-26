package serr

import (
	"errors"
	"fmt"
	"runtime"
)

func construct(stack []uintptr, level ErrorLevel, code int, key string, err error, callerSkip int) *serr {
	errx := &serr{
		level:    level,
		err:      err,
		key:      key,
		code:     code,
		comments: []string{},
		payload:  make(ErrorPayload),
		stack:    stack,
	}
	errx.addRawComment(err.Error(), callerSkip+1)
	return errx
}

// New serr
func New(message string) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2, stack[:])
	)

	return construct(stack[:length], ErrorLevelFatal, ErrorCodeNothing, ErrorKeyNothing, errors.New(message), 0)
}

// Newf serr with message binding
func Newf(format string, args ...interface{}) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2, stack[:])
	)

	return construct(stack[:length], ErrorLevelFatal, ErrorCodeNothing, ErrorKeyNothing, fmt.Errorf(format, args...), 0)
}

// News serr from stack callerSkip
func News(callerSkip int, message string) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2+callerSkip, stack[:])
	)

	return construct(stack[:length], ErrorLevelFatal, ErrorCodeNothing, ErrorKeyNothing, errors.New(message), callerSkip)
}

// Newsl serr from stack callerSkip and error level
func Newsl(callerSkip int, level ErrorLevel, message string) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2+callerSkip, stack[:])
	)

	return construct(stack[:length], level, ErrorCodeNothing, ErrorKeyNothing, errors.New(message), callerSkip)
}

// Newsli serr from stack callerSkip, error level and error code
func Newsli(callerSkip int, level ErrorLevel, code int, message string) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2+callerSkip, stack[:])
	)

	return construct(stack[:length], level, code, ErrorKeyNothing, errors.New(message), callerSkip)
}

// Newslik serr from stack callerSkip, error level, error code and error key
func Newslik(callerSkip int, level ErrorLevel, code int, key string, message string) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2+callerSkip, stack[:])
	)

	return construct(stack[:length], level, code, key, errors.New(message), callerSkip)
}

// Newslikc serr from stack callerSkip, error level, error code, error key and title
func Newslikc(callerSkip int, level ErrorLevel, code int, key string, message string, title string) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2+callerSkip, stack[:])
		errx   = construct(stack[:length], level, code, key, errors.New(message), callerSkip)
	)

	if title != "@" {
		errx.comments[0] = title
	}
	return errx
}

// Newsi serr from stack callerSkip and error code
func Newsi(callerSkip int, code int, message string) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2+callerSkip, stack[:])
	)

	return construct(stack[:length], ErrorLevelFatal, code, ErrorKeyNothing, errors.New(message), callerSkip)
}

// Newsik serr from stack callerSkip, error code and error key
func Newsik(callerSkip int, code int, key string, message string) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2+callerSkip, stack[:])
	)

	return construct(stack[:length], ErrorLevelFatal, code, key, errors.New(message), callerSkip)
}

// Newsikc serr from stack callerSkip, error code, error key and title
func Newsikc(callerSkip int, code int, key string, message string, title string) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2+callerSkip, stack[:])
		errx   = construct(stack[:length], ErrorLevelFatal, code, key, errors.New(message), callerSkip)
	)

	if title != "@" {
		errx.comments[0] = title
	}
	return errx
}

// Newsk serr from stack callerSkip and error key
func Newsk(callerSkip int, key string, message string) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2+callerSkip, stack[:])
	)

	return construct(stack[:length], ErrorLevelFatal, ErrorCodeNothing, key, errors.New(message), callerSkip)
}

// Newskc serr from stack callerSkip, error key and title
func Newskc(callerSkip int, key string, message string, title string) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2+callerSkip, stack[:])
		errx   = construct(stack[:length], ErrorLevelFatal, ErrorCodeNothing, key, errors.New(message), callerSkip)
	)

	if title != "@" {
		errx.comments[0] = title
	}
	return errx
}

// Newsf serr from stack callerSkip with message binding
func Newsf(callerSkip int, message string, args ...interface{}) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2+callerSkip, stack[:])
	)

	return construct(stack[:length], ErrorLevelFatal, ErrorCodeNothing, ErrorKeyNothing, fmt.Errorf(message, args...), callerSkip)
}

// Newsc serr from stack callerSkip and title
func Newsc(callerSkip int, message string, title string) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2+callerSkip, stack[:])
		errx   = construct(stack[:length], ErrorLevelFatal, ErrorCodeNothing, ErrorKeyNothing, errors.New(message), callerSkip)
	)

	if title != "@" {
		errx.comments[0] = title
	}
	return errx
}

// Newl serr from error level
func Newl(level ErrorLevel, message string) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2, stack[:])
	)

	return construct(stack[:length], level, ErrorCodeNothing, ErrorKeyNothing, errors.New(message), 0)
}

// Newlf serr from error level with message binding
func Newlf(level ErrorLevel, message string, args ...interface{}) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2, stack[:])
	)

	return construct(stack[:length], level, ErrorCodeNothing, ErrorKeyNothing, fmt.Errorf(message, args...), 0)
}

// Newli serr from error level and error code
func Newli(level ErrorLevel, code int, message string) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2, stack[:])
	)

	return construct(stack[:length], level, code, ErrorKeyNothing, errors.New(message), 0)
}

// Newlik serr from error level, error code and error key
func Newlik(level ErrorLevel, code int, key string, message string) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2, stack[:])
	)

	return construct(stack[:length], level, code, key, errors.New(message), 0)
}

// Newlikc serr from error level, error code, error key and title
func Newlikc(level ErrorLevel, code int, key string, message string, title string) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2, stack[:])
		errx   = construct(stack[:length], level, code, key, errors.New(message), 0)
	)

	if title != "@" {
		errx.comments[0] = title
	}
	return errx
}

// Newi serr from error code
func Newi(code int, message string) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2, stack[:])
	)

	return construct(stack[:length], ErrorLevelFatal, code, ErrorKeyNothing, errors.New(message), 0)
}

// Newif serr from error code with message binding
func Newif(code int, frmt string, args ...interface{}) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2, stack[:])
	)

	return construct(stack[:length], ErrorLevelFatal, code, ErrorKeyNothing, fmt.Errorf(frmt, args...), 0)
}

// Newik serr from error code and error key
func Newik(code int, key string, message string) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2, stack[:])
	)

	return construct(stack[:length], ErrorLevelFatal, code, key, errors.New(message), 0)
}

// Newikc serr from error code, error key and title
func Newikc(code int, key string, message string, title string) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2, stack[:])
		errx   = construct(stack[:length], ErrorLevelFatal, code, key, errors.New(message), 0)
	)

	if title != "@" {
		errx.comments[0] = title
	}
	return errx
}

// Newk serr from error key
func Newk(key string, message string) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2, stack[:])
	)

	return construct(stack[:length], ErrorLevelFatal, ErrorCodeNothing, key, errors.New(message), 0)
}

// Newkf serr from error key with message binding
func Newkf(key string, frmt string, args ...interface{}) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2, stack[:])
	)

	return construct(stack[:length], ErrorLevelFatal, ErrorCodeNothing, key, fmt.Errorf(frmt, args...), 0)
}

// Newkc serr from error key and title
func Newkc(key string, message string, title string) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2, stack[:])
		errx   = construct(stack[:length], ErrorLevelFatal, ErrorCodeNothing, key, errors.New(message), 0)
	)

	if title != "@" {
		errx.comments[0] = title
	}
	return errx
}

// Newc serr from title
func Newc(message string, title string) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2, stack[:])
		errx   = construct(stack[:length], ErrorLevelFatal, ErrorCodeNothing, ErrorKeyNothing, errors.New(message), 0)
	)

	if title != "@" {
		errx.comments[0] = title
	}
	return errx
}

// Newslic serr from stack callerSkip, error level, error code and title
func Newslic(callerSkip int, level ErrorLevel, code int, message string, title string) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2+callerSkip, stack[:])
		errx   = construct(stack[:length], level, code, ErrorKeyNothing, errors.New(message), callerSkip)
	)

	if title != "@" {
		errx.comments[0] = title
	}
	return errx
}

// Newslk serr from stack callerSkip, error level and error key
func Newslk(callerSkip int, level ErrorLevel, key string, message string) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2+callerSkip, stack[:])
	)

	return construct(stack[:length], level, ErrorCodeNothing, key, errors.New(message), 0)
}

// Newsic serr from stack callerSkip, error code and title
func Newsic(callerSkip int, code int, message string, title string) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2+callerSkip, stack[:])
	)

	errx := construct(stack[:length], ErrorLevelFatal, code, ErrorKeyNothing, errors.New(message), callerSkip)

	if title != "@" {
		errx.comments[0] = title
	}
	return errx
}

// Newslc serr from stack callerSkip, error level and title
func Newslc(callerSkip int, level ErrorLevel, message string, title string) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2+callerSkip, stack[:])
		errx   = construct(stack[:length], level, ErrorCodeNothing, ErrorKeyNothing, errors.New(message), callerSkip)
	)

	if title != "@" {
		errx.comments[0] = title
	}
	return errx
}

// Newlkc serr from error level, error key and title
func Newlkc(level ErrorLevel, key string, message string, title string) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2, stack[:])
		errx   = construct(stack[:length], level, ErrorCodeNothing, key, errors.New(message), 0)
	)

	if title != "@" {
		errx.comments[0] = title
	}
	return errx
}

// Newlic serr from error level, error code and title
func Newlic(level ErrorLevel, code int, message string, title string) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2, stack[:])
		errx   = construct(stack[:length], level, code, ErrorKeyNothing, errors.New(message), 0)
	)

	if title != "@" {
		errx.comments[0] = title
	}
	return errx
}

// Newlk serr from error level and error key
func Newlk(level ErrorLevel, key string, message string) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2, stack[:])
	)

	return construct(stack[:length], level, ErrorCodeNothing, key, errors.New(message), 0)
}

// Newlc serr from error level and title
func Newlc(level ErrorLevel, message string, title string) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2, stack[:])
		errx   = construct(stack[:length], level, ErrorCodeNothing, ErrorKeyNothing, errors.New(message), 0)
	)

	if title != "@" {
		errx.comments[0] = title
	}
	return errx
}

// NewFromError serr from error
func NewFromError(err error) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2, stack[:])
	)

	return construct(stack[:length], ErrorLevelFatal, ErrorCodeNothing, ErrorKeyNothing, err, 0)
}

// NewFromErrors serr from error and stack callerSkip
func NewFromErrors(callerSkip int, err error) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2+callerSkip, stack[:])
	)

	return construct(stack[:length], ErrorLevelFatal, ErrorCodeNothing, ErrorKeyNothing, err, callerSkip)
}

// NewFromErrorl serr from error and error level
func NewFromErrorl(level ErrorLevel, err error) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2, stack[:])
	)

	return construct(stack[:length], level, ErrorCodeNothing, ErrorKeyNothing, err, 0)
}

// NewFromErrori serr from error and error code
func NewFromErrori(code int, err error) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2, stack[:])
	)

	return construct(stack[:length], ErrorLevelFatal, code, ErrorKeyNothing, err, 0)
}

// NewFromErrork serr from error and error key
func NewFromErrork(key string, err error) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2, stack[:])
	)

	return construct(stack[:length], ErrorLevelFatal, ErrorCodeNothing, key, err, 0)
}

// NewFromErrorc serr from error and title
func NewFromErrorc(err error, title string) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2, stack[:])
		errx   = construct(stack[:length], ErrorLevelFatal, ErrorCodeNothing, ErrorKeyNothing, err, 0)
	)

	if title != "@" {
		errx.comments[0] = title
	}

	return errx
}

// NewFromErrorkc serr from error, error key, and title
func NewFromErrorkc(err error, key string, title string) SErr {
	var (
		stack  = make([]uintptr, 50)
		length = runtime.Callers(2, stack[:])
		errx   = construct(stack[:length], ErrorLevelFatal, ErrorCodeNothing, key, err, 0)
	)

	if title != "@" {
		errx.comments[0] = title
	}

	return errx
}
