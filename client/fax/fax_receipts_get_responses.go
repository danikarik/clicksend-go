// Code generated by go-swagger; DO NOT EDIT.

package fax

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// FAXReceiptsGetReader is a Reader for the FAXReceiptsGet structure.
type FAXReceiptsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FAXReceiptsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFAXReceiptsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewFAXReceiptsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewFAXReceiptsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewFAXReceiptsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewFAXReceiptsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewFAXReceiptsGetMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewFAXReceiptsGetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewFAXReceiptsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFAXReceiptsGetOK creates a FAXReceiptsGetOK with default headers values
func NewFAXReceiptsGetOK() *FAXReceiptsGetOK {
	return &FAXReceiptsGetOK{}
}

/*FAXReceiptsGetOK handles this case with default header values.

SUCCESS
*/
type FAXReceiptsGetOK struct {
	Payload interface{}
}

func (o *FAXReceiptsGetOK) Error() string {
	return fmt.Sprintf("[GET /fax/receipts][%d] faxReceiptsGetOK  %+v", 200, o.Payload)
}

func (o *FAXReceiptsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFAXReceiptsGetBadRequest creates a FAXReceiptsGetBadRequest with default headers values
func NewFAXReceiptsGetBadRequest() *FAXReceiptsGetBadRequest {
	return &FAXReceiptsGetBadRequest{}
}

/*FAXReceiptsGetBadRequest handles this case with default header values.

BAD_REQUEST
*/
type FAXReceiptsGetBadRequest struct {
	Payload interface{}
}

func (o *FAXReceiptsGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /fax/receipts][%d] faxReceiptsGetBadRequest  %+v", 400, o.Payload)
}

func (o *FAXReceiptsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFAXReceiptsGetUnauthorized creates a FAXReceiptsGetUnauthorized with default headers values
func NewFAXReceiptsGetUnauthorized() *FAXReceiptsGetUnauthorized {
	return &FAXReceiptsGetUnauthorized{}
}

/*FAXReceiptsGetUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type FAXReceiptsGetUnauthorized struct {
	Payload interface{}
}

func (o *FAXReceiptsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /fax/receipts][%d] faxReceiptsGetUnauthorized  %+v", 401, o.Payload)
}

func (o *FAXReceiptsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFAXReceiptsGetForbidden creates a FAXReceiptsGetForbidden with default headers values
func NewFAXReceiptsGetForbidden() *FAXReceiptsGetForbidden {
	return &FAXReceiptsGetForbidden{}
}

/*FAXReceiptsGetForbidden handles this case with default header values.

FORBIDDEN
*/
type FAXReceiptsGetForbidden struct {
	Payload interface{}
}

func (o *FAXReceiptsGetForbidden) Error() string {
	return fmt.Sprintf("[GET /fax/receipts][%d] faxReceiptsGetForbidden  %+v", 403, o.Payload)
}

func (o *FAXReceiptsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFAXReceiptsGetNotFound creates a FAXReceiptsGetNotFound with default headers values
func NewFAXReceiptsGetNotFound() *FAXReceiptsGetNotFound {
	return &FAXReceiptsGetNotFound{}
}

/*FAXReceiptsGetNotFound handles this case with default header values.

NOT_FOUND
*/
type FAXReceiptsGetNotFound struct {
	Payload interface{}
}

func (o *FAXReceiptsGetNotFound) Error() string {
	return fmt.Sprintf("[GET /fax/receipts][%d] faxReceiptsGetNotFound  %+v", 404, o.Payload)
}

func (o *FAXReceiptsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFAXReceiptsGetMethodNotAllowed creates a FAXReceiptsGetMethodNotAllowed with default headers values
func NewFAXReceiptsGetMethodNotAllowed() *FAXReceiptsGetMethodNotAllowed {
	return &FAXReceiptsGetMethodNotAllowed{}
}

/*FAXReceiptsGetMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type FAXReceiptsGetMethodNotAllowed struct {
	Payload interface{}
}

func (o *FAXReceiptsGetMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /fax/receipts][%d] faxReceiptsGetMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *FAXReceiptsGetMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFAXReceiptsGetTooManyRequests creates a FAXReceiptsGetTooManyRequests with default headers values
func NewFAXReceiptsGetTooManyRequests() *FAXReceiptsGetTooManyRequests {
	return &FAXReceiptsGetTooManyRequests{}
}

/*FAXReceiptsGetTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type FAXReceiptsGetTooManyRequests struct {
	Payload interface{}
}

func (o *FAXReceiptsGetTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /fax/receipts][%d] faxReceiptsGetTooManyRequests  %+v", 429, o.Payload)
}

func (o *FAXReceiptsGetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFAXReceiptsGetDefault creates a FAXReceiptsGetDefault with default headers values
func NewFAXReceiptsGetDefault(code int) *FAXReceiptsGetDefault {
	return &FAXReceiptsGetDefault{
		_statusCode: code,
	}
}

/*FAXReceiptsGetDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type FAXReceiptsGetDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the Fax receipts get default response
func (o *FAXReceiptsGetDefault) Code() int {
	return o._statusCode
}

func (o *FAXReceiptsGetDefault) Error() string {
	return fmt.Sprintf("[GET /fax/receipts][%d] FaxReceiptsGet default  %+v", o._statusCode, o.Payload)
}

func (o *FAXReceiptsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
