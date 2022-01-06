package echotron

import "testing"

var a APIError

func TestErrorCode(_ *t.Testing) {
	a.ErrorCode()
}

func TestDescription(_ *t.Testing) {
	a.Description()
}

func TestError(_ *t.Testing) {
	a.Error()
}
