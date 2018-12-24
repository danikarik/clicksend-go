// Code generated by go-swagger; DO NOT EDIT.

package voice

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// VoiceReceiptsGetReader is a Reader for the VoiceReceiptsGet structure.
type VoiceReceiptsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VoiceReceiptsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewVoiceReceiptsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewVoiceReceiptsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewVoiceReceiptsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewVoiceReceiptsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewVoiceReceiptsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewVoiceReceiptsGetMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewVoiceReceiptsGetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewVoiceReceiptsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVoiceReceiptsGetOK creates a VoiceReceiptsGetOK with default headers values
func NewVoiceReceiptsGetOK() *VoiceReceiptsGetOK {
	return &VoiceReceiptsGetOK{}
}

/*VoiceReceiptsGetOK handles this case with default header values.

SUCCESS
*/
type VoiceReceiptsGetOK struct {
	Payload interface{}
}

func (o *VoiceReceiptsGetOK) Error() string {
	return fmt.Sprintf("[GET /voice/receipts][%d] voiceReceiptsGetOK  %+v", 200, o.Payload)
}

func (o *VoiceReceiptsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceReceiptsGetBadRequest creates a VoiceReceiptsGetBadRequest with default headers values
func NewVoiceReceiptsGetBadRequest() *VoiceReceiptsGetBadRequest {
	return &VoiceReceiptsGetBadRequest{}
}

/*VoiceReceiptsGetBadRequest handles this case with default header values.

BAD_REQUEST
*/
type VoiceReceiptsGetBadRequest struct {
	Payload interface{}
}

func (o *VoiceReceiptsGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /voice/receipts][%d] voiceReceiptsGetBadRequest  %+v", 400, o.Payload)
}

func (o *VoiceReceiptsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceReceiptsGetUnauthorized creates a VoiceReceiptsGetUnauthorized with default headers values
func NewVoiceReceiptsGetUnauthorized() *VoiceReceiptsGetUnauthorized {
	return &VoiceReceiptsGetUnauthorized{}
}

/*VoiceReceiptsGetUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type VoiceReceiptsGetUnauthorized struct {
	Payload interface{}
}

func (o *VoiceReceiptsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /voice/receipts][%d] voiceReceiptsGetUnauthorized  %+v", 401, o.Payload)
}

func (o *VoiceReceiptsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceReceiptsGetForbidden creates a VoiceReceiptsGetForbidden with default headers values
func NewVoiceReceiptsGetForbidden() *VoiceReceiptsGetForbidden {
	return &VoiceReceiptsGetForbidden{}
}

/*VoiceReceiptsGetForbidden handles this case with default header values.

FORBIDDEN
*/
type VoiceReceiptsGetForbidden struct {
	Payload interface{}
}

func (o *VoiceReceiptsGetForbidden) Error() string {
	return fmt.Sprintf("[GET /voice/receipts][%d] voiceReceiptsGetForbidden  %+v", 403, o.Payload)
}

func (o *VoiceReceiptsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceReceiptsGetNotFound creates a VoiceReceiptsGetNotFound with default headers values
func NewVoiceReceiptsGetNotFound() *VoiceReceiptsGetNotFound {
	return &VoiceReceiptsGetNotFound{}
}

/*VoiceReceiptsGetNotFound handles this case with default header values.

NOT_FOUND
*/
type VoiceReceiptsGetNotFound struct {
	Payload interface{}
}

func (o *VoiceReceiptsGetNotFound) Error() string {
	return fmt.Sprintf("[GET /voice/receipts][%d] voiceReceiptsGetNotFound  %+v", 404, o.Payload)
}

func (o *VoiceReceiptsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceReceiptsGetMethodNotAllowed creates a VoiceReceiptsGetMethodNotAllowed with default headers values
func NewVoiceReceiptsGetMethodNotAllowed() *VoiceReceiptsGetMethodNotAllowed {
	return &VoiceReceiptsGetMethodNotAllowed{}
}

/*VoiceReceiptsGetMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type VoiceReceiptsGetMethodNotAllowed struct {
	Payload interface{}
}

func (o *VoiceReceiptsGetMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /voice/receipts][%d] voiceReceiptsGetMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *VoiceReceiptsGetMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceReceiptsGetTooManyRequests creates a VoiceReceiptsGetTooManyRequests with default headers values
func NewVoiceReceiptsGetTooManyRequests() *VoiceReceiptsGetTooManyRequests {
	return &VoiceReceiptsGetTooManyRequests{}
}

/*VoiceReceiptsGetTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type VoiceReceiptsGetTooManyRequests struct {
	Payload interface{}
}

func (o *VoiceReceiptsGetTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /voice/receipts][%d] voiceReceiptsGetTooManyRequests  %+v", 429, o.Payload)
}

func (o *VoiceReceiptsGetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceReceiptsGetDefault creates a VoiceReceiptsGetDefault with default headers values
func NewVoiceReceiptsGetDefault(code int) *VoiceReceiptsGetDefault {
	return &VoiceReceiptsGetDefault{
		_statusCode: code,
	}
}

/*VoiceReceiptsGetDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type VoiceReceiptsGetDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the voice receipts get default response
func (o *VoiceReceiptsGetDefault) Code() int {
	return o._statusCode
}

func (o *VoiceReceiptsGetDefault) Error() string {
	return fmt.Sprintf("[GET /voice/receipts][%d] VoiceReceiptsGet default  %+v", o._statusCode, o.Payload)
}

func (o *VoiceReceiptsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
