// Code generated by go-swagger; DO NOT EDIT.

package mms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// MMSSendPostReader is a Reader for the MMSSendPost structure.
type MMSSendPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MMSSendPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewMMSSendPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewMMSSendPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewMMSSendPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewMMSSendPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewMMSSendPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewMMSSendPostMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewMMSSendPostTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewMMSSendPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMMSSendPostOK creates a MMSSendPostOK with default headers values
func NewMMSSendPostOK() *MMSSendPostOK {
	return &MMSSendPostOK{}
}

/*MMSSendPostOK handles this case with default header values.

SUCCESS
*/
type MMSSendPostOK struct {
	Payload interface{}
}

func (o *MMSSendPostOK) Error() string {
	return fmt.Sprintf("[POST /mms/send][%d] mmsSendPostOK  %+v", 200, o.Payload)
}

func (o *MMSSendPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMMSSendPostBadRequest creates a MMSSendPostBadRequest with default headers values
func NewMMSSendPostBadRequest() *MMSSendPostBadRequest {
	return &MMSSendPostBadRequest{}
}

/*MMSSendPostBadRequest handles this case with default header values.

BAD_REQUEST
*/
type MMSSendPostBadRequest struct {
	Payload interface{}
}

func (o *MMSSendPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /mms/send][%d] mmsSendPostBadRequest  %+v", 400, o.Payload)
}

func (o *MMSSendPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMMSSendPostUnauthorized creates a MMSSendPostUnauthorized with default headers values
func NewMMSSendPostUnauthorized() *MMSSendPostUnauthorized {
	return &MMSSendPostUnauthorized{}
}

/*MMSSendPostUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type MMSSendPostUnauthorized struct {
	Payload interface{}
}

func (o *MMSSendPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /mms/send][%d] mmsSendPostUnauthorized  %+v", 401, o.Payload)
}

func (o *MMSSendPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMMSSendPostForbidden creates a MMSSendPostForbidden with default headers values
func NewMMSSendPostForbidden() *MMSSendPostForbidden {
	return &MMSSendPostForbidden{}
}

/*MMSSendPostForbidden handles this case with default header values.

FORBIDDEN
*/
type MMSSendPostForbidden struct {
	Payload interface{}
}

func (o *MMSSendPostForbidden) Error() string {
	return fmt.Sprintf("[POST /mms/send][%d] mmsSendPostForbidden  %+v", 403, o.Payload)
}

func (o *MMSSendPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMMSSendPostNotFound creates a MMSSendPostNotFound with default headers values
func NewMMSSendPostNotFound() *MMSSendPostNotFound {
	return &MMSSendPostNotFound{}
}

/*MMSSendPostNotFound handles this case with default header values.

NOT_FOUND
*/
type MMSSendPostNotFound struct {
	Payload interface{}
}

func (o *MMSSendPostNotFound) Error() string {
	return fmt.Sprintf("[POST /mms/send][%d] mmsSendPostNotFound  %+v", 404, o.Payload)
}

func (o *MMSSendPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMMSSendPostMethodNotAllowed creates a MMSSendPostMethodNotAllowed with default headers values
func NewMMSSendPostMethodNotAllowed() *MMSSendPostMethodNotAllowed {
	return &MMSSendPostMethodNotAllowed{}
}

/*MMSSendPostMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type MMSSendPostMethodNotAllowed struct {
	Payload interface{}
}

func (o *MMSSendPostMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /mms/send][%d] mmsSendPostMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *MMSSendPostMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMMSSendPostTooManyRequests creates a MMSSendPostTooManyRequests with default headers values
func NewMMSSendPostTooManyRequests() *MMSSendPostTooManyRequests {
	return &MMSSendPostTooManyRequests{}
}

/*MMSSendPostTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type MMSSendPostTooManyRequests struct {
	Payload interface{}
}

func (o *MMSSendPostTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /mms/send][%d] mmsSendPostTooManyRequests  %+v", 429, o.Payload)
}

func (o *MMSSendPostTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMMSSendPostDefault creates a MMSSendPostDefault with default headers values
func NewMMSSendPostDefault(code int) *MMSSendPostDefault {
	return &MMSSendPostDefault{
		_statusCode: code,
	}
}

/*MMSSendPostDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type MMSSendPostDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the Mms send post default response
func (o *MMSSendPostDefault) Code() int {
	return o._statusCode
}

func (o *MMSSendPostDefault) Error() string {
	return fmt.Sprintf("[POST /mms/send][%d] MmsSendPost default  %+v", o._statusCode, o.Payload)
}

func (o *MMSSendPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
