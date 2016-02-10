package strlit


import (
	"bytes"
)


// Code is the original code of a string literal, pre-compiling.
// It provides the Bytes, Runes and String methods; used to
// retrieve the original code of the string literal. in
// []byte, []rune and string formats, respectively.
type Code interface {
	Code()

	Bytes() []byte
	Runes() []rune
	String() string
}


type internalCode struct {
	buffer bytes.Buffer
}


func newCode() *internalCode {
	code := internalCode{}

	return &code
}


func (*internalCode) Code() {
	// Nothing here.
}


func (code *internalCode) Bytes() []byte {
	return code.buffer.Bytes()
}


func (code *internalCode) Runes() []rune {
	return []rune(code.buffer.String())
}


func (code *internalCode) String() string {
	return code.buffer.String()
}
