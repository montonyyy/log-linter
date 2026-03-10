package logcheck

import (
	"log-linter/logcheck"
	"testing"
)

func TestHasSymbol(t *testing.T) {
	if !logcheck.HasSymbol("hello!") {
		t.Error("must be true")
	}
	if !logcheck.HasSymbol("what?") {
		t.Error("must be true")
	}

	if logcheck.HasSymbol("hello") {
		t.Error("must be false")
	}
	if logcheck.HasSymbol("123") {
		t.Error("must be false")
	}
}
