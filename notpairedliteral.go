package strlit

import (
	"fmt"
)

type NotPairedLiteral interface {
	error
	NotPairedLiteral() rune
}

func errNotPairedLiteral(r rune) error {
	var e NotPairedLiteral = &internalNotPairedLiteral{
		r:r,
	}

	return e
}

type internalNotPairedLiteral struct {
	r rune
}

func (receiver internalNotPairedLiteral) Error() string {
	return fmt.Sprintf("strlit: Not Paired Literal: paired literal cannot being with %q (%d)", receiver.r, receiver.r)
}

func (receiver internalNotPairedLiteral) NotPairedLiteral() rune {
	return receiver.r
}
