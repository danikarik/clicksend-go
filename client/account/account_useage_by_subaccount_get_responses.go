// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// AccountUseageBySubaccountGetReader is a Reader for the AccountUseageBySubaccountGet structure.
type AccountUseageBySubaccountGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountUseageBySubaccountGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAccountUseageBySubaccountGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAccountUseageBySubaccountGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAccountUseageBySubaccountGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewAccountUseageBySubaccountGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAccountUseageBySubaccountGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewAccountUseageBySubaccountGetMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewAccountUseageBySubaccountGetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewAccountUseageBySubaccountGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAccountUseageBySubaccountGetOK creates a AccountUseageBySubaccountGetOK with default headers values
func NewAccountUseageBySubaccountGetOK() *AccountUseageBySubaccountGetOK {
	return &AccountUseageBySubaccountGetOK{}
}

/*AccountUseageBySubaccountGetOK handles this case with default header values.

SUCCESS
*/
type AccountUseageBySubaccountGetOK struct {
	Payload string
}

func (o *AccountUseageBySubaccountGetOK) Error() string {
	return fmt.Sprintf("[GET /account/usage/{year}/{month}/subaccount][%d] accountUseageBySubaccountGetOK  %+v", 200, o.Payload)
}

func (o *AccountUseageBySubaccountGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountUseageBySubaccountGetBadRequest creates a AccountUseageBySubaccountGetBadRequest with default headers values
func NewAccountUseageBySubaccountGetBadRequest() *AccountUseageBySubaccountGetBadRequest {
	return &AccountUseageBySubaccountGetBadRequest{}
}

/*AccountUseageBySubaccountGetBadRequest handles this case with default header values.

BAD_REQUEST
*/
type AccountUseageBySubaccountGetBadRequest struct {
	Payload string
}

func (o *AccountUseageBySubaccountGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /account/usage/{year}/{month}/subaccount][%d] accountUseageBySubaccountGetBadRequest  %+v", 400, o.Payload)
}

func (o *AccountUseageBySubaccountGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountUseageBySubaccountGetUnauthorized creates a AccountUseageBySubaccountGetUnauthorized with default headers values
func NewAccountUseageBySubaccountGetUnauthorized() *AccountUseageBySubaccountGetUnauthorized {
	return &AccountUseageBySubaccountGetUnauthorized{}
}

/*AccountUseageBySubaccountGetUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type AccountUseageBySubaccountGetUnauthorized struct {
	Payload string
}

func (o *AccountUseageBySubaccountGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /account/usage/{year}/{month}/subaccount][%d] accountUseageBySubaccountGetUnauthorized  %+v", 401, o.Payload)
}

func (o *AccountUseageBySubaccountGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountUseageBySubaccountGetForbidden creates a AccountUseageBySubaccountGetForbidden with default headers values
func NewAccountUseageBySubaccountGetForbidden() *AccountUseageBySubaccountGetForbidden {
	return &AccountUseageBySubaccountGetForbidden{}
}

/*AccountUseageBySubaccountGetForbidden handles this case with default header values.

FORBIDDEN
*/
type AccountUseageBySubaccountGetForbidden struct {
	Payload string
}

func (o *AccountUseageBySubaccountGetForbidden) Error() string {
	return fmt.Sprintf("[GET /account/usage/{year}/{month}/subaccount][%d] accountUseageBySubaccountGetForbidden  %+v", 403, o.Payload)
}

func (o *AccountUseageBySubaccountGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountUseageBySubaccountGetNotFound creates a AccountUseageBySubaccountGetNotFound with default headers values
func NewAccountUseageBySubaccountGetNotFound() *AccountUseageBySubaccountGetNotFound {
	return &AccountUseageBySubaccountGetNotFound{}
}

/*AccountUseageBySubaccountGetNotFound handles this case with default header values.

NOT_FOUND
*/
type AccountUseageBySubaccountGetNotFound struct {
	Payload string
}

func (o *AccountUseageBySubaccountGetNotFound) Error() string {
	return fmt.Sprintf("[GET /account/usage/{year}/{month}/subaccount][%d] accountUseageBySubaccountGetNotFound  %+v", 404, o.Payload)
}

func (o *AccountUseageBySubaccountGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountUseageBySubaccountGetMethodNotAllowed creates a AccountUseageBySubaccountGetMethodNotAllowed with default headers values
func NewAccountUseageBySubaccountGetMethodNotAllowed() *AccountUseageBySubaccountGetMethodNotAllowed {
	return &AccountUseageBySubaccountGetMethodNotAllowed{}
}

/*AccountUseageBySubaccountGetMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type AccountUseageBySubaccountGetMethodNotAllowed struct {
	Payload string
}

func (o *AccountUseageBySubaccountGetMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /account/usage/{year}/{month}/subaccount][%d] accountUseageBySubaccountGetMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *AccountUseageBySubaccountGetMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountUseageBySubaccountGetTooManyRequests creates a AccountUseageBySubaccountGetTooManyRequests with default headers values
func NewAccountUseageBySubaccountGetTooManyRequests() *AccountUseageBySubaccountGetTooManyRequests {
	return &AccountUseageBySubaccountGetTooManyRequests{}
}

/*AccountUseageBySubaccountGetTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type AccountUseageBySubaccountGetTooManyRequests struct {
	Payload string
}

func (o *AccountUseageBySubaccountGetTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /account/usage/{year}/{month}/subaccount][%d] accountUseageBySubaccountGetTooManyRequests  %+v", 429, o.Payload)
}

func (o *AccountUseageBySubaccountGetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountUseageBySubaccountGetDefault creates a AccountUseageBySubaccountGetDefault with default headers values
func NewAccountUseageBySubaccountGetDefault(code int) *AccountUseageBySubaccountGetDefault {
	return &AccountUseageBySubaccountGetDefault{
		_statusCode: code,
	}
}

/*AccountUseageBySubaccountGetDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type AccountUseageBySubaccountGetDefault struct {
	_statusCode int

	Payload string
}

// Code gets the status code for the account useage by subaccount get default response
func (o *AccountUseageBySubaccountGetDefault) Code() int {
	return o._statusCode
}

func (o *AccountUseageBySubaccountGetDefault) Error() string {
	return fmt.Sprintf("[GET /account/usage/{year}/{month}/subaccount][%d] AccountUseageBySubaccountGet default  %+v", o._statusCode, o.Payload)
}

func (o *AccountUseageBySubaccountGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
