package strlit

import (
	"fmt"
)

type NotLiteral interface {
	error
	NotLiteral() rune
}

func errNotLiteral(r rune) error {
	var e NotLiteral = &internalNotLiteral{
		r:r,
	}

	return e
}

type internalNotLiteral struct {
	r rune
}

func (receiver internalNotLiteral) Error() string {
	return fmt.Sprintf("strlit: Not Literal: literal cannot being with %q (%d)", receiver.r, receiver.r)
}

func (receiver internalNotLiteral) NotLiteral() rune {
	return receiver.r
}
