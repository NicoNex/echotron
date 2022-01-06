package echotron

func TestPassportElementErrorDataField(_ *t.Testing) {
	p := PassportElementErrorDataField
	p.ImplementsPassportElementError()
}

func TestPassportElementErrorFrontSide(_ *t.Testing) {
	p := PassportElementErrorFrontSide
	p.ImplementsPassportElementError()
}

func TestPassportElementErrorReverseSide(_ *t.Testing) {
	p := PassportElementErrorReverseSide
	p.ImplementsPassportElementError()
}

func TestPassportElementErrorSelfie(_ *t.Testing) {
	p := PassportElementErrorSelfie
	p.ImplementsPassportElementError()
}

func TestPassportElementErrorFile(_ *t.Testing) {
	p := PassportElementErrorFile
	p.ImplementsPassportElementError()
}

func TestPassportElementErrorFiles(_ *t.Testing) {
	p := PassportElementErrorFiles
	p.ImplementsPassportElementError()
}

func TestPassportElementErrorTranslationFile(_ *t.Testing) {
	p := PassportElementErrorTranslationFile
	p.ImplementsPassportElementError()
}

func TestPassportElementErrorTranslationFiles(_ *t.Testing) {
	p := PassportElementErrorTranslationFiles
	p.ImplementsPassportElementError()
}

func TestPassportElementErrorUnspecified(_ *t.Testing) {
	p := PassportElementErrorUnspecified
	p.ImplementsPassportElementError()
}
