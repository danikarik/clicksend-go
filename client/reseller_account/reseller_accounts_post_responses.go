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

// ResellerAccountsPostReader is a Reader for the ResellerAccountsPost structure.
type ResellerAccountsPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResellerAccountsPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewResellerAccountsPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewResellerAccountsPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewResellerAccountsPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewResellerAccountsPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewResellerAccountsPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewResellerAccountsPostMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewResellerAccountsPostTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewResellerAccountsPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewResellerAccountsPostOK creates a ResellerAccountsPostOK with default headers values
func NewResellerAccountsPostOK() *ResellerAccountsPostOK {
	return &ResellerAccountsPostOK{}
}

/*ResellerAccountsPostOK handles this case with default header values.

SUCCESS
*/
type ResellerAccountsPostOK struct {
	Payload interface{}
}

func (o *ResellerAccountsPostOK) Error() string {
	return fmt.Sprintf("[POST /reseller/accounts][%d] resellerAccountsPostOK  %+v", 200, o.Payload)
}

func (o *ResellerAccountsPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResellerAccountsPostBadRequest creates a ResellerAccountsPostBadRequest with default headers values
func NewResellerAccountsPostBadRequest() *ResellerAccountsPostBadRequest {
	return &ResellerAccountsPostBadRequest{}
}

/*ResellerAccountsPostBadRequest handles this case with default header values.

BAD_REQUEST
*/
type ResellerAccountsPostBadRequest struct {
	Payload interface{}
}

func (o *ResellerAccountsPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /reseller/accounts][%d] resellerAccountsPostBadRequest  %+v", 400, o.Payload)
}

func (o *ResellerAccountsPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResellerAccountsPostUnauthorized creates a ResellerAccountsPostUnauthorized with default headers values
func NewResellerAccountsPostUnauthorized() *ResellerAccountsPostUnauthorized {
	return &ResellerAccountsPostUnauthorized{}
}

/*ResellerAccountsPostUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type ResellerAccountsPostUnauthorized struct {
	Payload interface{}
}

func (o *ResellerAccountsPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /reseller/accounts][%d] resellerAccountsPostUnauthorized  %+v", 401, o.Payload)
}

func (o *ResellerAccountsPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResellerAccountsPostForbidden creates a ResellerAccountsPostForbidden with default headers values
func NewResellerAccountsPostForbidden() *ResellerAccountsPostForbidden {
	return &ResellerAccountsPostForbidden{}
}

/*ResellerAccountsPostForbidden handles this case with default header values.

FORBIDDEN
*/
type ResellerAccountsPostForbidden struct {
	Payload interface{}
}

func (o *ResellerAccountsPostForbidden) Error() string {
	return fmt.Sprintf("[POST /reseller/accounts][%d] resellerAccountsPostForbidden  %+v", 403, o.Payload)
}

func (o *ResellerAccountsPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResellerAccountsPostNotFound creates a ResellerAccountsPostNotFound with default headers values
func NewResellerAccountsPostNotFound() *ResellerAccountsPostNotFound {
	return &ResellerAccountsPostNotFound{}
}

/*ResellerAccountsPostNotFound handles this case with default header values.

NOT_FOUND
*/
type ResellerAccountsPostNotFound struct {
	Payload interface{}
}

func (o *ResellerAccountsPostNotFound) Error() string {
	return fmt.Sprintf("[POST /reseller/accounts][%d] resellerAccountsPostNotFound  %+v", 404, o.Payload)
}

func (o *ResellerAccountsPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResellerAccountsPostMethodNotAllowed creates a ResellerAccountsPostMethodNotAllowed with default headers values
func NewResellerAccountsPostMethodNotAllowed() *ResellerAccountsPostMethodNotAllowed {
	return &ResellerAccountsPostMethodNotAllowed{}
}

/*ResellerAccountsPostMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type ResellerAccountsPostMethodNotAllowed struct {
	Payload interface{}
}

func (o *ResellerAccountsPostMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /reseller/accounts][%d] resellerAccountsPostMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *ResellerAccountsPostMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResellerAccountsPostTooManyRequests creates a ResellerAccountsPostTooManyRequests with default headers values
func NewResellerAccountsPostTooManyRequests() *ResellerAccountsPostTooManyRequests {
	return &ResellerAccountsPostTooManyRequests{}
}

/*ResellerAccountsPostTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type ResellerAccountsPostTooManyRequests struct {
	Payload interface{}
}

func (o *ResellerAccountsPostTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /reseller/accounts][%d] resellerAccountsPostTooManyRequests  %+v", 429, o.Payload)
}

func (o *ResellerAccountsPostTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResellerAccountsPostDefault creates a ResellerAccountsPostDefault with default headers values
func NewResellerAccountsPostDefault(code int) *ResellerAccountsPostDefault {
	return &ResellerAccountsPostDefault{
		_statusCode: code,
	}
}

/*ResellerAccountsPostDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type ResellerAccountsPostDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the reseller accounts post default response
func (o *ResellerAccountsPostDefault) Code() int {
	return o._statusCode
}

func (o *ResellerAccountsPostDefault) Error() string {
	return fmt.Sprintf("[POST /reseller/accounts][%d] ResellerAccountsPost default  %+v", o._statusCode, o.Payload)
}

func (o *ResellerAccountsPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
