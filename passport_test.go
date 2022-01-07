package echotron

import "testing"

func TestPassportElementErrorDataField(_ *testing.T) {
	p := PassportElementErrorDataField{}
	p.ImplementsPassportElementError()
}

func TestPassportElementErrorFrontSide(_ *testing.T) {
	p := PassportElementErrorFrontSide{}
	p.ImplementsPassportElementError()
}

func TestPassportElementErrorReverseSide(_ *testing.T) {
	p := PassportElementErrorReverseSide{}
	p.ImplementsPassportElementError()
}

func TestPassportElementErrorSelfie(_ *testing.T) {
	p := PassportElementErrorSelfie{}
	p.ImplementsPassportElementError()
}

func TestPassportElementErrorFile(_ *testing.T) {
	p := PassportElementErrorFile{}
	p.ImplementsPassportElementError()
}

func TestPassportElementErrorFiles(_ *testing.T) {
	p := PassportElementErrorFiles{}
	p.ImplementsPassportElementError()
}

func TestPassportElementErrorTranslationFile(_ *testing.T) {
	p := PassportElementErrorTranslationFile{}
	p.ImplementsPassportElementError()
}

func TestPassportElementErrorTranslationFiles(_ *testing.T) {
	p := PassportElementErrorTranslationFiles{}
	p.ImplementsPassportElementError()
}

func TestPassportElementErrorUnspecified(_ *testing.T) {
	p := PassportElementErrorUnspecified{}
	p.ImplementsPassportElementError()
}
