package errors

import (
	"casher-server/internal/i18n"
	"errors"
)

type IErrorL interface {
	Error() string
	New(id string, defaultMsg string, args ...string) *ErrorL
	FromError(err error) *ErrorL
}

func NewErrorL(messageId string, defaultMsg string) *ErrorL {
	return &ErrorL{
		MessageId: messageId,
		Args:      []i18n.Field{},
	}
}

func FromLang(lang i18n.IParse) *ErrorL {
	return &ErrorL{
		Lang: lang,
		Args: []i18n.Field{},
	}
}

type ErrorL struct {
	Lang       i18n.IParse
	MessageId  string       `json:"id"`
	DefaultMsg string       `json:"-"`
	Args       []i18n.Field `json:"-"`
}

func (e *ErrorL) Error() string {
	return e.ei18n().Error()
}

func (e *ErrorL) New(id string, defaultMsg string, args ...i18n.Field) *ErrorL {
	e.MessageId = id
	e.DefaultMsg = defaultMsg
	e.Args = args
	return e
}

func (e *ErrorL) ei18n() error {
	return errors.New(e.Lang.L(e.MessageId, e.DefaultMsg, e.Args...))
}
