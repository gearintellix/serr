package serr

import (
	"fmt"
	"path"
	"reflect"
	"runtime"
	"strings"

	gerr "github.com/go-errors/errors"
)

type (
	// SErr interface
	SErr interface {
		Error() string
		Cause() error

		Level() ErrLevel
		Code() int
		Key() string
		Title() string
		Comments() string
		CommentStack() []string
		Payload() ErrPayload

		Callers() []uintptr
		StackFrames() []gerr.StackFrame

		Type() string
		File() string
		Line() int
		FN() string
		Package() string

		String() string
		ColoredString() string
		SimpleString() string

		SetKey(key string)
		SetCode(code int)
		SetLevel(lvl ErrLevel)
		AddComment(msg string)
		AddComments(skip int, msg string)
		AddCommentf(msg string, opts ...interface{})
		ApplyPayload(payload ErrPayload)
		SetPayload(key string, value interface{})
	}

	// ErrPayload type
	ErrPayload map[string]interface{}

	// ErrLevel type
	ErrLevel string
)

type serr struct {
	level    ErrLevel
	err      error
	key      string
	code     int
	comments []string
	payload  ErrPayload
	frames   []gerr.StackFrame
	stack    []uintptr
}

const (
	// ErrLevelFatal constant for fatal error level
	ErrLevelFatal ErrLevel = "fatal"

	// ErrLevelWarn constant for warning error level
	ErrLevelWarn ErrLevel = "warn"

	// ErrLevelInfo constant for info error level
	ErrLevelInfo ErrLevel = "info"
)

const (
	// ErrKeyNothing constant for empty error key
	ErrKeyNothing string = "-"

	// ErrKeyUnexpected constant for unexpected error key
	ErrKeyUnexpected string = "unexpected"

	// ErrKeyExpected constant for expected error key
	ErrKeyExpected string = "expected"

	// ErrCodeNothing constant for empty error code
	ErrCodeNothing int = 0
)

var (
	rootPaths []string
)

// RegisterRootPath function
func RegisterRootPath(paths []string) {
	rootPaths = append(rootPaths, paths...)
}

// RegisterThisAsRoot function
func RegisterThisAsRoot(cskip int, pskip int) SErr {
	_, file, _, ok := runtime.Caller(cskip + 1)
	if !ok {
		return New("Failed to get path")
	}

	sep := "/"
	if runtime.GOOS == "windows" {
		sep = "\\"
	}

	file = path.Dir(file)
	paths := strings.Split(file, sep)
	if len(paths) > pskip {
		paths = paths[:len(paths)-pskip]
	}
	RegisterRootPath([]string{strings.Join(paths, sep)})

	return nil
}

func construct(stack []uintptr, level ErrLevel, code int, key string, err error, skip int) *serr {
	res := &serr{
		level:    level,
		err:      err,
		key:      key,
		code:     code,
		comments: []string{},
		payload:  make(ErrPayload),
		stack:    stack,
	}
	res.addRawComment(err.Error(), skip+1)
	return res
}

// Error to get error message
func (ox serr) Error() string {
	return fmt.Sprintf("%+v", ox.err)
}

// Cause to get original error
func (ox serr) Cause() error {
	return ox.err
}

// Level to get error level
func (ox serr) Level() ErrLevel {
	return ox.level
}

// Code to get error code
func (ox serr) Code() int {
	return ox.code
}

// Key to get error key
func (ox serr) Key() string {
	return ox.key
}

// Title to get error title
func (ox serr) Title() string {
	if len(ox.comments) > 0 {
		return ox.comments[0]
	}
	return ox.Error()
}

// Comments to get error comments
func (ox serr) Comments() string {
	return strings.Join(ox.comments, ", ")
}

// CommentStack to get error comment stack
func (ox serr) CommentStack() []string {
	return ox.comments
}

// Payload to get error payload
func (ox serr) Payload() ErrPayload {
	return ox.payload
}

// Callers to get error callers stack
func (ox serr) Callers() []uintptr {
	return ox.stack
}

// StackFrames to get error stack frames
func (ox *serr) StackFrames() []gerr.StackFrame {
	if ox.frames == nil {
		ox.frames = make([]gerr.StackFrame, len(ox.stack))

		for i, pc := range ox.stack {
			item := gerr.NewStackFrame(pc)
			item.File = resolvePath(item.File)
			ox.frames[i] = item
		}
	}

	return ox.frames
}

