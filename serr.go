package serr

import (
	"fmt"
	"runtime"
	"strings"
)

// SErr interface
type SErr interface {
	Level() ErrLevel
	Code() int
	Key() string
	Error() error
	String() string
	Comments() string
	CommentStack() []string
	Payload() ErrPayload
	File() string
	Line() int
	FN() string

	SetKey(key string)
	SetCode(code int)
	SetLevel(lvl ErrLevel)
	AddComment(msg string)
	AddCommentf(msg string, opts ...interface{})
	ApplyPayload(payload ErrPayload)
	SetPayload(key string, value interface{})

	Sprint() string
	SprintWithColor() string
}

type serr struct {
	level    ErrLevel
	err      error
	key      string
	code     int
	comments []string
	payload  ErrPayload
	file     string
	line     int
	fn       string
}

// ErrPayload type
type ErrPayload map[string]interface{}

// ErrLevel type
type ErrLevel string

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

func construct(level ErrLevel, code int, key string, err error) SErr {
	pc, fileName, line, _ := runtime.Caller(2)

	return &serr{
		level:    level,
		err:      err,
		key:      key,
		code:     code,
		comments: []string{},
		payload:  make(ErrPayload),
		file:     fileName,
		line:     line,
		fn:       runtime.FuncForPC(pc).Name(),
	}
}

func (ox serr) Level() ErrLevel {
	return ox.level
}

func (ox serr) Code() int {
	return ox.code
}

func (ox serr) Key() string {
	return ox.key
}

func (ox serr) Error() error {
	return ox.err
}

func (ox serr) String() string {
	return ox.err.Error()
}

func (ox serr) Comments() string {
	return strings.Join(ox.comments, ", ")
}

func (ox serr) CommentStack() []string {
	return ox.comments
}

func (ox serr) File() string {
	return ox.file
}

func (ox serr) Payload() ErrPayload {
	return ox.payload
}

func (ox serr) Line() int {
	return ox.line
}

func (ox serr) FN() string {
	return ox.fn
}

func (ox *serr) SetKey(key string) {
	ox.key = key
}

func (ox *serr) SetCode(code int) {
	ox.code = code
}

func (ox *serr) SetLevel(lvl ErrLevel) {
	ox.level = lvl
}

func (ox *serr) AddComment(msg string) {
	ox.comments = append(ox.comments, msg)
}

func (ox *serr) AddCommentf(msg string, opts ...interface{}) {
	ox.comments = append(ox.comments, fmt.Sprintf(msg, opts))
}

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

func (ox *serr) SetPayload(key string, value interface{}) {
	if ox.payload == nil {
		ox.payload = make(ErrPayload)
	}

	ox.payload[key] = value
}

func (ox serr) Sprint() string {
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
		GetStandardFormat(),
		ox.FN(),
		ox.File(),
		ox.Line(),
		ox.String(),
		comments,
	)
}

func (ox serr) SprintWithColor() string {
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
		GetStandardColorFormat(),
		ox.FN(),
		ox.File(),
		ox.Line(),
		ox.String(),
		comments,
	)
}

func (ox serr) JSON() string {
	// TODO: this still draft
	return ""
}

func (ox serr) Panic() {
	// TODO: this still draft
}
