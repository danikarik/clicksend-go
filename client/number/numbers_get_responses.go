// Code generated by go-swagger; DO NOT EDIT.

package number

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NumbersGetReader is a Reader for the NumbersGet structure.
type NumbersGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NumbersGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewNumbersGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewNumbersGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewNumbersGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewNumbersGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewNumbersGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewNumbersGetMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewNumbersGetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewNumbersGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNumbersGetOK creates a NumbersGetOK with default headers values
func NewNumbersGetOK() *NumbersGetOK {
	return &NumbersGetOK{}
}

/*NumbersGetOK handles this case with default header values.

SUCCESS
*/
type NumbersGetOK struct {
	Payload string
}

func (o *NumbersGetOK) Error() string {
	return fmt.Sprintf("[GET /numbers][%d] numbersGetOK  %+v", 200, o.Payload)
}

func (o *NumbersGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNumbersGetBadRequest creates a NumbersGetBadRequest with default headers values
func NewNumbersGetBadRequest() *NumbersGetBadRequest {
	return &NumbersGetBadRequest{}
}

/*NumbersGetBadRequest handles this case with default header values.

BAD_REQUEST
*/
type NumbersGetBadRequest struct {
	Payload string
}

func (o *NumbersGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /numbers][%d] numbersGetBadRequest  %+v", 400, o.Payload)
}

func (o *NumbersGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNumbersGetUnauthorized creates a NumbersGetUnauthorized with default headers values
func NewNumbersGetUnauthorized() *NumbersGetUnauthorized {
	return &NumbersGetUnauthorized{}
}

/*NumbersGetUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type NumbersGetUnauthorized struct {
	Payload string
}

func (o *NumbersGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /numbers][%d] numbersGetUnauthorized  %+v", 401, o.Payload)
}

func (o *NumbersGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNumbersGetForbidden creates a NumbersGetForbidden with default headers values
func NewNumbersGetForbidden() *NumbersGetForbidden {
	return &NumbersGetForbidden{}
}

/*NumbersGetForbidden handles this case with default header values.

FORBIDDEN
*/
type NumbersGetForbidden struct {
	Payload string
}

func (o *NumbersGetForbidden) Error() string {
	return fmt.Sprintf("[GET /numbers][%d] numbersGetForbidden  %+v", 403, o.Payload)
}

func (o *NumbersGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNumbersGetNotFound creates a NumbersGetNotFound with default headers values
func NewNumbersGetNotFound() *NumbersGetNotFound {
	return &NumbersGetNotFound{}
}

/*NumbersGetNotFound handles this case with default header values.

NOT_FOUND
*/
type NumbersGetNotFound struct {
	Payload string
}

func (o *NumbersGetNotFound) Error() string {
	return fmt.Sprintf("[GET /numbers][%d] numbersGetNotFound  %+v", 404, o.Payload)
}

func (o *NumbersGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNumbersGetMethodNotAllowed creates a NumbersGetMethodNotAllowed with default headers values
func NewNumbersGetMethodNotAllowed() *NumbersGetMethodNotAllowed {
	return &NumbersGetMethodNotAllowed{}
}

/*NumbersGetMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type NumbersGetMethodNotAllowed struct {
	Payload string
}

func (o *NumbersGetMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /numbers][%d] numbersGetMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *NumbersGetMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNumbersGetTooManyRequests creates a NumbersGetTooManyRequests with default headers values
func NewNumbersGetTooManyRequests() *NumbersGetTooManyRequests {
	return &NumbersGetTooManyRequests{}
}

/*NumbersGetTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type NumbersGetTooManyRequests struct {
	Payload string
}

func (o *NumbersGetTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /numbers][%d] numbersGetTooManyRequests  %+v", 429, o.Payload)
}

func (o *NumbersGetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNumbersGetDefault creates a NumbersGetDefault with default headers values
func NewNumbersGetDefault(code int) *NumbersGetDefault {
	return &NumbersGetDefault{
		_statusCode: code,
	}
}

/*NumbersGetDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type NumbersGetDefault struct {
	_statusCode int

	Payload string
}

// Code gets the status code for the numbers get default response
func (o *NumbersGetDefault) Code() int {
	return o._statusCode
}

func (o *NumbersGetDefault) Error() string {
	return fmt.Sprintf("[GET /numbers][%d] NumbersGet default  %+v", o._statusCode, o.Payload)
}

func (o *NumbersGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
