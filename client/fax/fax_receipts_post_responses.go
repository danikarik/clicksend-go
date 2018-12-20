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

// FAXReceiptsPostReader is a Reader for the FAXReceiptsPost structure.
type FAXReceiptsPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FAXReceiptsPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFAXReceiptsPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewFAXReceiptsPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewFAXReceiptsPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewFAXReceiptsPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewFAXReceiptsPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewFAXReceiptsPostMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewFAXReceiptsPostTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewFAXReceiptsPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFAXReceiptsPostOK creates a FAXReceiptsPostOK with default headers values
func NewFAXReceiptsPostOK() *FAXReceiptsPostOK {
	return &FAXReceiptsPostOK{}
}

/*FAXReceiptsPostOK handles this case with default header values.

SUCCESS
*/
type FAXReceiptsPostOK struct {
	Payload string
}

func (o *FAXReceiptsPostOK) Error() string {
	return fmt.Sprintf("[POST /fax/receipts][%d] faxReceiptsPostOK  %+v", 200, o.Payload)
}

func (o *FAXReceiptsPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFAXReceiptsPostBadRequest creates a FAXReceiptsPostBadRequest with default headers values
func NewFAXReceiptsPostBadRequest() *FAXReceiptsPostBadRequest {
	return &FAXReceiptsPostBadRequest{}
}

/*FAXReceiptsPostBadRequest handles this case with default header values.

BAD_REQUEST
*/
type FAXReceiptsPostBadRequest struct {
	Payload string
}

func (o *FAXReceiptsPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /fax/receipts][%d] faxReceiptsPostBadRequest  %+v", 400, o.Payload)
}

func (o *FAXReceiptsPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFAXReceiptsPostUnauthorized creates a FAXReceiptsPostUnauthorized with default headers values
func NewFAXReceiptsPostUnauthorized() *FAXReceiptsPostUnauthorized {
	return &FAXReceiptsPostUnauthorized{}
}

/*FAXReceiptsPostUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type FAXReceiptsPostUnauthorized struct {
	Payload string
}

func (o *FAXReceiptsPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fax/receipts][%d] faxReceiptsPostUnauthorized  %+v", 401, o.Payload)
}

func (o *FAXReceiptsPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFAXReceiptsPostForbidden creates a FAXReceiptsPostForbidden with default headers values
func NewFAXReceiptsPostForbidden() *FAXReceiptsPostForbidden {
	return &FAXReceiptsPostForbidden{}
}

/*FAXReceiptsPostForbidden handles this case with default header values.

FORBIDDEN
*/
type FAXReceiptsPostForbidden struct {
	Payload string
}

func (o *FAXReceiptsPostForbidden) Error() string {
	return fmt.Sprintf("[POST /fax/receipts][%d] faxReceiptsPostForbidden  %+v", 403, o.Payload)
}

func (o *FAXReceiptsPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFAXReceiptsPostNotFound creates a FAXReceiptsPostNotFound with default headers values
func NewFAXReceiptsPostNotFound() *FAXReceiptsPostNotFound {
	return &FAXReceiptsPostNotFound{}
}

/*FAXReceiptsPostNotFound handles this case with default header values.

NOT_FOUND
*/
type FAXReceiptsPostNotFound struct {
	Payload string
}

func (o *FAXReceiptsPostNotFound) Error() string {
	return fmt.Sprintf("[POST /fax/receipts][%d] faxReceiptsPostNotFound  %+v", 404, o.Payload)
}

func (o *FAXReceiptsPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFAXReceiptsPostMethodNotAllowed creates a FAXReceiptsPostMethodNotAllowed with default headers values
func NewFAXReceiptsPostMethodNotAllowed() *FAXReceiptsPostMethodNotAllowed {
	return &FAXReceiptsPostMethodNotAllowed{}
}

/*FAXReceiptsPostMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type FAXReceiptsPostMethodNotAllowed struct {
	Payload string
}

func (o *FAXReceiptsPostMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /fax/receipts][%d] faxReceiptsPostMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *FAXReceiptsPostMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFAXReceiptsPostTooManyRequests creates a FAXReceiptsPostTooManyRequests with default headers values
func NewFAXReceiptsPostTooManyRequests() *FAXReceiptsPostTooManyRequests {
	return &FAXReceiptsPostTooManyRequests{}
}

/*FAXReceiptsPostTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type FAXReceiptsPostTooManyRequests struct {
	Payload string
}

func (o *FAXReceiptsPostTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fax/receipts][%d] faxReceiptsPostTooManyRequests  %+v", 429, o.Payload)
}

func (o *FAXReceiptsPostTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFAXReceiptsPostDefault creates a FAXReceiptsPostDefault with default headers values
func NewFAXReceiptsPostDefault(code int) *FAXReceiptsPostDefault {
	return &FAXReceiptsPostDefault{
		_statusCode: code,
	}
}

/*FAXReceiptsPostDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type FAXReceiptsPostDefault struct {
	_statusCode int

	Payload string
}

// Code gets the status code for the Fax receipts post default response
func (o *FAXReceiptsPostDefault) Code() int {
	return o._statusCode
}

func (o *FAXReceiptsPostDefault) Error() string {
	return fmt.Sprintf("[POST /fax/receipts][%d] FaxReceiptsPost default  %+v", o._statusCode, o.Payload)
}

func (o *FAXReceiptsPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}