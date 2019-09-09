package strlit

import (
	"testing"
)

func TestNotBareLiteralAsError(t *testing.T) {

	var err error = internalNotBareLiteral{} // THIS IS WHAT ACTUALLY MATTERS!

	if nil == err {
		t.Error("This should never happen.")
	}
}

func TestNotBareLiteralAsNotLiteral(t *testing.T) {

	var complainer NotLiteral = internalNotBareLiteral{} // THIS IS WHAT ACTUALLY MATTERS!

	if nil == complainer {
		t.Error("This should never happen.")
	}
}
