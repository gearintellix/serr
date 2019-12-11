package serr

import (
	"errors"
	"fmt"
)

// New serr from raw message
func New(msg string) SErr {
	return construct(ErrLevelFatal, ErrCodeNothing, ErrKeyNothing, errors.New(msg))
}

func Newf(frmt string, args ...interface{}) SErr {
	return construct(ErrLevelFatal, ErrCodeNothing, ErrKeyNothing, fmt.Errorf(frmt, args...))
}

func Newl(lvl ErrLevel, msg string) SErr {
	return construct(lvl, ErrCodeNothing, ErrKeyNothing, errors.New(msg))
}

func Newi(code int, msg string) SErr {
	return construct(ErrLevelFatal, code, ErrKeyNothing, errors.New(msg))
}

func Newk(key string, msg string) SErr {
	return construct(ErrLevelFatal, ErrCodeNothing, key, errors.New(msg))
}

func Newc(msg string, comment string) SErr {
	errx := construct(ErrLevelFatal, ErrCodeNothing, ErrKeyNothing, errors.New(msg))
	errx.AddComment(comment)
	return errx
}

func Newlf(lvl ErrLevel, frmt string, args ...interface{}) SErr {
	return construct(lvl, ErrCodeNothing, ErrKeyNothing, fmt.Errorf(frmt, args...))
}

func Newli(lvl ErrLevel, code int, msg string) SErr {
	return construct(lvl, code, ErrKeyNothing, errors.New(msg))
}

func Newlk(lvl ErrLevel, key string, msg string) SErr {
	return construct(lvl, ErrCodeNothing, key, errors.New(msg))
}

func Newlc(lvl ErrLevel, msg string, comment string) SErr {
	errx := construct(lvl, ErrCodeNothing, ErrKeyNothing, errors.New(msg))
	errx.AddComment(comment)
	return errx
}

func Newif(code int, frmt string, args ...interface{}) SErr {
	return construct(ErrLevelFatal, code, ErrKeyNothing, fmt.Errorf(frmt, args...))
}

func Newik(code int, key string, msg string) SErr {
	return construct(ErrLevelFatal, code, key, errors.New(msg))
}

func Newic(code int, msg string, comment string) SErr {
	errx := construct(ErrLevelFatal, code, ErrKeyNothing, errors.New(msg))
	errx.AddComment(comment)
	return errx
}

func Newkf(key string, frmt string, args ...interface{}) SErr {
	return construct(ErrLevelFatal, ErrCodeNothing, key, fmt.Errorf(frmt, args...))
}

func Newkc(key string, msg string, comment string) SErr {
	errx := construct(ErrLevelFatal, ErrCodeNothing, key, errors.New(msg))
	errx.AddComment(comment)
	return errx
}

func NewFromError(err error) SErr {
	return construct(ErrLevelFatal, ErrCodeNothing, ErrKeyNothing, err)
}

func NewFromErrorl(lvl ErrLevel, err error) SErr {
	return construct(lvl, ErrCodeNothing, ErrKeyNothing, err)
}

func NewFromErrori(code int, err error) SErr {
	return construct(ErrLevelFatal, code, ErrKeyNothing, err)
}

func NewFromErrork(key string, err error) SErr {
	return construct(ErrLevelFatal, ErrCodeNothing, key, err)
}

func NewFromErrorc(err error, comment string) SErr {
	errx := construct(ErrLevelFatal, ErrCodeNothing, ErrKeyNothing, err)
	errx.AddComment(comment)
	return errx
}
