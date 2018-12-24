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

// ResellerAccountsGetReader is a Reader for the ResellerAccountsGet structure.
type ResellerAccountsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResellerAccountsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewResellerAccountsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewResellerAccountsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewResellerAccountsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewResellerAccountsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewResellerAccountsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewResellerAccountsGetMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewResellerAccountsGetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewResellerAccountsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewResellerAccountsGetOK creates a ResellerAccountsGetOK with default headers values
func NewResellerAccountsGetOK() *ResellerAccountsGetOK {
	return &ResellerAccountsGetOK{}
}

/*ResellerAccountsGetOK handles this case with default header values.

SUCCESS
*/
type ResellerAccountsGetOK struct {
	Payload interface{}
}

func (o *ResellerAccountsGetOK) Error() string {
	return fmt.Sprintf("[GET /reseller/accounts][%d] resellerAccountsGetOK  %+v", 200, o.Payload)
}

func (o *ResellerAccountsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResellerAccountsGetBadRequest creates a ResellerAccountsGetBadRequest with default headers values
func NewResellerAccountsGetBadRequest() *ResellerAccountsGetBadRequest {
	return &ResellerAccountsGetBadRequest{}
}

/*ResellerAccountsGetBadRequest handles this case with default header values.

BAD_REQUEST
*/
type ResellerAccountsGetBadRequest struct {
	Payload interface{}
}

func (o *ResellerAccountsGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /reseller/accounts][%d] resellerAccountsGetBadRequest  %+v", 400, o.Payload)
}

func (o *ResellerAccountsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResellerAccountsGetUnauthorized creates a ResellerAccountsGetUnauthorized with default headers values
func NewResellerAccountsGetUnauthorized() *ResellerAccountsGetUnauthorized {
	return &ResellerAccountsGetUnauthorized{}
}

/*ResellerAccountsGetUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type ResellerAccountsGetUnauthorized struct {
	Payload interface{}
}

func (o *ResellerAccountsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /reseller/accounts][%d] resellerAccountsGetUnauthorized  %+v", 401, o.Payload)
}

func (o *ResellerAccountsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResellerAccountsGetForbidden creates a ResellerAccountsGetForbidden with default headers values
func NewResellerAccountsGetForbidden() *ResellerAccountsGetForbidden {
	return &ResellerAccountsGetForbidden{}
}

/*ResellerAccountsGetForbidden handles this case with default header values.

FORBIDDEN
*/
type ResellerAccountsGetForbidden struct {
	Payload interface{}
}

func (o *ResellerAccountsGetForbidden) Error() string {
	return fmt.Sprintf("[GET /reseller/accounts][%d] resellerAccountsGetForbidden  %+v", 403, o.Payload)
}

func (o *ResellerAccountsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResellerAccountsGetNotFound creates a ResellerAccountsGetNotFound with default headers values
func NewResellerAccountsGetNotFound() *ResellerAccountsGetNotFound {
	return &ResellerAccountsGetNotFound{}
}

/*ResellerAccountsGetNotFound handles this case with default header values.

NOT_FOUND
*/
type ResellerAccountsGetNotFound struct {
	Payload interface{}
}

func (o *ResellerAccountsGetNotFound) Error() string {
	return fmt.Sprintf("[GET /reseller/accounts][%d] resellerAccountsGetNotFound  %+v", 404, o.Payload)
}

func (o *ResellerAccountsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResellerAccountsGetMethodNotAllowed creates a ResellerAccountsGetMethodNotAllowed with default headers values
func NewResellerAccountsGetMethodNotAllowed() *ResellerAccountsGetMethodNotAllowed {
	return &ResellerAccountsGetMethodNotAllowed{}
}

/*ResellerAccountsGetMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type ResellerAccountsGetMethodNotAllowed struct {
	Payload interface{}
}

func (o *ResellerAccountsGetMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /reseller/accounts][%d] resellerAccountsGetMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *ResellerAccountsGetMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResellerAccountsGetTooManyRequests creates a ResellerAccountsGetTooManyRequests with default headers values
func NewResellerAccountsGetTooManyRequests() *ResellerAccountsGetTooManyRequests {
	return &ResellerAccountsGetTooManyRequests{}
}

/*ResellerAccountsGetTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type ResellerAccountsGetTooManyRequests struct {
	Payload interface{}
}

func (o *ResellerAccountsGetTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /reseller/accounts][%d] resellerAccountsGetTooManyRequests  %+v", 429, o.Payload)
}

func (o *ResellerAccountsGetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResellerAccountsGetDefault creates a ResellerAccountsGetDefault with default headers values
func NewResellerAccountsGetDefault(code int) *ResellerAccountsGetDefault {
	return &ResellerAccountsGetDefault{
		_statusCode: code,
	}
}

/*ResellerAccountsGetDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type ResellerAccountsGetDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the reseller accounts get default response
func (o *ResellerAccountsGetDefault) Code() int {
	return o._statusCode
}

func (o *ResellerAccountsGetDefault) Error() string {
	return fmt.Sprintf("[GET /reseller/accounts][%d] ResellerAccountsGet default  %+v", o._statusCode, o.Payload)
}

func (o *ResellerAccountsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
