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

// AccountVerifyVerifyByActivationTokenPutReader is a Reader for the AccountVerifyVerifyByActivationTokenPut structure.
type AccountVerifyVerifyByActivationTokenPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountVerifyVerifyByActivationTokenPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAccountVerifyVerifyByActivationTokenPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAccountVerifyVerifyByActivationTokenPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAccountVerifyVerifyByActivationTokenPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewAccountVerifyVerifyByActivationTokenPutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAccountVerifyVerifyByActivationTokenPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewAccountVerifyVerifyByActivationTokenPutMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewAccountVerifyVerifyByActivationTokenPutTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewAccountVerifyVerifyByActivationTokenPutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAccountVerifyVerifyByActivationTokenPutOK creates a AccountVerifyVerifyByActivationTokenPutOK with default headers values
func NewAccountVerifyVerifyByActivationTokenPutOK() *AccountVerifyVerifyByActivationTokenPutOK {
	return &AccountVerifyVerifyByActivationTokenPutOK{}
}

/*AccountVerifyVerifyByActivationTokenPutOK handles this case with default header values.

SUCCESS
*/
type AccountVerifyVerifyByActivationTokenPutOK struct {
	Payload interface{}
}

func (o *AccountVerifyVerifyByActivationTokenPutOK) Error() string {
	return fmt.Sprintf("[PUT /account-verify/verify/{activation_token}][%d] accountVerifyVerifyByActivationTokenPutOK  %+v", 200, o.Payload)
}

func (o *AccountVerifyVerifyByActivationTokenPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountVerifyVerifyByActivationTokenPutBadRequest creates a AccountVerifyVerifyByActivationTokenPutBadRequest with default headers values
func NewAccountVerifyVerifyByActivationTokenPutBadRequest() *AccountVerifyVerifyByActivationTokenPutBadRequest {
	return &AccountVerifyVerifyByActivationTokenPutBadRequest{}
}

/*AccountVerifyVerifyByActivationTokenPutBadRequest handles this case with default header values.

BAD_REQUEST
*/
type AccountVerifyVerifyByActivationTokenPutBadRequest struct {
	Payload interface{}
}

func (o *AccountVerifyVerifyByActivationTokenPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /account-verify/verify/{activation_token}][%d] accountVerifyVerifyByActivationTokenPutBadRequest  %+v", 400, o.Payload)
}

func (o *AccountVerifyVerifyByActivationTokenPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountVerifyVerifyByActivationTokenPutUnauthorized creates a AccountVerifyVerifyByActivationTokenPutUnauthorized with default headers values
func NewAccountVerifyVerifyByActivationTokenPutUnauthorized() *AccountVerifyVerifyByActivationTokenPutUnauthorized {
	return &AccountVerifyVerifyByActivationTokenPutUnauthorized{}
}

/*AccountVerifyVerifyByActivationTokenPutUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type AccountVerifyVerifyByActivationTokenPutUnauthorized struct {
	Payload interface{}
}

func (o *AccountVerifyVerifyByActivationTokenPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /account-verify/verify/{activation_token}][%d] accountVerifyVerifyByActivationTokenPutUnauthorized  %+v", 401, o.Payload)
}

func (o *AccountVerifyVerifyByActivationTokenPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountVerifyVerifyByActivationTokenPutForbidden creates a AccountVerifyVerifyByActivationTokenPutForbidden with default headers values
func NewAccountVerifyVerifyByActivationTokenPutForbidden() *AccountVerifyVerifyByActivationTokenPutForbidden {
	return &AccountVerifyVerifyByActivationTokenPutForbidden{}
}

/*AccountVerifyVerifyByActivationTokenPutForbidden handles this case with default header values.

FORBIDDEN
*/
type AccountVerifyVerifyByActivationTokenPutForbidden struct {
	Payload interface{}
}

func (o *AccountVerifyVerifyByActivationTokenPutForbidden) Error() string {
	return fmt.Sprintf("[PUT /account-verify/verify/{activation_token}][%d] accountVerifyVerifyByActivationTokenPutForbidden  %+v", 403, o.Payload)
}

func (o *AccountVerifyVerifyByActivationTokenPutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountVerifyVerifyByActivationTokenPutNotFound creates a AccountVerifyVerifyByActivationTokenPutNotFound with default headers values
func NewAccountVerifyVerifyByActivationTokenPutNotFound() *AccountVerifyVerifyByActivationTokenPutNotFound {
	return &AccountVerifyVerifyByActivationTokenPutNotFound{}
}

/*AccountVerifyVerifyByActivationTokenPutNotFound handles this case with default header values.

NOT_FOUND
*/
type AccountVerifyVerifyByActivationTokenPutNotFound struct {
	Payload interface{}
}

func (o *AccountVerifyVerifyByActivationTokenPutNotFound) Error() string {
	return fmt.Sprintf("[PUT /account-verify/verify/{activation_token}][%d] accountVerifyVerifyByActivationTokenPutNotFound  %+v", 404, o.Payload)
}

func (o *AccountVerifyVerifyByActivationTokenPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountVerifyVerifyByActivationTokenPutMethodNotAllowed creates a AccountVerifyVerifyByActivationTokenPutMethodNotAllowed with default headers values
func NewAccountVerifyVerifyByActivationTokenPutMethodNotAllowed() *AccountVerifyVerifyByActivationTokenPutMethodNotAllowed {
	return &AccountVerifyVerifyByActivationTokenPutMethodNotAllowed{}
}

/*AccountVerifyVerifyByActivationTokenPutMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type AccountVerifyVerifyByActivationTokenPutMethodNotAllowed struct {
	Payload interface{}
}

func (o *AccountVerifyVerifyByActivationTokenPutMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /account-verify/verify/{activation_token}][%d] accountVerifyVerifyByActivationTokenPutMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *AccountVerifyVerifyByActivationTokenPutMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountVerifyVerifyByActivationTokenPutTooManyRequests creates a AccountVerifyVerifyByActivationTokenPutTooManyRequests with default headers values
func NewAccountVerifyVerifyByActivationTokenPutTooManyRequests() *AccountVerifyVerifyByActivationTokenPutTooManyRequests {
	return &AccountVerifyVerifyByActivationTokenPutTooManyRequests{}
}

/*AccountVerifyVerifyByActivationTokenPutTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type AccountVerifyVerifyByActivationTokenPutTooManyRequests struct {
	Payload interface{}
}

func (o *AccountVerifyVerifyByActivationTokenPutTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /account-verify/verify/{activation_token}][%d] accountVerifyVerifyByActivationTokenPutTooManyRequests  %+v", 429, o.Payload)
}

func (o *AccountVerifyVerifyByActivationTokenPutTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountVerifyVerifyByActivationTokenPutDefault creates a AccountVerifyVerifyByActivationTokenPutDefault with default headers values
func NewAccountVerifyVerifyByActivationTokenPutDefault(code int) *AccountVerifyVerifyByActivationTokenPutDefault {
	return &AccountVerifyVerifyByActivationTokenPutDefault{
		_statusCode: code,
	}
}

/*AccountVerifyVerifyByActivationTokenPutDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type AccountVerifyVerifyByActivationTokenPutDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the account verify verify by activation token put default response
func (o *AccountVerifyVerifyByActivationTokenPutDefault) Code() int {
	return o._statusCode
}

func (o *AccountVerifyVerifyByActivationTokenPutDefault) Error() string {
	return fmt.Sprintf("[PUT /account-verify/verify/{activation_token}][%d] AccountVerifyVerifyByActivationTokenPut default  %+v", o._statusCode, o.Payload)
}

func (o *AccountVerifyVerifyByActivationTokenPutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
