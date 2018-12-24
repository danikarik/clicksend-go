// Code generated by go-swagger; DO NOT EDIT.

package sms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// SMSCancelByMessageIDPutReader is a Reader for the SMSCancelByMessageIDPut structure.
type SMSCancelByMessageIDPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SMSCancelByMessageIDPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSMSCancelByMessageIDPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSMSCancelByMessageIDPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSMSCancelByMessageIDPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSMSCancelByMessageIDPutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSMSCancelByMessageIDPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewSMSCancelByMessageIDPutMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewSMSCancelByMessageIDPutTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSMSCancelByMessageIDPutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSMSCancelByMessageIDPutOK creates a SMSCancelByMessageIDPutOK with default headers values
func NewSMSCancelByMessageIDPutOK() *SMSCancelByMessageIDPutOK {
	return &SMSCancelByMessageIDPutOK{}
}

/*SMSCancelByMessageIDPutOK handles this case with default header values.

SUCCESS
*/
type SMSCancelByMessageIDPutOK struct {
	Payload interface{}
}

func (o *SMSCancelByMessageIDPutOK) Error() string {
	return fmt.Sprintf("[PUT /sms/{message_id}/cancel][%d] smsCancelByMessageIdPutOK  %+v", 200, o.Payload)
}

func (o *SMSCancelByMessageIDPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSCancelByMessageIDPutBadRequest creates a SMSCancelByMessageIDPutBadRequest with default headers values
func NewSMSCancelByMessageIDPutBadRequest() *SMSCancelByMessageIDPutBadRequest {
	return &SMSCancelByMessageIDPutBadRequest{}
}

/*SMSCancelByMessageIDPutBadRequest handles this case with default header values.

BAD_REQUEST
*/
type SMSCancelByMessageIDPutBadRequest struct {
	Payload interface{}
}

func (o *SMSCancelByMessageIDPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /sms/{message_id}/cancel][%d] smsCancelByMessageIdPutBadRequest  %+v", 400, o.Payload)
}

func (o *SMSCancelByMessageIDPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSCancelByMessageIDPutUnauthorized creates a SMSCancelByMessageIDPutUnauthorized with default headers values
func NewSMSCancelByMessageIDPutUnauthorized() *SMSCancelByMessageIDPutUnauthorized {
	return &SMSCancelByMessageIDPutUnauthorized{}
}

/*SMSCancelByMessageIDPutUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type SMSCancelByMessageIDPutUnauthorized struct {
	Payload interface{}
}

func (o *SMSCancelByMessageIDPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /sms/{message_id}/cancel][%d] smsCancelByMessageIdPutUnauthorized  %+v", 401, o.Payload)
}

func (o *SMSCancelByMessageIDPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSCancelByMessageIDPutForbidden creates a SMSCancelByMessageIDPutForbidden with default headers values
func NewSMSCancelByMessageIDPutForbidden() *SMSCancelByMessageIDPutForbidden {
	return &SMSCancelByMessageIDPutForbidden{}
}

/*SMSCancelByMessageIDPutForbidden handles this case with default header values.

FORBIDDEN
*/
type SMSCancelByMessageIDPutForbidden struct {
	Payload interface{}
}

func (o *SMSCancelByMessageIDPutForbidden) Error() string {
	return fmt.Sprintf("[PUT /sms/{message_id}/cancel][%d] smsCancelByMessageIdPutForbidden  %+v", 403, o.Payload)
}

func (o *SMSCancelByMessageIDPutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSCancelByMessageIDPutNotFound creates a SMSCancelByMessageIDPutNotFound with default headers values
func NewSMSCancelByMessageIDPutNotFound() *SMSCancelByMessageIDPutNotFound {
	return &SMSCancelByMessageIDPutNotFound{}
}

/*SMSCancelByMessageIDPutNotFound handles this case with default header values.

NOT_FOUND
*/
type SMSCancelByMessageIDPutNotFound struct {
	Payload interface{}
}

func (o *SMSCancelByMessageIDPutNotFound) Error() string {
	return fmt.Sprintf("[PUT /sms/{message_id}/cancel][%d] smsCancelByMessageIdPutNotFound  %+v", 404, o.Payload)
}

func (o *SMSCancelByMessageIDPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSCancelByMessageIDPutMethodNotAllowed creates a SMSCancelByMessageIDPutMethodNotAllowed with default headers values
func NewSMSCancelByMessageIDPutMethodNotAllowed() *SMSCancelByMessageIDPutMethodNotAllowed {
	return &SMSCancelByMessageIDPutMethodNotAllowed{}
}

/*SMSCancelByMessageIDPutMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type SMSCancelByMessageIDPutMethodNotAllowed struct {
	Payload interface{}
}

func (o *SMSCancelByMessageIDPutMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /sms/{message_id}/cancel][%d] smsCancelByMessageIdPutMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *SMSCancelByMessageIDPutMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSCancelByMessageIDPutTooManyRequests creates a SMSCancelByMessageIDPutTooManyRequests with default headers values
func NewSMSCancelByMessageIDPutTooManyRequests() *SMSCancelByMessageIDPutTooManyRequests {
	return &SMSCancelByMessageIDPutTooManyRequests{}
}

/*SMSCancelByMessageIDPutTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type SMSCancelByMessageIDPutTooManyRequests struct {
	Payload interface{}
}

func (o *SMSCancelByMessageIDPutTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /sms/{message_id}/cancel][%d] smsCancelByMessageIdPutTooManyRequests  %+v", 429, o.Payload)
}

func (o *SMSCancelByMessageIDPutTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSCancelByMessageIDPutDefault creates a SMSCancelByMessageIDPutDefault with default headers values
func NewSMSCancelByMessageIDPutDefault(code int) *SMSCancelByMessageIDPutDefault {
	return &SMSCancelByMessageIDPutDefault{
		_statusCode: code,
	}
}

/*SMSCancelByMessageIDPutDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type SMSCancelByMessageIDPutDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the Sms cancel by message Id put default response
func (o *SMSCancelByMessageIDPutDefault) Code() int {
	return o._statusCode
}

func (o *SMSCancelByMessageIDPutDefault) Error() string {
	return fmt.Sprintf("[PUT /sms/{message_id}/cancel][%d] SmsCancelByMessageIdPut default  %+v", o._statusCode, o.Payload)
}

func (o *SMSCancelByMessageIDPutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
