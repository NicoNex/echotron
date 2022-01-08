package echotron

import "testing"

var a APIError

func TestErrorCode(_ *testing.T) {
	a.ErrorCode()
}

func TestDescription(_ *testing.T) {
	a.Description()
}

func TestError(_ *testing.T) {
	_ = a.Error()
}
