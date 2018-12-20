// Code generated by go-swagger; DO NOT EDIT.

package email_marketing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// AllowedEmailAddressPostReader is a Reader for the AllowedEmailAddressPost structure.
type AllowedEmailAddressPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllowedEmailAddressPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAllowedEmailAddressPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAllowedEmailAddressPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAllowedEmailAddressPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewAllowedEmailAddressPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAllowedEmailAddressPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewAllowedEmailAddressPostMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewAllowedEmailAddressPostTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewAllowedEmailAddressPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAllowedEmailAddressPostOK creates a AllowedEmailAddressPostOK with default headers values
func NewAllowedEmailAddressPostOK() *AllowedEmailAddressPostOK {
	return &AllowedEmailAddressPostOK{}
}

/*AllowedEmailAddressPostOK handles this case with default header values.

SUCCESS
*/
type AllowedEmailAddressPostOK struct {
	Payload string
}

func (o *AllowedEmailAddressPostOK) Error() string {
	return fmt.Sprintf("[POST /email/addresses][%d] allowedEmailAddressPostOK  %+v", 200, o.Payload)
}

func (o *AllowedEmailAddressPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllowedEmailAddressPostBadRequest creates a AllowedEmailAddressPostBadRequest with default headers values
func NewAllowedEmailAddressPostBadRequest() *AllowedEmailAddressPostBadRequest {
	return &AllowedEmailAddressPostBadRequest{}
}

/*AllowedEmailAddressPostBadRequest handles this case with default header values.

BAD_REQUEST
*/
type AllowedEmailAddressPostBadRequest struct {
	Payload string
}

func (o *AllowedEmailAddressPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /email/addresses][%d] allowedEmailAddressPostBadRequest  %+v", 400, o.Payload)
}

func (o *AllowedEmailAddressPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllowedEmailAddressPostUnauthorized creates a AllowedEmailAddressPostUnauthorized with default headers values
func NewAllowedEmailAddressPostUnauthorized() *AllowedEmailAddressPostUnauthorized {
	return &AllowedEmailAddressPostUnauthorized{}
}

/*AllowedEmailAddressPostUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type AllowedEmailAddressPostUnauthorized struct {
	Payload string
}

func (o *AllowedEmailAddressPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /email/addresses][%d] allowedEmailAddressPostUnauthorized  %+v", 401, o.Payload)
}

func (o *AllowedEmailAddressPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllowedEmailAddressPostForbidden creates a AllowedEmailAddressPostForbidden with default headers values
func NewAllowedEmailAddressPostForbidden() *AllowedEmailAddressPostForbidden {
	return &AllowedEmailAddressPostForbidden{}
}

/*AllowedEmailAddressPostForbidden handles this case with default header values.

FORBIDDEN
*/
type AllowedEmailAddressPostForbidden struct {
	Payload string
}

func (o *AllowedEmailAddressPostForbidden) Error() string {
	return fmt.Sprintf("[POST /email/addresses][%d] allowedEmailAddressPostForbidden  %+v", 403, o.Payload)
}

func (o *AllowedEmailAddressPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllowedEmailAddressPostNotFound creates a AllowedEmailAddressPostNotFound with default headers values
func NewAllowedEmailAddressPostNotFound() *AllowedEmailAddressPostNotFound {
	return &AllowedEmailAddressPostNotFound{}
}

/*AllowedEmailAddressPostNotFound handles this case with default header values.

NOT_FOUND
*/
type AllowedEmailAddressPostNotFound struct {
	Payload string
}

func (o *AllowedEmailAddressPostNotFound) Error() string {
	return fmt.Sprintf("[POST /email/addresses][%d] allowedEmailAddressPostNotFound  %+v", 404, o.Payload)
}

func (o *AllowedEmailAddressPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllowedEmailAddressPostMethodNotAllowed creates a AllowedEmailAddressPostMethodNotAllowed with default headers values
func NewAllowedEmailAddressPostMethodNotAllowed() *AllowedEmailAddressPostMethodNotAllowed {
	return &AllowedEmailAddressPostMethodNotAllowed{}
}

/*AllowedEmailAddressPostMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type AllowedEmailAddressPostMethodNotAllowed struct {
	Payload string
}

func (o *AllowedEmailAddressPostMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /email/addresses][%d] allowedEmailAddressPostMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *AllowedEmailAddressPostMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllowedEmailAddressPostTooManyRequests creates a AllowedEmailAddressPostTooManyRequests with default headers values
func NewAllowedEmailAddressPostTooManyRequests() *AllowedEmailAddressPostTooManyRequests {
	return &AllowedEmailAddressPostTooManyRequests{}
}

/*AllowedEmailAddressPostTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type AllowedEmailAddressPostTooManyRequests struct {
	Payload string
}

func (o *AllowedEmailAddressPostTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /email/addresses][%d] allowedEmailAddressPostTooManyRequests  %+v", 429, o.Payload)
}

func (o *AllowedEmailAddressPostTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllowedEmailAddressPostDefault creates a AllowedEmailAddressPostDefault with default headers values
func NewAllowedEmailAddressPostDefault(code int) *AllowedEmailAddressPostDefault {
	return &AllowedEmailAddressPostDefault{
		_statusCode: code,
	}
}

/*AllowedEmailAddressPostDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type AllowedEmailAddressPostDefault struct {
	_statusCode int

	Payload string
}

// Code gets the status code for the allowed email address post default response
func (o *AllowedEmailAddressPostDefault) Code() int {
	return o._statusCode
}

func (o *AllowedEmailAddressPostDefault) Error() string {
	return fmt.Sprintf("[POST /email/addresses][%d] AllowedEmailAddressPost default  %+v", o._statusCode, o.Payload)
}

func (o *AllowedEmailAddressPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
