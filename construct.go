package serr

import (
	"errors"
	"fmt"
	"runtime"
)

// New serr
func New(msg string) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2, stack[:])

	return construct(stack[:length], ErrLevelFatal, ErrCodeNothing, ErrKeyNothing, errors.New(msg), 0)
}

// Newf serr with message binding
func Newf(frmt string, args ...interface{}) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2, stack[:])

	return construct(stack[:length], ErrLevelFatal, ErrCodeNothing, ErrKeyNothing, fmt.Errorf(frmt, args...), 0)
}

// News serr from stack skip
func News(skip int, msg string) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2+skip, stack[:])

	return construct(stack[:length], ErrLevelFatal, ErrCodeNothing, ErrKeyNothing, errors.New(msg), skip)
}

// Newsl serr from stack skip and error level
func Newsl(skip int, lvl ErrLevel, msg string) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2+skip, stack[:])

	return construct(stack[:length], lvl, ErrCodeNothing, ErrKeyNothing, errors.New(msg), skip)
}

// Newsli serr from stack skip, error level and error code
func Newsli(skip int, lvl ErrLevel, code int, msg string) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2+skip, stack[:])

	return construct(stack[:length], lvl, code, ErrKeyNothing, errors.New(msg), skip)
}

// Newslik serr from stack skip, error level, error code and error key
func Newslik(skip int, lvl ErrLevel, code int, key string, msg string) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2+skip, stack[:])

	return construct(stack[:length], lvl, code, key, errors.New(msg), skip)
}

// Newslikc serr from stack skip, error level, error code, error key and comments
func Newslikc(skip int, lvl ErrLevel, code int, key string, msg string, comment string) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2+skip, stack[:])

	errx := construct(stack[:length], lvl, code, key, errors.New(msg), skip)

	if comment != "@" {
		errx.comments[0] = comment
	}
	return errx
}

// Newsi serr from stack skip and error code
func Newsi(skip int, code int, msg string) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2+skip, stack[:])

	return construct(stack[:length], ErrLevelFatal, code, ErrKeyNothing, errors.New(msg), skip)
}

// Newsik serr from stack skip, error code and error key
func Newsik(skip int, code int, key string, msg string) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2+skip, stack[:])

	return construct(stack[:length], ErrLevelFatal, code, key, errors.New(msg), skip)
}

// Newsikc serr from stack skip, error code, error key and comment
func Newsikc(skip int, code int, key string, msg string, comment string) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2+skip, stack[:])

	errx := construct(stack[:length], ErrLevelFatal, code, key, errors.New(msg), skip)

	if comment != "@" {
		errx.comments[0] = comment
	}
	return errx
}

// Newsk serr from stack skip and error key
func Newsk(skip int, key string, msg string) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2+skip, stack[:])

	return construct(stack[:length], ErrLevelFatal, ErrCodeNothing, key, errors.New(msg), skip)
}

// Newskc serr from stack skip, error key and comment
func Newskc(skip int, key string, msg string, comment string) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2+skip, stack[:])

	errx := construct(stack[:length], ErrLevelFatal, ErrCodeNothing, key, errors.New(msg), skip)

	if comment != "@" {
		errx.comments[0] = comment
	}
	return errx
}

// Newsf serr from stack skip with message binding
func Newsf(skip int, msg string, args ...interface{}) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2+skip, stack[:])

	return construct(stack[:length], ErrLevelFatal, ErrCodeNothing, ErrKeyNothing, fmt.Errorf(msg, args...), skip)
}

// Newsc serr from stack skip and comment
func Newsc(skip int, msg string, comment string) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2+skip, stack[:])

	errx := construct(stack[:length], ErrLevelFatal, ErrCodeNothing, ErrKeyNothing, errors.New(msg), skip)

	if comment != "@" {
		errx.comments[0] = comment
	}
	return errx
}

// Newl serr from error level
func Newl(lvl ErrLevel, msg string) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2, stack[:])

	return construct(stack[:length], lvl, ErrCodeNothing, ErrKeyNothing, errors.New(msg), 0)
}

// Newlf serr from error level with message binding
func Newlf(lvl ErrLevel, msg string, args ...interface{}) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2, stack[:])

	return construct(stack[:length], lvl, ErrCodeNothing, ErrKeyNothing, fmt.Errorf(msg, args...), 0)
}

// Newli serr from error level and error code
func Newli(lvl ErrLevel, code int, msg string) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2, stack[:])

	return construct(stack[:length], lvl, code, ErrKeyNothing, errors.New(msg), 0)
}

// Newlik serr from error level, error code and error key
func Newlik(lvl ErrLevel, code int, key string, msg string) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2, stack[:])

	return construct(stack[:length], lvl, code, key, errors.New(msg), 0)
}

// Newlikc serr from error level, error code, error key and comment
func Newlikc(lvl ErrLevel, code int, key string, msg string, comment string) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2, stack[:])

	errx := construct(stack[:length], lvl, code, key, errors.New(msg), 0)

	if comment != "@" {
		errx.comments[0] = comment
	}
	return errx
}

// Newi serr from error code
func Newi(code int, msg string) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2, stack[:])

	return construct(stack[:length], ErrLevelFatal, code, ErrKeyNothing, errors.New(msg), 0)
}

// Newif serr from error code with message binding
func Newif(code int, frmt string, args ...interface{}) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2, stack[:])

	return construct(stack[:length], ErrLevelFatal, code, ErrKeyNothing, fmt.Errorf(frmt, args...), 0)
}

// Newik serr from error code and error key
func Newik(code int, key string, msg string) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2, stack[:])

	return construct(stack[:length], ErrLevelFatal, code, key, errors.New(msg), 0)
}

