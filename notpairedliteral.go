package strlit

import (
	"fmt"
)

func errNotPairedLiteral(r rune) error {
	var e NotLiteral = &internalNotPairedLiteral{
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

func (receiver internalNotPairedLiteral) NotLiteral() {
	// Nothing here.
}
