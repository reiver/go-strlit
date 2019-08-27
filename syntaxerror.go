package strlit

import (
	"fmt"
)

// SyntaxError is an error that is used to represent a syntax error in a
// string literal.
//
// SyntaxError allows one to respond to errors returned from strlit.Compile()
// in a more precise way. For example:
//
//	compiled, err := strlit.Compile(runeReader)
//	if nil != err {
//		switch err.(type) {
//		case strlit.SyntaxError:
//			//@TODO
//		default:
//			//@TODO
//		}
//	}
//
// An error message can be returned from the Error() method. In addition to this, one
// can construct their own error message using the Code() method.
//
// The Code() method returns "code" from the string literal that caused the error.
type SyntaxError interface {
	error

	// SyntaxError is used for typing. Calling it will do nothing.
	SyntaxError()

	// Code returns "code" from the string literal that caused the error.
	Code() string
}

type internalSyntaxError struct {
	code    string
	message string
}

func newSyntaxError(code string, message string) SyntaxError {
	complainer := internalSyntaxError{
		code:code,
		message:message,
	}

	return &complainer
}

func (complainer *internalSyntaxError) Error() string {
	msg := fmt.Sprintf("Syntax Error: %s: %q", complainer.message, complainer.code)

	return msg
}

func (*internalSyntaxError) SyntaxError() {
	// Nothing here.
}

func (complainer *internalSyntaxError) Code() string {
	return complainer.code
}
