// Code generated by go-swagger; DO NOT EDIT.

package reseller_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ResellerAccountsByClientUserIDGetReader is a Reader for the ResellerAccountsByClientUserIDGet structure.
type ResellerAccountsByClientUserIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResellerAccountsByClientUserIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewResellerAccountsByClientUserIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewResellerAccountsByClientUserIDGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewResellerAccountsByClientUserIDGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewResellerAccountsByClientUserIDGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewResellerAccountsByClientUserIDGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewResellerAccountsByClientUserIDGetMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewResellerAccountsByClientUserIDGetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewResellerAccountsByClientUserIDGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewResellerAccountsByClientUserIDGetOK creates a ResellerAccountsByClientUserIDGetOK with default headers values
func NewResellerAccountsByClientUserIDGetOK() *ResellerAccountsByClientUserIDGetOK {
	return &ResellerAccountsByClientUserIDGetOK{}
}

/*ResellerAccountsByClientUserIDGetOK handles this case with default header values.

SUCCESS
*/
type ResellerAccountsByClientUserIDGetOK struct {
	Payload interface{}
}

func (o *ResellerAccountsByClientUserIDGetOK) Error() string {
	return fmt.Sprintf("[GET /reseller/accounts/{client_user_id}][%d] resellerAccountsByClientUserIdGetOK  %+v", 200, o.Payload)
}

func (o *ResellerAccountsByClientUserIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResellerAccountsByClientUserIDGetBadRequest creates a ResellerAccountsByClientUserIDGetBadRequest with default headers values
func NewResellerAccountsByClientUserIDGetBadRequest() *ResellerAccountsByClientUserIDGetBadRequest {
	return &ResellerAccountsByClientUserIDGetBadRequest{}
}

/*ResellerAccountsByClientUserIDGetBadRequest handles this case with default header values.

BAD_REQUEST
*/
type ResellerAccountsByClientUserIDGetBadRequest struct {
	Payload interface{}
}

func (o *ResellerAccountsByClientUserIDGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /reseller/accounts/{client_user_id}][%d] resellerAccountsByClientUserIdGetBadRequest  %+v", 400, o.Payload)
}

func (o *ResellerAccountsByClientUserIDGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResellerAccountsByClientUserIDGetUnauthorized creates a ResellerAccountsByClientUserIDGetUnauthorized with default headers values
func NewResellerAccountsByClientUserIDGetUnauthorized() *ResellerAccountsByClientUserIDGetUnauthorized {
	return &ResellerAccountsByClientUserIDGetUnauthorized{}
}

/*ResellerAccountsByClientUserIDGetUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type ResellerAccountsByClientUserIDGetUnauthorized struct {
	Payload interface{}
}

func (o *ResellerAccountsByClientUserIDGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /reseller/accounts/{client_user_id}][%d] resellerAccountsByClientUserIdGetUnauthorized  %+v", 401, o.Payload)
}

func (o *ResellerAccountsByClientUserIDGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResellerAccountsByClientUserIDGetForbidden creates a ResellerAccountsByClientUserIDGetForbidden with default headers values
func NewResellerAccountsByClientUserIDGetForbidden() *ResellerAccountsByClientUserIDGetForbidden {
	return &ResellerAccountsByClientUserIDGetForbidden{}
}

/*ResellerAccountsByClientUserIDGetForbidden handles this case with default header values.

FORBIDDEN
*/
type ResellerAccountsByClientUserIDGetForbidden struct {
	Payload interface{}
}

func (o *ResellerAccountsByClientUserIDGetForbidden) Error() string {
	return fmt.Sprintf("[GET /reseller/accounts/{client_user_id}][%d] resellerAccountsByClientUserIdGetForbidden  %+v", 403, o.Payload)
}

func (o *ResellerAccountsByClientUserIDGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResellerAccountsByClientUserIDGetNotFound creates a ResellerAccountsByClientUserIDGetNotFound with default headers values
func NewResellerAccountsByClientUserIDGetNotFound() *ResellerAccountsByClientUserIDGetNotFound {
	return &ResellerAccountsByClientUserIDGetNotFound{}
}

/*ResellerAccountsByClientUserIDGetNotFound handles this case with default header values.

NOT_FOUND
*/
type ResellerAccountsByClientUserIDGetNotFound struct {
	Payload interface{}
}

func (o *ResellerAccountsByClientUserIDGetNotFound) Error() string {
	return fmt.Sprintf("[GET /reseller/accounts/{client_user_id}][%d] resellerAccountsByClientUserIdGetNotFound  %+v", 404, o.Payload)
}

func (o *ResellerAccountsByClientUserIDGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResellerAccountsByClientUserIDGetMethodNotAllowed creates a ResellerAccountsByClientUserIDGetMethodNotAllowed with default headers values
func NewResellerAccountsByClientUserIDGetMethodNotAllowed() *ResellerAccountsByClientUserIDGetMethodNotAllowed {
	return &ResellerAccountsByClientUserIDGetMethodNotAllowed{}
}

/*ResellerAccountsByClientUserIDGetMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type ResellerAccountsByClientUserIDGetMethodNotAllowed struct {
	Payload interface{}
}

func (o *ResellerAccountsByClientUserIDGetMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /reseller/accounts/{client_user_id}][%d] resellerAccountsByClientUserIdGetMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *ResellerAccountsByClientUserIDGetMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResellerAccountsByClientUserIDGetTooManyRequests creates a ResellerAccountsByClientUserIDGetTooManyRequests with default headers values
func NewResellerAccountsByClientUserIDGetTooManyRequests() *ResellerAccountsByClientUserIDGetTooManyRequests {
	return &ResellerAccountsByClientUserIDGetTooManyRequests{}
}

/*ResellerAccountsByClientUserIDGetTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type ResellerAccountsByClientUserIDGetTooManyRequests struct {
	Payload interface{}
}

func (o *ResellerAccountsByClientUserIDGetTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /reseller/accounts/{client_user_id}][%d] resellerAccountsByClientUserIdGetTooManyRequests  %+v", 429, o.Payload)
}

func (o *ResellerAccountsByClientUserIDGetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResellerAccountsByClientUserIDGetDefault creates a ResellerAccountsByClientUserIDGetDefault with default headers values
func NewResellerAccountsByClientUserIDGetDefault(code int) *ResellerAccountsByClientUserIDGetDefault {
	return &ResellerAccountsByClientUserIDGetDefault{
		_statusCode: code,
	}
}

/*ResellerAccountsByClientUserIDGetDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type ResellerAccountsByClientUserIDGetDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the reseller accounts by client user Id get default response
func (o *ResellerAccountsByClientUserIDGetDefault) Code() int {
	return o._statusCode
}

func (o *ResellerAccountsByClientUserIDGetDefault) Error() string {
	return fmt.Sprintf("[GET /reseller/accounts/{client_user_id}][%d] ResellerAccountsByClientUserIdGet default  %+v", o._statusCode, o.Payload)
}

func (o *ResellerAccountsByClientUserIDGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
