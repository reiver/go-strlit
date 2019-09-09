package strlit

import (
	"testing"
)

func TestNotPairedLiteralAsError(t *testing.T) {

	var err error = internalNotPairedLiteral{} // THIS IS WHAT ACTUALLY MATTERS!

	if nil == err {
		t.Error("This should never happen.")
	}
}

func TestNotPairedLiteralAsNotLiteral(t *testing.T) {

	var complainer NotLiteral = internalNotPairedLiteral{} // THIS IS WHAT ACTUALLY MATTERS!

	if nil == complainer {
		t.Error("This should never happen.")
	}
}
