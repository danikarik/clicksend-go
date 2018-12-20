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

// EmailPricePostReader is a Reader for the EmailPricePost structure.
type EmailPricePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmailPricePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEmailPricePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewEmailPricePostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewEmailPricePostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewEmailPricePostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewEmailPricePostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewEmailPricePostMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewEmailPricePostTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewEmailPricePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEmailPricePostOK creates a EmailPricePostOK with default headers values
func NewEmailPricePostOK() *EmailPricePostOK {
	return &EmailPricePostOK{}
}

/*EmailPricePostOK handles this case with default header values.

SUCCESS
*/
type EmailPricePostOK struct {
	Payload string
}

func (o *EmailPricePostOK) Error() string {
	return fmt.Sprintf("[POST /email/price][%d] emailPricePostOK  %+v", 200, o.Payload)
}

func (o *EmailPricePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailPricePostBadRequest creates a EmailPricePostBadRequest with default headers values
func NewEmailPricePostBadRequest() *EmailPricePostBadRequest {
	return &EmailPricePostBadRequest{}
}

/*EmailPricePostBadRequest handles this case with default header values.

BAD_REQUEST
*/
type EmailPricePostBadRequest struct {
	Payload string
}

func (o *EmailPricePostBadRequest) Error() string {
	return fmt.Sprintf("[POST /email/price][%d] emailPricePostBadRequest  %+v", 400, o.Payload)
}

func (o *EmailPricePostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailPricePostUnauthorized creates a EmailPricePostUnauthorized with default headers values
func NewEmailPricePostUnauthorized() *EmailPricePostUnauthorized {
	return &EmailPricePostUnauthorized{}
}

/*EmailPricePostUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type EmailPricePostUnauthorized struct {
	Payload string
}

func (o *EmailPricePostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /email/price][%d] emailPricePostUnauthorized  %+v", 401, o.Payload)
}

func (o *EmailPricePostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailPricePostForbidden creates a EmailPricePostForbidden with default headers values
func NewEmailPricePostForbidden() *EmailPricePostForbidden {
	return &EmailPricePostForbidden{}
}

/*EmailPricePostForbidden handles this case with default header values.

FORBIDDEN
*/
type EmailPricePostForbidden struct {
	Payload string
}

func (o *EmailPricePostForbidden) Error() string {
	return fmt.Sprintf("[POST /email/price][%d] emailPricePostForbidden  %+v", 403, o.Payload)
}

func (o *EmailPricePostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailPricePostNotFound creates a EmailPricePostNotFound with default headers values
func NewEmailPricePostNotFound() *EmailPricePostNotFound {
	return &EmailPricePostNotFound{}
}

/*EmailPricePostNotFound handles this case with default header values.

NOT_FOUND
*/
type EmailPricePostNotFound struct {
	Payload string
}

func (o *EmailPricePostNotFound) Error() string {
	return fmt.Sprintf("[POST /email/price][%d] emailPricePostNotFound  %+v", 404, o.Payload)
}

func (o *EmailPricePostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailPricePostMethodNotAllowed creates a EmailPricePostMethodNotAllowed with default headers values
func NewEmailPricePostMethodNotAllowed() *EmailPricePostMethodNotAllowed {
	return &EmailPricePostMethodNotAllowed{}
}

/*EmailPricePostMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type EmailPricePostMethodNotAllowed struct {
	Payload string
}

func (o *EmailPricePostMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /email/price][%d] emailPricePostMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *EmailPricePostMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailPricePostTooManyRequests creates a EmailPricePostTooManyRequests with default headers values
func NewEmailPricePostTooManyRequests() *EmailPricePostTooManyRequests {
	return &EmailPricePostTooManyRequests{}
}

/*EmailPricePostTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type EmailPricePostTooManyRequests struct {
	Payload string
}

func (o *EmailPricePostTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /email/price][%d] emailPricePostTooManyRequests  %+v", 429, o.Payload)
}

func (o *EmailPricePostTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailPricePostDefault creates a EmailPricePostDefault with default headers values
func NewEmailPricePostDefault(code int) *EmailPricePostDefault {
	return &EmailPricePostDefault{
		_statusCode: code,
	}
}

/*EmailPricePostDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type EmailPricePostDefault struct {
	_statusCode int

	Payload string
}

// Code gets the status code for the email price post default response
func (o *EmailPricePostDefault) Code() int {
	return o._statusCode
}

func (o *EmailPricePostDefault) Error() string {
	return fmt.Sprintf("[POST /email/price][%d] EmailPricePost default  %+v", o._statusCode, o.Payload)
}

func (o *EmailPricePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
