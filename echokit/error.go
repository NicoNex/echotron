package main

import (
	"fmt"

	"github.com/NicoNex/echotron/v3"
)

var ErrInvalidAPIResponse = errors.New("invalid API response object")

// APIError is an error generated by a Telegram resonse
type APIError struct {
	code int
	desc string
}

// NewAPIError creates a new NewAPIError with given errorCode and description
func NewAPIError(errorCode int, description string) APIError {
	return APIError{code: errorCode, desc: description}
}

// Code returns the error-code status of the error
func (a APIError) Code() int {
	return a.code
}

// Desc return the description of the error
func (a APIError) Desc() string {
	return a.desc
}

// Error returns a complete error description
// thanks to this method ResponseError can be used as an error interface
func (a APIError) Error() string {
	return fmt.Sprint("api error, code ", a.code, ", description ", a.desc)
}

// CheckResponseBase transform an error and an API response object into a *ResponseError (error interface)
// The only valid API response object types are the followings, if a different
// one is provided the error will be of type ErrInvalidAPIResponse.
// - APIResponseAdministrators
// - APIResponseBase
// - APIResponseBool
// - APIResponseChat
// - APIResponseChatMember
// - APIResponseCommands
// - APIResponseFile
// - APIResponseGameHighScore
// - APIResponseInteger
// - APIResponseInviteLink
// - APIResponseMessage
// - APIResponseMessageArray
// - APIResponseMessageID
// - APIResponsePoll
// - APIResponseStickerSet
// - APIResponseString
// - APIResponseUpdate
// - APIResponseUser
// - APIResponseUserProfile
// - APIResponseWebhook
func Check(APIResponse interface{}, err error) error {
	if err != nil {
		return err
	}

	var base echotron.APIResponseBase
	switch res := APIResponse.(type) {
	case APIResponseAdministrators:
		base = res.APIResponseBase
	case APIResponseBase:
		base = res.APIResponseBase
	case APIResponseBool:
		base = res.APIResponseBase
	case APIResponseChat:
		base = res.APIResponseBase
	case APIResponseChatMember:
		base = res.APIResponseBase
	case APIResponseCommands:
		base = res.APIResponseBase
	case APIResponseFile:
		base = res.APIResponseBase
	case APIResponseGameHighScore:
		base = res.APIResponseBase
	case APIResponseInteger:
		base = res.APIResponseBase
	case APIResponseInviteLink:
		base = res.APIResponseBase
	case APIResponseMessage:
		base = res.APIResponseBase
	case APIResponseMessageArray:
		base = res.APIResponseBase
	case APIResponseMessageID:
		base = res.APIResponseBase
	case APIResponsePoll:
		base = res.APIResponseBase
	case APIResponseStickerSet:
		base = res.APIResponseBase
	case APIResponseString:
		base = res.APIResponseBase
	case APIResponseUpdate:
		base = res.APIResponseBase
	case APIResponseUser:
		base = res.APIResponseBase
	case APIResponseUserProfile:
		base = res.APIResponseBase
	case APIResponseWebhook:
		base = res.APIResponseBase
	default:
		return ErrInvalidAPIResponse
	}

	if !base.Ok {
		return NewAPIError(base.ErrorCode, base.Description)
	}

	return nil
}
