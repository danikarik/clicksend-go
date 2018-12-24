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

// NumbersSearchByCountryGetReader is a Reader for the NumbersSearchByCountryGet structure.
type NumbersSearchByCountryGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NumbersSearchByCountryGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewNumbersSearchByCountryGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewNumbersSearchByCountryGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewNumbersSearchByCountryGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewNumbersSearchByCountryGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewNumbersSearchByCountryGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewNumbersSearchByCountryGetMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewNumbersSearchByCountryGetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewNumbersSearchByCountryGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNumbersSearchByCountryGetOK creates a NumbersSearchByCountryGetOK with default headers values
func NewNumbersSearchByCountryGetOK() *NumbersSearchByCountryGetOK {
	return &NumbersSearchByCountryGetOK{}
}

/*NumbersSearchByCountryGetOK handles this case with default header values.

SUCCESS
*/
type NumbersSearchByCountryGetOK struct {
	Payload interface{}
}

func (o *NumbersSearchByCountryGetOK) Error() string {
	return fmt.Sprintf("[GET /numbers/search/{country}][%d] numbersSearchByCountryGetOK  %+v", 200, o.Payload)
}

func (o *NumbersSearchByCountryGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNumbersSearchByCountryGetBadRequest creates a NumbersSearchByCountryGetBadRequest with default headers values
func NewNumbersSearchByCountryGetBadRequest() *NumbersSearchByCountryGetBadRequest {
	return &NumbersSearchByCountryGetBadRequest{}
}

/*NumbersSearchByCountryGetBadRequest handles this case with default header values.

BAD_REQUEST
*/
type NumbersSearchByCountryGetBadRequest struct {
	Payload interface{}
}

func (o *NumbersSearchByCountryGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /numbers/search/{country}][%d] numbersSearchByCountryGetBadRequest  %+v", 400, o.Payload)
}

func (o *NumbersSearchByCountryGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNumbersSearchByCountryGetUnauthorized creates a NumbersSearchByCountryGetUnauthorized with default headers values
func NewNumbersSearchByCountryGetUnauthorized() *NumbersSearchByCountryGetUnauthorized {
	return &NumbersSearchByCountryGetUnauthorized{}
}

/*NumbersSearchByCountryGetUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type NumbersSearchByCountryGetUnauthorized struct {
	Payload interface{}
}

func (o *NumbersSearchByCountryGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /numbers/search/{country}][%d] numbersSearchByCountryGetUnauthorized  %+v", 401, o.Payload)
}

func (o *NumbersSearchByCountryGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNumbersSearchByCountryGetForbidden creates a NumbersSearchByCountryGetForbidden with default headers values
func NewNumbersSearchByCountryGetForbidden() *NumbersSearchByCountryGetForbidden {
	return &NumbersSearchByCountryGetForbidden{}
}

/*NumbersSearchByCountryGetForbidden handles this case with default header values.

FORBIDDEN
*/
type NumbersSearchByCountryGetForbidden struct {
	Payload interface{}
}

func (o *NumbersSearchByCountryGetForbidden) Error() string {
	return fmt.Sprintf("[GET /numbers/search/{country}][%d] numbersSearchByCountryGetForbidden  %+v", 403, o.Payload)
}

func (o *NumbersSearchByCountryGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNumbersSearchByCountryGetNotFound creates a NumbersSearchByCountryGetNotFound with default headers values
func NewNumbersSearchByCountryGetNotFound() *NumbersSearchByCountryGetNotFound {
	return &NumbersSearchByCountryGetNotFound{}
}

/*NumbersSearchByCountryGetNotFound handles this case with default header values.

NOT_FOUND
*/
type NumbersSearchByCountryGetNotFound struct {
	Payload interface{}
}

func (o *NumbersSearchByCountryGetNotFound) Error() string {
	return fmt.Sprintf("[GET /numbers/search/{country}][%d] numbersSearchByCountryGetNotFound  %+v", 404, o.Payload)
}

func (o *NumbersSearchByCountryGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNumbersSearchByCountryGetMethodNotAllowed creates a NumbersSearchByCountryGetMethodNotAllowed with default headers values
func NewNumbersSearchByCountryGetMethodNotAllowed() *NumbersSearchByCountryGetMethodNotAllowed {
	return &NumbersSearchByCountryGetMethodNotAllowed{}
}

/*NumbersSearchByCountryGetMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type NumbersSearchByCountryGetMethodNotAllowed struct {
	Payload interface{}
}

func (o *NumbersSearchByCountryGetMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /numbers/search/{country}][%d] numbersSearchByCountryGetMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *NumbersSearchByCountryGetMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNumbersSearchByCountryGetTooManyRequests creates a NumbersSearchByCountryGetTooManyRequests with default headers values
func NewNumbersSearchByCountryGetTooManyRequests() *NumbersSearchByCountryGetTooManyRequests {
	return &NumbersSearchByCountryGetTooManyRequests{}
}

/*NumbersSearchByCountryGetTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type NumbersSearchByCountryGetTooManyRequests struct {
	Payload interface{}
}

func (o *NumbersSearchByCountryGetTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /numbers/search/{country}][%d] numbersSearchByCountryGetTooManyRequests  %+v", 429, o.Payload)
}

func (o *NumbersSearchByCountryGetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNumbersSearchByCountryGetDefault creates a NumbersSearchByCountryGetDefault with default headers values
func NewNumbersSearchByCountryGetDefault(code int) *NumbersSearchByCountryGetDefault {
	return &NumbersSearchByCountryGetDefault{
		_statusCode: code,
	}
}

/*NumbersSearchByCountryGetDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type NumbersSearchByCountryGetDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the numbers search by country get default response
func (o *NumbersSearchByCountryGetDefault) Code() int {
	return o._statusCode
}

func (o *NumbersSearchByCountryGetDefault) Error() string {
	return fmt.Sprintf("[GET /numbers/search/{country}][%d] NumbersSearchByCountryGet default  %+v", o._statusCode, o.Payload)
}

func (o *NumbersSearchByCountryGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
