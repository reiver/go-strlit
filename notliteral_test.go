package strlit

import (
	"testing"
)

func TestNotLiteralAsError(t *testing.T) {

	var err error = internalNotLiteral{} // THIS IS WHAT ACTUALLY MATTERS!

	if nil == err {
		t.Error("This should never happen.")
	}
}

func TestNotLiteralAsNotLiteral(t *testing.T) {

	var complainer NotLiteral = internalNotLiteral{} // THIS IS WHAT ACTUALLY MATTERS!

	if nil == complainer {
		t.Error("This should never happen.")
	}
}
