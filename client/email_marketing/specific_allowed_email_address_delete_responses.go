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

// SpecificAllowedEmailAddressDeleteReader is a Reader for the SpecificAllowedEmailAddressDelete structure.
type SpecificAllowedEmailAddressDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SpecificAllowedEmailAddressDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSpecificAllowedEmailAddressDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSpecificAllowedEmailAddressDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSpecificAllowedEmailAddressDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSpecificAllowedEmailAddressDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSpecificAllowedEmailAddressDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewSpecificAllowedEmailAddressDeleteMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewSpecificAllowedEmailAddressDeleteTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSpecificAllowedEmailAddressDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSpecificAllowedEmailAddressDeleteOK creates a SpecificAllowedEmailAddressDeleteOK with default headers values
func NewSpecificAllowedEmailAddressDeleteOK() *SpecificAllowedEmailAddressDeleteOK {
	return &SpecificAllowedEmailAddressDeleteOK{}
}

/*SpecificAllowedEmailAddressDeleteOK handles this case with default header values.

SUCCESS
*/
type SpecificAllowedEmailAddressDeleteOK struct {
	Payload interface{}
}

func (o *SpecificAllowedEmailAddressDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /email/addresses/{email_address_id}][%d] specificAllowedEmailAddressDeleteOK  %+v", 200, o.Payload)
}

func (o *SpecificAllowedEmailAddressDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSpecificAllowedEmailAddressDeleteBadRequest creates a SpecificAllowedEmailAddressDeleteBadRequest with default headers values
func NewSpecificAllowedEmailAddressDeleteBadRequest() *SpecificAllowedEmailAddressDeleteBadRequest {
	return &SpecificAllowedEmailAddressDeleteBadRequest{}
}

/*SpecificAllowedEmailAddressDeleteBadRequest handles this case with default header values.

BAD_REQUEST
*/
type SpecificAllowedEmailAddressDeleteBadRequest struct {
	Payload interface{}
}

func (o *SpecificAllowedEmailAddressDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /email/addresses/{email_address_id}][%d] specificAllowedEmailAddressDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *SpecificAllowedEmailAddressDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSpecificAllowedEmailAddressDeleteUnauthorized creates a SpecificAllowedEmailAddressDeleteUnauthorized with default headers values
func NewSpecificAllowedEmailAddressDeleteUnauthorized() *SpecificAllowedEmailAddressDeleteUnauthorized {
	return &SpecificAllowedEmailAddressDeleteUnauthorized{}
}

/*SpecificAllowedEmailAddressDeleteUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type SpecificAllowedEmailAddressDeleteUnauthorized struct {
	Payload interface{}
}

func (o *SpecificAllowedEmailAddressDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /email/addresses/{email_address_id}][%d] specificAllowedEmailAddressDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *SpecificAllowedEmailAddressDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSpecificAllowedEmailAddressDeleteForbidden creates a SpecificAllowedEmailAddressDeleteForbidden with default headers values
func NewSpecificAllowedEmailAddressDeleteForbidden() *SpecificAllowedEmailAddressDeleteForbidden {
	return &SpecificAllowedEmailAddressDeleteForbidden{}
}

/*SpecificAllowedEmailAddressDeleteForbidden handles this case with default header values.

FORBIDDEN
*/
type SpecificAllowedEmailAddressDeleteForbidden struct {
	Payload interface{}
}

func (o *SpecificAllowedEmailAddressDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /email/addresses/{email_address_id}][%d] specificAllowedEmailAddressDeleteForbidden  %+v", 403, o.Payload)
}

func (o *SpecificAllowedEmailAddressDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSpecificAllowedEmailAddressDeleteNotFound creates a SpecificAllowedEmailAddressDeleteNotFound with default headers values
func NewSpecificAllowedEmailAddressDeleteNotFound() *SpecificAllowedEmailAddressDeleteNotFound {
	return &SpecificAllowedEmailAddressDeleteNotFound{}
}

/*SpecificAllowedEmailAddressDeleteNotFound handles this case with default header values.

NOT_FOUND
*/
type SpecificAllowedEmailAddressDeleteNotFound struct {
	Payload interface{}
}

func (o *SpecificAllowedEmailAddressDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /email/addresses/{email_address_id}][%d] specificAllowedEmailAddressDeleteNotFound  %+v", 404, o.Payload)
}

func (o *SpecificAllowedEmailAddressDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSpecificAllowedEmailAddressDeleteMethodNotAllowed creates a SpecificAllowedEmailAddressDeleteMethodNotAllowed with default headers values
func NewSpecificAllowedEmailAddressDeleteMethodNotAllowed() *SpecificAllowedEmailAddressDeleteMethodNotAllowed {
	return &SpecificAllowedEmailAddressDeleteMethodNotAllowed{}
}

/*SpecificAllowedEmailAddressDeleteMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type SpecificAllowedEmailAddressDeleteMethodNotAllowed struct {
	Payload interface{}
}

func (o *SpecificAllowedEmailAddressDeleteMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /email/addresses/{email_address_id}][%d] specificAllowedEmailAddressDeleteMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *SpecificAllowedEmailAddressDeleteMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSpecificAllowedEmailAddressDeleteTooManyRequests creates a SpecificAllowedEmailAddressDeleteTooManyRequests with default headers values
func NewSpecificAllowedEmailAddressDeleteTooManyRequests() *SpecificAllowedEmailAddressDeleteTooManyRequests {
	return &SpecificAllowedEmailAddressDeleteTooManyRequests{}
}

/*SpecificAllowedEmailAddressDeleteTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type SpecificAllowedEmailAddressDeleteTooManyRequests struct {
	Payload interface{}
}

func (o *SpecificAllowedEmailAddressDeleteTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /email/addresses/{email_address_id}][%d] specificAllowedEmailAddressDeleteTooManyRequests  %+v", 429, o.Payload)
}

func (o *SpecificAllowedEmailAddressDeleteTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSpecificAllowedEmailAddressDeleteDefault creates a SpecificAllowedEmailAddressDeleteDefault with default headers values
func NewSpecificAllowedEmailAddressDeleteDefault(code int) *SpecificAllowedEmailAddressDeleteDefault {
	return &SpecificAllowedEmailAddressDeleteDefault{
		_statusCode: code,
	}
}

/*SpecificAllowedEmailAddressDeleteDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type SpecificAllowedEmailAddressDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the specific allowed email address delete default response
func (o *SpecificAllowedEmailAddressDeleteDefault) Code() int {
	return o._statusCode
}

func (o *SpecificAllowedEmailAddressDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /email/addresses/{email_address_id}][%d] SpecificAllowedEmailAddressDelete default  %+v", o._statusCode, o.Payload)
}

func (o *SpecificAllowedEmailAddressDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
