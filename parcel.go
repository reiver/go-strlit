package strlit


import (
	"bytes"
)


// Parcel is used for both containing the result of compiling
// a string literal, and for containing the original pre-compiled
// string literal code.
//
// It provides the Bytes, Runes and String methods; used to
// retrieve the value of the string literal. in []byte, []rune
// and string formats, respectively.
type Parcel interface {
	Bytes() []byte
	Runes() []rune
	String() string
}


type internalParcel struct {
	buffer bytes.Buffer
}


func newParcel() *internalParcel {
	parcel := internalParcel{}

	return &parcel
}


func (parcel *internalParcel) Bytes() []byte {
	return parcel.buffer.Bytes()
}


func (parcel *internalParcel) Runes() []rune {
	return []rune(parcel.buffer.String())
}


func (parcel *internalParcel) String() string {
	return parcel.buffer.String()
}