// Newikc serr from error code, error key and comment
func Newikc(code int, key string, msg string, comment string) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2, stack[:])

	errx := construct(stack[:length], ErrLevelFatal, code, key, errors.New(msg), 0)

	if comment != "@" {
		errx.comments[0] = comment
	}
	return errx
}

// Newk serr from error key
func Newk(key string, msg string) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2, stack[:])

	return construct(stack[:length], ErrLevelFatal, ErrCodeNothing, key, errors.New(msg), 0)
}

// Newkf serr from error key with message binding
func Newkf(key string, frmt string, args ...interface{}) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2, stack[:])

	return construct(stack[:length], ErrLevelFatal, ErrCodeNothing, key, fmt.Errorf(frmt, args...), 0)
}

// Newkc serr from error key and comment
func Newkc(key string, msg string, comment string) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2, stack[:])

	errx := construct(stack[:length], ErrLevelFatal, ErrCodeNothing, key, errors.New(msg), 0)

	if comment != "@" {
		errx.comments[0] = comment
	}
	return errx
}

// Newc serr from comment
func Newc(msg string, comment string) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2, stack[:])

	errx := construct(stack[:length], ErrLevelFatal, ErrCodeNothing, ErrKeyNothing, errors.New(msg), 0)

	if comment != "@" {
		errx.comments[0] = comment
	}
	return errx
}

// Newslic serr from stack skip, error level, error code and comment
func Newslic(skip int, lvl ErrLevel, code int, msg string, comment string) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2+skip, stack[:])

	errx := construct(stack[:length], lvl, code, ErrKeyNothing, errors.New(msg), skip)

	if comment != "@" {
		errx.comments[0] = comment
	}
	return errx
}

// Newslk serr from stack skip, error level and error key
func Newslk(skip int, lvl ErrLevel, key string, msg string) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2+skip, stack[:])

	return construct(stack[:length], lvl, ErrCodeNothing, key, errors.New(msg), 0)
}

// Newsic serr from stack skip, error code and comment
func Newsic(skip int, code int, msg string, comment string) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2+skip, stack[:])

	errx := construct(stack[:length], ErrLevelFatal, code, ErrKeyNothing, errors.New(msg), skip)

	if comment != "@" {
		errx.comments[0] = comment
	}
	return errx
}

// Newslc serr from stack skip, error level and comment
func Newslc(skip int, lvl ErrLevel, msg string, comment string) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2+skip, stack[:])

	errx := construct(stack[:length], lvl, ErrCodeNothing, ErrKeyNothing, errors.New(msg), skip)

	if comment != "@" {
		errx.comments[0] = comment
	}
	return errx
}

// Newlkc serr from error level, error key and comment
func Newlkc(lvl ErrLevel, key string, msg string, comment string) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2, stack[:])

	errx := construct(stack[:length], lvl, ErrCodeNothing, key, errors.New(msg), 0)

	if comment != "@" {
		errx.comments[0] = comment
	}
	return errx
}

// Newlic serr from error level, error code and comment
func Newlic(lvl ErrLevel, code int, msg string, comment string) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2, stack[:])

	errx := construct(stack[:length], lvl, code, ErrKeyNothing, errors.New(msg), 0)

	if comment != "@" {
		errx.comments[0] = comment
	}
	return errx
}

// Newlk serr from error level and error key
func Newlk(lvl ErrLevel, key string, msg string) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2, stack[:])

	return construct(stack[:length], lvl, ErrCodeNothing, key, errors.New(msg), 0)
}

// Newlc serr from error level and comment
func Newlc(lvl ErrLevel, msg string, comment string) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2, stack[:])

	errx := construct(stack[:length], lvl, ErrCodeNothing, ErrKeyNothing, errors.New(msg), 0)

	if comment != "@" {
		errx.comments[0] = comment
	}
	return errx
}

// NewFromError serr from error
func NewFromError(err error) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2, stack[:])

	return construct(stack[:length], ErrLevelFatal, ErrCodeNothing, ErrKeyNothing, err, 0)
}

// NewFromErrors serr from error and stack skip
func NewFromErrors(skip int, err error) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2+skip, stack[:])

	return construct(stack[:length], ErrLevelFatal, ErrCodeNothing, ErrKeyNothing, err, skip)
}

// NewFromErrorl serr from error and error level
func NewFromErrorl(lvl ErrLevel, err error) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2, stack[:])

	return construct(stack[:length], lvl, ErrCodeNothing, ErrKeyNothing, err, 0)
}

// NewFromErrori serr from error and error code
func NewFromErrori(code int, err error) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2, stack[:])

	return construct(stack[:length], ErrLevelFatal, code, ErrKeyNothing, err, 0)
}

// NewFromErrork serr from error and error key
func NewFromErrork(key string, err error) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2, stack[:])

	return construct(stack[:length], ErrLevelFatal, ErrCodeNothing, key, err, 0)
}

// NewFromErrorc serr from error and comment
func NewFromErrorc(err error, comment string) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2, stack[:])

	errx := construct(stack[:length], ErrLevelFatal, ErrCodeNothing, ErrKeyNothing, err, 0)
	if comment != "@" {
		errx.comments[0] = comment
	}

	return errx
}

// NewFromErrorkc serr from error, error key, and comment
func NewFromErrorkc(err error, key string, comment string) SErr {
	stack := make([]uintptr, 50)
	length := runtime.Callers(2, stack[:])

	errx := construct(stack[:length], ErrLevelFatal, ErrCodeNothing, key, err, 0)
	if comment != "@" {
		errx.comments[0] = comment
	}

	return errx
}
