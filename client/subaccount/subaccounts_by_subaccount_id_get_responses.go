// Code generated by go-swagger; DO NOT EDIT.

package subaccount

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// SubaccountsBySubaccountIDGetReader is a Reader for the SubaccountsBySubaccountIDGet structure.
type SubaccountsBySubaccountIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubaccountsBySubaccountIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSubaccountsBySubaccountIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSubaccountsBySubaccountIDGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSubaccountsBySubaccountIDGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSubaccountsBySubaccountIDGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSubaccountsBySubaccountIDGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewSubaccountsBySubaccountIDGetMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewSubaccountsBySubaccountIDGetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSubaccountsBySubaccountIDGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSubaccountsBySubaccountIDGetOK creates a SubaccountsBySubaccountIDGetOK with default headers values
func NewSubaccountsBySubaccountIDGetOK() *SubaccountsBySubaccountIDGetOK {
	return &SubaccountsBySubaccountIDGetOK{}
}

/*SubaccountsBySubaccountIDGetOK handles this case with default header values.

SUCCESS
*/
type SubaccountsBySubaccountIDGetOK struct {
	Payload interface{}
}

func (o *SubaccountsBySubaccountIDGetOK) Error() string {
	return fmt.Sprintf("[GET /subaccounts/{subaccount_id}][%d] subaccountsBySubaccountIdGetOK  %+v", 200, o.Payload)
}

func (o *SubaccountsBySubaccountIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubaccountsBySubaccountIDGetBadRequest creates a SubaccountsBySubaccountIDGetBadRequest with default headers values
func NewSubaccountsBySubaccountIDGetBadRequest() *SubaccountsBySubaccountIDGetBadRequest {
	return &SubaccountsBySubaccountIDGetBadRequest{}
}

/*SubaccountsBySubaccountIDGetBadRequest handles this case with default header values.

BAD_REQUEST
*/
type SubaccountsBySubaccountIDGetBadRequest struct {
	Payload interface{}
}

func (o *SubaccountsBySubaccountIDGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /subaccounts/{subaccount_id}][%d] subaccountsBySubaccountIdGetBadRequest  %+v", 400, o.Payload)
}

func (o *SubaccountsBySubaccountIDGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubaccountsBySubaccountIDGetUnauthorized creates a SubaccountsBySubaccountIDGetUnauthorized with default headers values
func NewSubaccountsBySubaccountIDGetUnauthorized() *SubaccountsBySubaccountIDGetUnauthorized {
	return &SubaccountsBySubaccountIDGetUnauthorized{}
}

/*SubaccountsBySubaccountIDGetUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type SubaccountsBySubaccountIDGetUnauthorized struct {
	Payload interface{}
}

func (o *SubaccountsBySubaccountIDGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /subaccounts/{subaccount_id}][%d] subaccountsBySubaccountIdGetUnauthorized  %+v", 401, o.Payload)
}

func (o *SubaccountsBySubaccountIDGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubaccountsBySubaccountIDGetForbidden creates a SubaccountsBySubaccountIDGetForbidden with default headers values
func NewSubaccountsBySubaccountIDGetForbidden() *SubaccountsBySubaccountIDGetForbidden {
	return &SubaccountsBySubaccountIDGetForbidden{}
}

/*SubaccountsBySubaccountIDGetForbidden handles this case with default header values.

FORBIDDEN
*/
type SubaccountsBySubaccountIDGetForbidden struct {
	Payload interface{}
}

func (o *SubaccountsBySubaccountIDGetForbidden) Error() string {
	return fmt.Sprintf("[GET /subaccounts/{subaccount_id}][%d] subaccountsBySubaccountIdGetForbidden  %+v", 403, o.Payload)
}

func (o *SubaccountsBySubaccountIDGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubaccountsBySubaccountIDGetNotFound creates a SubaccountsBySubaccountIDGetNotFound with default headers values
func NewSubaccountsBySubaccountIDGetNotFound() *SubaccountsBySubaccountIDGetNotFound {
	return &SubaccountsBySubaccountIDGetNotFound{}
}

/*SubaccountsBySubaccountIDGetNotFound handles this case with default header values.

NOT_FOUND
*/
type SubaccountsBySubaccountIDGetNotFound struct {
	Payload interface{}
}

func (o *SubaccountsBySubaccountIDGetNotFound) Error() string {
	return fmt.Sprintf("[GET /subaccounts/{subaccount_id}][%d] subaccountsBySubaccountIdGetNotFound  %+v", 404, o.Payload)
}

func (o *SubaccountsBySubaccountIDGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubaccountsBySubaccountIDGetMethodNotAllowed creates a SubaccountsBySubaccountIDGetMethodNotAllowed with default headers values
func NewSubaccountsBySubaccountIDGetMethodNotAllowed() *SubaccountsBySubaccountIDGetMethodNotAllowed {
	return &SubaccountsBySubaccountIDGetMethodNotAllowed{}
}

/*SubaccountsBySubaccountIDGetMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type SubaccountsBySubaccountIDGetMethodNotAllowed struct {
	Payload interface{}
}

func (o *SubaccountsBySubaccountIDGetMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /subaccounts/{subaccount_id}][%d] subaccountsBySubaccountIdGetMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *SubaccountsBySubaccountIDGetMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubaccountsBySubaccountIDGetTooManyRequests creates a SubaccountsBySubaccountIDGetTooManyRequests with default headers values
func NewSubaccountsBySubaccountIDGetTooManyRequests() *SubaccountsBySubaccountIDGetTooManyRequests {
	return &SubaccountsBySubaccountIDGetTooManyRequests{}
}

/*SubaccountsBySubaccountIDGetTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type SubaccountsBySubaccountIDGetTooManyRequests struct {
	Payload interface{}
}

func (o *SubaccountsBySubaccountIDGetTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /subaccounts/{subaccount_id}][%d] subaccountsBySubaccountIdGetTooManyRequests  %+v", 429, o.Payload)
}

func (o *SubaccountsBySubaccountIDGetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubaccountsBySubaccountIDGetDefault creates a SubaccountsBySubaccountIDGetDefault with default headers values
func NewSubaccountsBySubaccountIDGetDefault(code int) *SubaccountsBySubaccountIDGetDefault {
	return &SubaccountsBySubaccountIDGetDefault{
		_statusCode: code,
	}
}

/*SubaccountsBySubaccountIDGetDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type SubaccountsBySubaccountIDGetDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the subaccounts by subaccount Id get default response
func (o *SubaccountsBySubaccountIDGetDefault) Code() int {
	return o._statusCode
}

func (o *SubaccountsBySubaccountIDGetDefault) Error() string {
	return fmt.Sprintf("[GET /subaccounts/{subaccount_id}][%d] SubaccountsBySubaccountIdGet default  %+v", o._statusCode, o.Payload)
}

func (o *SubaccountsBySubaccountIDGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
