// Code generated by go-swagger; DO NOT EDIT.

package transactional_email

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// EmailSendPostReader is a Reader for the EmailSendPost structure.
type EmailSendPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmailSendPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEmailSendPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewEmailSendPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewEmailSendPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewEmailSendPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewEmailSendPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewEmailSendPostMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewEmailSendPostTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewEmailSendPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEmailSendPostOK creates a EmailSendPostOK with default headers values
func NewEmailSendPostOK() *EmailSendPostOK {
	return &EmailSendPostOK{}
}

/*EmailSendPostOK handles this case with default header values.

SUCCESS
*/
type EmailSendPostOK struct {
	Payload interface{}
}

func (o *EmailSendPostOK) Error() string {
	return fmt.Sprintf("[POST /email/send][%d] emailSendPostOK  %+v", 200, o.Payload)
}

func (o *EmailSendPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailSendPostBadRequest creates a EmailSendPostBadRequest with default headers values
func NewEmailSendPostBadRequest() *EmailSendPostBadRequest {
	return &EmailSendPostBadRequest{}
}

/*EmailSendPostBadRequest handles this case with default header values.

BAD_REQUEST
*/
type EmailSendPostBadRequest struct {
	Payload interface{}
}

func (o *EmailSendPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /email/send][%d] emailSendPostBadRequest  %+v", 400, o.Payload)
}

func (o *EmailSendPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailSendPostUnauthorized creates a EmailSendPostUnauthorized with default headers values
func NewEmailSendPostUnauthorized() *EmailSendPostUnauthorized {
	return &EmailSendPostUnauthorized{}
}

/*EmailSendPostUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type EmailSendPostUnauthorized struct {
	Payload interface{}
}

func (o *EmailSendPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /email/send][%d] emailSendPostUnauthorized  %+v", 401, o.Payload)
}

func (o *EmailSendPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailSendPostForbidden creates a EmailSendPostForbidden with default headers values
func NewEmailSendPostForbidden() *EmailSendPostForbidden {
	return &EmailSendPostForbidden{}
}

/*EmailSendPostForbidden handles this case with default header values.

FORBIDDEN
*/
type EmailSendPostForbidden struct {
	Payload interface{}
}

func (o *EmailSendPostForbidden) Error() string {
	return fmt.Sprintf("[POST /email/send][%d] emailSendPostForbidden  %+v", 403, o.Payload)
}

func (o *EmailSendPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailSendPostNotFound creates a EmailSendPostNotFound with default headers values
func NewEmailSendPostNotFound() *EmailSendPostNotFound {
	return &EmailSendPostNotFound{}
}

/*EmailSendPostNotFound handles this case with default header values.

NOT_FOUND
*/
type EmailSendPostNotFound struct {
	Payload interface{}
}

func (o *EmailSendPostNotFound) Error() string {
	return fmt.Sprintf("[POST /email/send][%d] emailSendPostNotFound  %+v", 404, o.Payload)
}

func (o *EmailSendPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailSendPostMethodNotAllowed creates a EmailSendPostMethodNotAllowed with default headers values
func NewEmailSendPostMethodNotAllowed() *EmailSendPostMethodNotAllowed {
	return &EmailSendPostMethodNotAllowed{}
}

/*EmailSendPostMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type EmailSendPostMethodNotAllowed struct {
	Payload interface{}
}

func (o *EmailSendPostMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /email/send][%d] emailSendPostMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *EmailSendPostMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailSendPostTooManyRequests creates a EmailSendPostTooManyRequests with default headers values
func NewEmailSendPostTooManyRequests() *EmailSendPostTooManyRequests {
	return &EmailSendPostTooManyRequests{}
}

/*EmailSendPostTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type EmailSendPostTooManyRequests struct {
	Payload interface{}
}

func (o *EmailSendPostTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /email/send][%d] emailSendPostTooManyRequests  %+v", 429, o.Payload)
}

func (o *EmailSendPostTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailSendPostDefault creates a EmailSendPostDefault with default headers values
func NewEmailSendPostDefault(code int) *EmailSendPostDefault {
	return &EmailSendPostDefault{
		_statusCode: code,
	}
}

/*EmailSendPostDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type EmailSendPostDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the email send post default response
func (o *EmailSendPostDefault) Code() int {
	return o._statusCode
}

func (o *EmailSendPostDefault) Error() string {
	return fmt.Sprintf("[POST /email/send][%d] EmailSendPost default  %+v", o._statusCode, o.Payload)
}

func (o *EmailSendPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
