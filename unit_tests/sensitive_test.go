package logcheck

import (
	"log-linter/logcheck"
	"testing"
)

func TestHasSensitive(t *testing.T) {
	if !logcheck.HasSensitive("password") {
		t.Error("must be true")
	}
	if !logcheck.HasSensitive("token") {
		t.Error("must be true")
	}

	if logcheck.HasSensitive("hello world") {
		t.Error("must be false")
	}
	if logcheck.HasSensitive("123") {
		t.Error("must be false")
	}
}
