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

// VoiceCancelByMessageIDPutReader is a Reader for the VoiceCancelByMessageIDPut structure.
type VoiceCancelByMessageIDPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VoiceCancelByMessageIDPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewVoiceCancelByMessageIDPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewVoiceCancelByMessageIDPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewVoiceCancelByMessageIDPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewVoiceCancelByMessageIDPutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewVoiceCancelByMessageIDPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewVoiceCancelByMessageIDPutMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewVoiceCancelByMessageIDPutTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewVoiceCancelByMessageIDPutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVoiceCancelByMessageIDPutOK creates a VoiceCancelByMessageIDPutOK with default headers values
func NewVoiceCancelByMessageIDPutOK() *VoiceCancelByMessageIDPutOK {
	return &VoiceCancelByMessageIDPutOK{}
}

/*VoiceCancelByMessageIDPutOK handles this case with default header values.

SUCCESS
*/
type VoiceCancelByMessageIDPutOK struct {
	Payload string
}

func (o *VoiceCancelByMessageIDPutOK) Error() string {
	return fmt.Sprintf("[PUT /voice/{message_id}/cancel][%d] voiceCancelByMessageIdPutOK  %+v", 200, o.Payload)
}

func (o *VoiceCancelByMessageIDPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceCancelByMessageIDPutBadRequest creates a VoiceCancelByMessageIDPutBadRequest with default headers values
func NewVoiceCancelByMessageIDPutBadRequest() *VoiceCancelByMessageIDPutBadRequest {
	return &VoiceCancelByMessageIDPutBadRequest{}
}

/*VoiceCancelByMessageIDPutBadRequest handles this case with default header values.

BAD_REQUEST
*/
type VoiceCancelByMessageIDPutBadRequest struct {
	Payload string
}

func (o *VoiceCancelByMessageIDPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /voice/{message_id}/cancel][%d] voiceCancelByMessageIdPutBadRequest  %+v", 400, o.Payload)
}

func (o *VoiceCancelByMessageIDPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceCancelByMessageIDPutUnauthorized creates a VoiceCancelByMessageIDPutUnauthorized with default headers values
func NewVoiceCancelByMessageIDPutUnauthorized() *VoiceCancelByMessageIDPutUnauthorized {
	return &VoiceCancelByMessageIDPutUnauthorized{}
}

/*VoiceCancelByMessageIDPutUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type VoiceCancelByMessageIDPutUnauthorized struct {
	Payload string
}

func (o *VoiceCancelByMessageIDPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /voice/{message_id}/cancel][%d] voiceCancelByMessageIdPutUnauthorized  %+v", 401, o.Payload)
}

func (o *VoiceCancelByMessageIDPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceCancelByMessageIDPutForbidden creates a VoiceCancelByMessageIDPutForbidden with default headers values
func NewVoiceCancelByMessageIDPutForbidden() *VoiceCancelByMessageIDPutForbidden {
	return &VoiceCancelByMessageIDPutForbidden{}
}

/*VoiceCancelByMessageIDPutForbidden handles this case with default header values.

FORBIDDEN
*/
type VoiceCancelByMessageIDPutForbidden struct {
	Payload string
}

func (o *VoiceCancelByMessageIDPutForbidden) Error() string {
	return fmt.Sprintf("[PUT /voice/{message_id}/cancel][%d] voiceCancelByMessageIdPutForbidden  %+v", 403, o.Payload)
}

func (o *VoiceCancelByMessageIDPutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceCancelByMessageIDPutNotFound creates a VoiceCancelByMessageIDPutNotFound with default headers values
func NewVoiceCancelByMessageIDPutNotFound() *VoiceCancelByMessageIDPutNotFound {
	return &VoiceCancelByMessageIDPutNotFound{}
}

/*VoiceCancelByMessageIDPutNotFound handles this case with default header values.

NOT_FOUND
*/
type VoiceCancelByMessageIDPutNotFound struct {
	Payload string
}

func (o *VoiceCancelByMessageIDPutNotFound) Error() string {
	return fmt.Sprintf("[PUT /voice/{message_id}/cancel][%d] voiceCancelByMessageIdPutNotFound  %+v", 404, o.Payload)
}

func (o *VoiceCancelByMessageIDPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceCancelByMessageIDPutMethodNotAllowed creates a VoiceCancelByMessageIDPutMethodNotAllowed with default headers values
func NewVoiceCancelByMessageIDPutMethodNotAllowed() *VoiceCancelByMessageIDPutMethodNotAllowed {
	return &VoiceCancelByMessageIDPutMethodNotAllowed{}
}

/*VoiceCancelByMessageIDPutMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type VoiceCancelByMessageIDPutMethodNotAllowed struct {
	Payload string
}

func (o *VoiceCancelByMessageIDPutMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /voice/{message_id}/cancel][%d] voiceCancelByMessageIdPutMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *VoiceCancelByMessageIDPutMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceCancelByMessageIDPutTooManyRequests creates a VoiceCancelByMessageIDPutTooManyRequests with default headers values
func NewVoiceCancelByMessageIDPutTooManyRequests() *VoiceCancelByMessageIDPutTooManyRequests {
	return &VoiceCancelByMessageIDPutTooManyRequests{}
}

/*VoiceCancelByMessageIDPutTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type VoiceCancelByMessageIDPutTooManyRequests struct {
	Payload string
}

func (o *VoiceCancelByMessageIDPutTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /voice/{message_id}/cancel][%d] voiceCancelByMessageIdPutTooManyRequests  %+v", 429, o.Payload)
}

func (o *VoiceCancelByMessageIDPutTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceCancelByMessageIDPutDefault creates a VoiceCancelByMessageIDPutDefault with default headers values
func NewVoiceCancelByMessageIDPutDefault(code int) *VoiceCancelByMessageIDPutDefault {
	return &VoiceCancelByMessageIDPutDefault{
		_statusCode: code,
	}
}

/*VoiceCancelByMessageIDPutDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type VoiceCancelByMessageIDPutDefault struct {
	_statusCode int

	Payload string
}

// Code gets the status code for the voice cancel by message Id put default response
func (o *VoiceCancelByMessageIDPutDefault) Code() int {
	return o._statusCode
}

func (o *VoiceCancelByMessageIDPutDefault) Error() string {
	return fmt.Sprintf("[PUT /voice/{message_id}/cancel][%d] VoiceCancelByMessageIdPut default  %+v", o._statusCode, o.Payload)
}

func (o *VoiceCancelByMessageIDPutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
