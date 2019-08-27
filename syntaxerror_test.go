package strlit

import (
	"testing"
)

func TestInternalSyntaxErrorAsError(t *testing.T) {

	var err error = internalSyntaxError{} // THIS IS WHAT ACTUALLY MATTERS!

	if nil == err {
		t.Error("This should never happen.")
	}
}
func TestInternalSyntaxErrorAsSyntaxError(t *testing.T) {

	var complainer SyntaxError = internalSyntaxError{} // THIS IS WHAT ACTUALLY MATTERS!

	if nil == complainer {
		t.Error("This should never happen.")
	}
}
