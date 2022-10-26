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

		Level() ErrorLevel
		Code() int
		Key() string
		Title() string
		Comments() string
		CommentStack() []string
		Payload() ErrorPayload

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
		SetLevel(lvl ErrorLevel)
		AddComment(msg string)
		AddComments(skip int, msg string)
		AddCommentf(msg string, opts ...interface{})
		ApplyPayload(payload ErrorPayload)
		SetPayload(key string, value interface{})
	}

	// ErrorPayload type
	ErrorPayload map[string]interface{}

	// ErrorLevel type
	ErrorLevel string
)

type serr struct {
	level    ErrorLevel
	err      error
	key      string
	code     int
	comments []string
	payload  ErrorPayload
	frames   []gerr.StackFrame
	stack    []uintptr
}

const (
	// ErrorLevelFatal for fatal error
	ErrorLevelFatal ErrorLevel = "fatal"

	// ErrorLevelValidation for validation error
	ErrorLevelValidation ErrorLevel = "warn"

	// ErrorLevelWarn for warning error
	ErrorLevelWarn ErrorLevel = "warn"

	// ErrorLevelInfo for information error
	ErrorLevelInfo ErrorLevel = "info"
)

const (
	// ErrorKeyNothing constant for empty error key
	ErrorKeyNothing = "-"

	// ErrorKeyUnexpected constant for unexpected error key
	ErrorKeyUnexpected = "unexpected"

	// ErrorKeyExpected constant for expected error key
	ErrorKeyExpected = "expected"

	// ErrorCodeNothing constant for empty error code
	ErrorCodeNothing = 0
)

var rootPaths []string

// RegisterRootPath to registering the root path.
func RegisterRootPath(paths []string) {
	rootPaths = append(rootPaths, paths...)
}

// RegisterHereAsRootPath to registering here (current file path) as a root path.
func RegisterHereAsRootPath(callerSkip int, pathSkip int) SErr {
	_, file, _, ok := runtime.Caller(callerSkip + 1)
	if !ok {
		return New("Failed to get path")
	}

	sep := "/"
	if runtime.GOOS == "windows" {
		sep = "\\"
	}

	file = path.Dir(file)
	paths := strings.Split(file, sep)
	if len(paths) > pathSkip {
		paths = paths[:len(paths)-pathSkip]
	}
	RegisterRootPath([]string{strings.Join(paths, sep)})

	return nil
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
func (ox serr) Level() ErrorLevel {
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
func (ox serr) Payload() ErrorPayload {
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
	var comments string

	if ox.Code() != 0 {
		comments += fmt.Sprintf(" <code: %d>", ox.Code())
	}

	if isExists(ox.Key(), []string{"-", "!", ""}) {
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
	errorMessage := ox.Error()
	if len(ox.comments) > 0 {
		errorMessage = fmt.Sprintf("%s, detail: %s [%s:%d]", ox.Comments(), errorMessage, ox.File(), ox.Line())
	}

	return errorMessage
}

// ColoredString to get formated error message with color (cli color code)
func (ox serr) ColoredString() string {
	var comments string

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
func (ox *serr) SetLevel(level ErrorLevel) {
	ox.level = level
}

func (ox *serr) addRawComment(message string, callerSkip int) {
	if len(message) <= 0 {
		return
	}

	if len(ox.comments) <= 0 {
		ox.comments = append(ox.comments, strings.ToUpper(string(message[0]))+string(message[1:]))
		return
	}

	_, file, line, _ := runtime.Caller(callerSkip + 1)
	ox.comments = append(ox.comments, fmt.Sprintf("%s on [%s:%d]", message, resolvePath(file), line))
}

// AddComment to add error comment
func (ox *serr) AddComment(message string) {
	ox.addRawComment(message, 1)
}

// AddComments to add error comment with skip
func (ox *serr) AddComments(callerSkip int, message string) {
	ox.addRawComment(message, 1+callerSkip)
}

// AddCommentf to add error comment with string binding
func (ox *serr) AddCommentf(message string, args ...interface{}) {
	ox.addRawComment(fmt.Sprintf(message, args...), 1)
}

// ApplyPayload to apply error payload
func (ox *serr) ApplyPayload(payload ErrorPayload) {
	if ox.payload == nil {
		ox.payload = make(ErrorPayload)
	}

	if payload != nil {
		for k, v := range payload {
			ox.payload[k] = v
		}
	}
}

// SetPayload to set error payload
func (ox *serr) SetPayload(key string, value interface{}) {
	if ox.payload == nil {
		ox.payload = make(ErrorPayload)
	}

	ox.payload[key] = value
}
