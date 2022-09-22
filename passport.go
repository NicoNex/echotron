package echotron

import (
	"encoding/json"
	"net/url"
)

// PassportData contains information about Telegram Passport data shared with the bot by the user.
type PassportData struct {
	Credentials EncryptedCredentials       `json:"encrypted_credentials"`
	Data        []EncryptedPassportElement `json:"encrypted_passport_element"`
}

// PassportFile represents a file uploaded to Telegram Passport.
// Currently all Telegram Passport files are in JPEG format when decrypted and don't exceed 10MB.
type PassportFile struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int64  `json:"file_size"`
	FileDate     int64  `json:"file_date"`
}

// EncryptedPassportElementType is a custom type for the various possible options used as Type in EncryptedPassportElement.
type EncryptedPassportElementType string

// These are all the possible options that can be used as Type in EncryptedPassportElement.
const (
	TypePersonalDetails       EncryptedPassportElementType = "personal_details"
	TypePassport                                           = "passport"
	TypeDriverLicense                                      = "driver_license"
	TypeIdentityCard                                       = "identity_card"
	TypeInternalPassport                                   = "internal_passport"
	TypeAddress                                            = "address"
	TypeUtilityBill                                        = "utility_bill"
	TypeBankStatement                                      = "bank_statement"
	TypeRentalAgreement                                    = "rental_agreement"
	TypePassportRegistration                               = "passport_registration"
	TypeTemporaryRegistration                              = "temporary_registration"
	TypePhoneNumber                                        = "phone_number"
	TypeEmail                                              = "email"
)

// EncryptedPassportElement contains information about documents or other Telegram Passport elements shared with the bot by the user.
type EncryptedPassportElement struct {
	Type        EncryptedPassportElementType `json:"type"`
	Data        string                       `json:"data,omitempty"`
	PhoneNumber string                       `json:"phone_number,omitempty"`
	Email       string                       `json:"email,omitempty"`
	Files       *[]PassportFile              `json:"files,omitempty"`
	FrontSide   *PassportFile                `json:"front_side,omitempty"`
	ReverseSide *PassportFile                `json:"reverse_side,omitempty"`
	Selfie      *PassportFile                `json:"selfie,omitempty"`
	Translation *[]PassportFile              `json:"translation,omitempty"`
	Hash        string                       `json:"hash"`
}

// EncryptedCredentials contains data required for decrypting and authenticating EncryptedPassportElement.
// See the Telegram Passport Documentation for a complete description of the data decryption and authentication processes.
// https://core.telegram.org/passport#receiving-information
type EncryptedCredentials struct {
	Data   string `json:"data"`
	Hash   string `json:"hash"`
	Secret string `json:"secret"`
}

// PassportElementErrorSource is a custom type for the various possible options used as Source in PassportElementSource.
type PassportElementErrorSource string

// These are all the possible options that can be used as Source in PassportElementSource.
const (
	SourceData             PassportElementErrorSource = "data"
	SourceFrontSide                                   = "front_side"
	SourceReverseSide                                 = "reverse_side"
	SourceSelfie                                      = "selfie"
	SourceFile                                        = "file"
	SourceFiles                                       = "files"
	SourceTranslationFile                             = "translation_file"
	SourceTranslationFiles                            = "translation_files"
	SourceUnspecified                                 = "unspecified"
)

// PassportElementError is an interface for the various PassportElementError types.
type PassportElementError interface {
	ImplementsPassportElementError()
}

// PassportElementErrorDataField represents an issue in one of the data fields that was provided by the user.
// The error is considered resolved when the field's value changes.
// Source MUST BE SourceData.
// Type MUST BE one of TypePersonalDetails, TypePassport, TypeDriverLicense, TypeIdentityCard, TypeInternalPassport and TypeAddress.
type PassportElementErrorDataField struct {
	Source    PassportElementErrorSource   `json:"source"`
	Type      EncryptedPassportElementType `json:"type"`
	FieldName string                       `json:"field_name"`
	DataHash  string                       `json:"data_hash"`
	Message   string                       `json:"message"`
}

// ImplementsPassportElementError us a dummy method which exists to implement the interface PassportElementError.
func (p PassportElementErrorDataField) ImplementsPassportElementError() {}

// PassportElementErrorFrontSide represents an issue with the front side of a document.
// The error is considered resolved when the file with the front side of the document changes.
// Source MUST BE SourceFrontSide.
// Type MUST BE one of TypeDriverLicense and TypeIdentityCard.
type PassportElementErrorFrontSide struct {
	Source   PassportElementErrorSource   `json:"source"`
	Type     EncryptedPassportElementType `json:"type"`
	FileHash string                       `json:"file_hash"`
	Message  string                       `json:"message"`
}

// ImplementsPassportElementError us a dummy method which exists to implement the interface PassportElementError.
func (p PassportElementErrorFrontSide) ImplementsPassportElementError() {}

// PassportElementErrorReverseSide represents an issue with the reverse side of a document.
// The error is considered resolved when the file with the reverse side of the document changes.
// Source MUST BE SourceReverseSide.
// Type MUST BE one of TypeDriverLicense and TypeIdentityCard.
type PassportElementErrorReverseSide struct {
	Source   PassportElementErrorSource   `json:"source"`
	Type     EncryptedPassportElementType `json:"type"`
	FileHash string                       `json:"file_hash"`
	Message  string                       `json:"message"`
}

