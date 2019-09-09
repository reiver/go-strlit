package strlit

import (
	"fmt"
)

func errNotBareLiteral(r rune) error {
	var e NotLiteral = &internalNotBareLiteral{
		r:r,
	}

	return e
}

type internalNotBareLiteral struct {
	r rune
}

func (receiver internalNotBareLiteral) Error() string {
	return fmt.Sprintf("strlit: Not Bare Literal: bare literal cannot being with %q (%d)", receiver.r, receiver.r)
}

func (receiver internalNotBareLiteral) NotLiteral() {
	// Nothing here.
}
