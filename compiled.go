package strlit


import (
	"bytes"
)


// Compiled is the result of compiling a string literal.
// It provides the Bytes, Runes and String methods; used to
// retrieve the value of the string literal. in []byte, []rune
// and string formats, respectively.
type Compiled interface {
	Bytes() []byte
	Runes() []rune
	String() string
}


type internalCompiled struct {
	buffer bytes.Buffer
}


func newCompiled() *internalCompiled {
	compiled := internalCompiled{}

	return &compiled
}


func (compiled *internalCompiled) Bytes() []byte {
	return compiled.buffer.Bytes()
}


func (compiled *internalCompiled) Runes() []rune {
	return []rune(compiled.buffer.String())
}


func (compiled *internalCompiled) String() string {
	return compiled.buffer.String()
}
