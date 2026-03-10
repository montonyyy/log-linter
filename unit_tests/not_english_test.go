package logcheck

import (
	"log-linter/logcheck"
	"testing"
)

func TestHasNonEnglish(t *testing.T) {
	if !logcheck.HasNonEnglish("привет") {
		t.Error("must be true")
	}

	if logcheck.HasNonEnglish("hello") {
		t.Error("must be false")
	}
	if logcheck.HasNonEnglish("123") {
		t.Error("must be false")
	}
}
