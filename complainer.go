package strlit


import (
	"fmt"
)


// SyntaxErrorComplainer is an error that is used to represent a syntax error in a
// string literal.
//
// SyntaxErrorComplainer allows one to respond to errors returned from strlit.Compile()
// in a more precise way. For example:
//
//	compiled, err := strlit.Compile(runeReader)
//	if nil != err {
//		switch err.(type) {
//		case strlit.SyntaxErrorComplainer:
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
type SyntaxErrorComplainer interface {
	error

	// SyntaxErrorComplainer is used for typing. Calling it will do nothing.
	SyntaxErrorComplainer()

	// Code returns "code" from the string literal that caused the error.
	Code() string
}


type internalSyntaxErrorComplainer struct {
	code    string
	message string
}


func newSyntaxErrorComplainer(code string, message string) SyntaxErrorComplainer {
	complainer := internalSyntaxErrorComplainer{
		code:code,
		message:message,
	}

	return &complainer
}


func (complainer *internalSyntaxErrorComplainer) Error() string {
	msg := fmt.Sprintf("Syntax Error: %s: %q", complainer.message, complainer.code)

	return msg
}


func (*internalSyntaxErrorComplainer) SyntaxErrorComplainer() {
	// Nothing here.
}


func (complainer *internalSyntaxErrorComplainer) Code() string {
	return complainer.code
}
