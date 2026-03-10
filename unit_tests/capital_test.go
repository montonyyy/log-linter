package logcheck

import (
	"log-linter/logcheck"
	"testing"
)

func TestHasCapitalLetter(t *testing.T) {
	if !logcheck.HasCapitalLetter("Hello") {
		t.Error("must be true")
	}

	if logcheck.HasCapitalLetter("hello") {
		t.Error("must be false")
	}
	if logcheck.HasCapitalLetter("123") {
		t.Error("must be false")
	}
	if logcheck.HasCapitalLetter("") {
		t.Error("must be false")
	}
}