// ImplementsPassportElementError us a dummy method which exists to implement the interface PassportElementError.
func (p PassportElementErrorReverseSide) ImplementsPassportElementError() {}

// PassportElementErrorSelfie represents an issue with the selfie with a document.
// The error is considered resolved when the file with the selfie changes.
// Source MUST BE SourceSelfie.
// Type MUST BE one of TypePassport, TypeDriverLicense, TypeIdentityCard and TypeIdentityPassport.
type PassportElementErrorSelfie struct {
	Source   PassportElementErrorSource   `json:"source"`
	Type     EncryptedPassportElementType `json:"type"`
	FileHash string                       `json:"file_hash"`
	Message  string                       `json:"message"`
}

// ImplementsPassportElementError us a dummy method which exists to implement the interface PassportElementError.
func (p PassportElementErrorSelfie) ImplementsPassportElementError() {}

// PassportElementErrorFile represents an issue with the document scan.
// The error is considered resolved when the file with the document scan changes.
// Source MUST BE SourceFile.
// Type MUST BE one of TypePassport, TypeDriverLicense, TypeIdentityCard and TypeIdentityPassport.
type PassportElementErrorFile struct {
	Source   PassportElementErrorSource   `json:"source"`
	Type     EncryptedPassportElementType `json:"type"`
	FileHash string                       `json:"file_hash"`
	Message  string                       `json:"message"`
}

// ImplementsPassportElementError us a dummy method which exists to implement the interface PassportElementError.
func (p PassportElementErrorFile) ImplementsPassportElementError() {}

// PassportElementErrorFiles represents an issue with a list of scans.
// The error is considered resolved when the list of files containing the scans changes.
// Source MUST BE SourceFiles.
// Type MUST BE one of TypeUtilityBill, TypeBankStatement, TypeRentalAgreement, TypePassportRegistration and TypeTemporaryRegistration.
type PassportElementErrorFiles struct {
	Source     PassportElementErrorSource   `json:"source"`
	Type       EncryptedPassportElementType `json:"type"`
	Message    string                       `json:"message"`
	FileHashes []string                     `json:"file_hashes"`
}

// ImplementsPassportElementError us a dummy method which exists to implement the interface PassportElementError.
func (p PassportElementErrorFiles) ImplementsPassportElementError() {}

// PassportElementErrorTranslationFile represents an issue with one of the files that constitute the translation of the document.
// The error is considered resolved when the file changes.
// Source MUST BE SourceTranslationFile.
// Type MUST BE one of TypePassport, TypeDriverLicense, TypeIdentityCard, TypeInternalPassport, TypeUtilityBill, TypeBankStatement,
// TypeRentalAgreement, TypePassportRegistration and TypeTemporaryRegistration.
type PassportElementErrorTranslationFile struct {
	Source   PassportElementErrorSource   `json:"source"`
	Type     EncryptedPassportElementType `json:"type"`
	FileHash string                       `json:"file_hash"`
	Message  string                       `json:"message"`
}

// ImplementsPassportElementError us a dummy method which exists to implement the interface PassportElementError.
func (p PassportElementErrorTranslationFile) ImplementsPassportElementError() {}

// PassportElementErrorTranslationFiles represents an issue with the translated version of a document.
// The error is considered resolved when a file with the document translation changes.
// Source MUST BE SourceTranslationFiles.
// Type MUST BE one of TypePassport, TypeDriverLicense, TypeIdentityCard, TypeInternalPassport, TypeUtilityBill, TypeBankStatement,
// TypeRentalAgreement, TypePassportRegistration and TypeTemporaryRegistration.
type PassportElementErrorTranslationFiles struct {
	Source     PassportElementErrorSource   `json:"source"`
	Type       EncryptedPassportElementType `json:"type"`
	Message    string                       `json:"message"`
	FileHashes []string                     `json:"file_hashes"`
}

// ImplementsPassportElementError us a dummy method which exists to implement the interface PassportElementError.
func (p PassportElementErrorTranslationFiles) ImplementsPassportElementError() {}

// PassportElementErrorUnspecified represents an issue in an unspecified place.
// The error is considered resolved when new data is added.
type PassportElementErrorUnspecified struct {
	Source      PassportElementErrorSource   `json:"source"`
	Type        EncryptedPassportElementType `json:"type"`
	ElementHash string                       `json:"element_hash"`
	Message     string                       `json:"message"`
}

// ImplementsPassportElementError us a dummy method which exists to implement the interface PassportElementError.
func (p PassportElementErrorUnspecified) ImplementsPassportElementError() {}

// SetPassportDataErrors Informs a user that some of the Telegram Passport elements they provided contains errors.
// The user will not be able to re-submit their Passport to you until the errors are fixed.
// The contents of the field for which you returned the error must change.
func (a API) SetPassportDataErrors(userID int64, errors []PassportElementError) (res APIResponseBool, err error) {
	var vals = make(url.Values)

	errorsArr, err := json.Marshal(errors)
	if err != nil {
		return res, err
	}

	vals.Set("user_id", itoa(userID))
	vals.Set("errors", string(errorsArr))
	return get[APIResponseBool](a.base, "setPassportDataErrors", vals)
}