// Type get error type
func (ox serr) Type() string {
	return reflect.TypeOf(ox.err).String()
}

// File get error file path
func (ox serr) File() string {
	frames := ox.StackFrames()
	if len(frames) > 0 {
		return frames[0].File
	}
	return ""
}

// Line get error line
func (ox serr) Line() int {
	frames := ox.StackFrames()
	if len(frames) > 0 {
		return frames[0].LineNumber
	}
	return 0
}

// FN to get error function name
func (ox serr) FN() string {
	frames := ox.StackFrames()
	if len(frames) > 0 {
		return frames[0].Name
	}
	return ""
}

// Package to get error package name
func (ox serr) Package() string {
	frames := ox.StackFrames()
	if len(frames) > 0 {
		return frames[0].Package
	}
	return ""
}

// String to get formated error message
func (ox serr) String() string {
	comments := ""

	if ox.Code() != 0 {
		comments += fmt.Sprintf(" <code: %d>", ox.Code())
	}

	if isExists(ox.Key(), []string{"-", "!"}) {
		comments += fmt.Sprintf(" <key: %s>", ox.Key())
	}

	if len(ox.comments) > 0 {
		comments += fmt.Sprintf(" <comments: %s>", ox.Comments())
	}

	return fmt.Sprintf(
		StandardFormat(),
		ox.FN(),
		ox.File(),
		ox.Line(),
		ox.Error(),
		comments,
	)
}

// SimpleString to get simple formatted error message
func (ox serr) SimpleString() string {
	msg := ox.Error()
	if len(ox.comments) > 0 {
		msg = fmt.Sprintf("%s, detail: %s [%s:%d]", ox.Comments(), msg, ox.File(), ox.Line())
	}

	return msg
}

// ColoredString to get formated error message with color (cli color code)
func (ox serr) ColoredString() string {
	comments := ""

	if ox.Code() != 0 {
		comments += fmt.Sprintf(" <code: %d>", ox.Code())
	}

	if !isExists(ox.Key(), []string{"", "-", "!"}) {
		comments += fmt.Sprintf(" <key: %s>", ox.Key())
	}

	if len(ox.comments) > 0 {
		comments += fmt.Sprintf(" <comments: %s>", ox.Comments())
	}

	return fmt.Sprintf(
		StandardColorFormat(),
		ox.FN(),
		ox.File(),
		ox.Line(),
		ox.Error(),
		comments,
	)
}

// SetKey to set error key
func (ox *serr) SetKey(key string) {
	ox.key = key
}

// SetCode to set error code
func (ox *serr) SetCode(code int) {
	ox.code = code
}

// SetLevel to set error level
func (ox *serr) SetLevel(lvl ErrLevel) {
	ox.level = lvl
}

func (ox *serr) addRawComment(note string, skip int) {
	if len(note) <= 0 {
		return
	}

	if len(ox.comments) <= 0 {
		ox.comments = append(ox.comments, strings.ToUpper(string(note[0]))+string(note[1:]))
		return
	}

	_, file, line, _ := runtime.Caller(skip + 1)
	ox.comments = append(ox.comments, fmt.Sprintf("%s on [%s:%d]", note, resolvePath(file), line))
}

// AddComment to add error comment
func (ox *serr) AddComment(msg string) {
	ox.addRawComment(msg, 1)
}

// AddComments to add error comment with skip
func (ox *serr) AddComments(skip int, msg string) {
	ox.addRawComment(msg, 1+skip)
}

// AddCommentf to add error comment with string binding
func (ox *serr) AddCommentf(msg string, opts ...interface{}) {
	ox.addRawComment(fmt.Sprintf(msg, opts...), 1)
}

// ApplyPayload to apply error payload
func (ox *serr) ApplyPayload(load ErrPayload) {
	if ox.payload == nil {
		ox.payload = make(ErrPayload)
	}

	if load != nil {
		for k, v := range load {
			ox.payload[k] = v
		}
	}
}

// SetPayload to set error payload
func (ox *serr) SetPayload(key string, value interface{}) {
	if ox.payload == nil {
		ox.payload = make(ErrPayload)
	}

	ox.payload[key] = value
}
