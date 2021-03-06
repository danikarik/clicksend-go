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

// ForgotPasswordPutReader is a Reader for the ForgotPasswordPut structure.
type ForgotPasswordPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ForgotPasswordPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewForgotPasswordPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewForgotPasswordPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewForgotPasswordPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewForgotPasswordPutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewForgotPasswordPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewForgotPasswordPutMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewForgotPasswordPutTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewForgotPasswordPutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewForgotPasswordPutOK creates a ForgotPasswordPutOK with default headers values
func NewForgotPasswordPutOK() *ForgotPasswordPutOK {
	return &ForgotPasswordPutOK{}
}

/*ForgotPasswordPutOK handles this case with default header values.

SUCCESS
*/
type ForgotPasswordPutOK struct {
	Payload interface{}
}

func (o *ForgotPasswordPutOK) Error() string {
	return fmt.Sprintf("[PUT /forgot-password][%d] forgotPasswordPutOK  %+v", 200, o.Payload)
}

func (o *ForgotPasswordPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewForgotPasswordPutBadRequest creates a ForgotPasswordPutBadRequest with default headers values
func NewForgotPasswordPutBadRequest() *ForgotPasswordPutBadRequest {
	return &ForgotPasswordPutBadRequest{}
}

/*ForgotPasswordPutBadRequest handles this case with default header values.

BAD_REQUEST
*/
type ForgotPasswordPutBadRequest struct {
	Payload interface{}
}

func (o *ForgotPasswordPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /forgot-password][%d] forgotPasswordPutBadRequest  %+v", 400, o.Payload)
}

func (o *ForgotPasswordPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewForgotPasswordPutUnauthorized creates a ForgotPasswordPutUnauthorized with default headers values
func NewForgotPasswordPutUnauthorized() *ForgotPasswordPutUnauthorized {
	return &ForgotPasswordPutUnauthorized{}
}

/*ForgotPasswordPutUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type ForgotPasswordPutUnauthorized struct {
	Payload interface{}
}

func (o *ForgotPasswordPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /forgot-password][%d] forgotPasswordPutUnauthorized  %+v", 401, o.Payload)
}

func (o *ForgotPasswordPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewForgotPasswordPutForbidden creates a ForgotPasswordPutForbidden with default headers values
func NewForgotPasswordPutForbidden() *ForgotPasswordPutForbidden {
	return &ForgotPasswordPutForbidden{}
}

/*ForgotPasswordPutForbidden handles this case with default header values.

FORBIDDEN
*/
type ForgotPasswordPutForbidden struct {
	Payload interface{}
}

func (o *ForgotPasswordPutForbidden) Error() string {
	return fmt.Sprintf("[PUT /forgot-password][%d] forgotPasswordPutForbidden  %+v", 403, o.Payload)
}

func (o *ForgotPasswordPutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewForgotPasswordPutNotFound creates a ForgotPasswordPutNotFound with default headers values
func NewForgotPasswordPutNotFound() *ForgotPasswordPutNotFound {
	return &ForgotPasswordPutNotFound{}
}

/*ForgotPasswordPutNotFound handles this case with default header values.

NOT_FOUND
*/
type ForgotPasswordPutNotFound struct {
	Payload interface{}
}

func (o *ForgotPasswordPutNotFound) Error() string {
	return fmt.Sprintf("[PUT /forgot-password][%d] forgotPasswordPutNotFound  %+v", 404, o.Payload)
}

func (o *ForgotPasswordPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewForgotPasswordPutMethodNotAllowed creates a ForgotPasswordPutMethodNotAllowed with default headers values
func NewForgotPasswordPutMethodNotAllowed() *ForgotPasswordPutMethodNotAllowed {
	return &ForgotPasswordPutMethodNotAllowed{}
}

/*ForgotPasswordPutMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type ForgotPasswordPutMethodNotAllowed struct {
	Payload interface{}
}

func (o *ForgotPasswordPutMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /forgot-password][%d] forgotPasswordPutMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *ForgotPasswordPutMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewForgotPasswordPutTooManyRequests creates a ForgotPasswordPutTooManyRequests with default headers values
func NewForgotPasswordPutTooManyRequests() *ForgotPasswordPutTooManyRequests {
	return &ForgotPasswordPutTooManyRequests{}
}

/*ForgotPasswordPutTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type ForgotPasswordPutTooManyRequests struct {
	Payload interface{}
}

func (o *ForgotPasswordPutTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /forgot-password][%d] forgotPasswordPutTooManyRequests  %+v", 429, o.Payload)
}

func (o *ForgotPasswordPutTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewForgotPasswordPutDefault creates a ForgotPasswordPutDefault with default headers values
func NewForgotPasswordPutDefault(code int) *ForgotPasswordPutDefault {
	return &ForgotPasswordPutDefault{
		_statusCode: code,
	}
}

/*ForgotPasswordPutDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type ForgotPasswordPutDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the forgot password put default response
func (o *ForgotPasswordPutDefault) Code() int {
	return o._statusCode
}

func (o *ForgotPasswordPutDefault) Error() string {
	return fmt.Sprintf("[PUT /forgot-password][%d] ForgotPasswordPut default  %+v", o._statusCode, o.Payload)
}

func (o *ForgotPasswordPutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
