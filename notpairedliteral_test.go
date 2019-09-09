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

func TestNotPairedLiteralAsNotPairedLiteral(t *testing.T) {

	var complainer NotPairedLiteral = internalNotPairedLiteral{} // THIS IS WHAT ACTUALLY MATTERS!

	if nil == complainer {
		t.Error("This should never happen.")
	}
}
