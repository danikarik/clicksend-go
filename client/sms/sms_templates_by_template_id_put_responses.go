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

// SMSTemplatesByTemplateIDPutReader is a Reader for the SMSTemplatesByTemplateIDPut structure.
type SMSTemplatesByTemplateIDPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SMSTemplatesByTemplateIDPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSMSTemplatesByTemplateIDPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSMSTemplatesByTemplateIDPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSMSTemplatesByTemplateIDPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSMSTemplatesByTemplateIDPutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSMSTemplatesByTemplateIDPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewSMSTemplatesByTemplateIDPutMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewSMSTemplatesByTemplateIDPutTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSMSTemplatesByTemplateIDPutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSMSTemplatesByTemplateIDPutOK creates a SMSTemplatesByTemplateIDPutOK with default headers values
func NewSMSTemplatesByTemplateIDPutOK() *SMSTemplatesByTemplateIDPutOK {
	return &SMSTemplatesByTemplateIDPutOK{}
}

/*SMSTemplatesByTemplateIDPutOK handles this case with default header values.

SUCCESS
*/
type SMSTemplatesByTemplateIDPutOK struct {
	Payload string
}

func (o *SMSTemplatesByTemplateIDPutOK) Error() string {
	return fmt.Sprintf("[PUT /sms/templates/{template_id}][%d] smsTemplatesByTemplateIdPutOK  %+v", 200, o.Payload)
}

func (o *SMSTemplatesByTemplateIDPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSTemplatesByTemplateIDPutBadRequest creates a SMSTemplatesByTemplateIDPutBadRequest with default headers values
func NewSMSTemplatesByTemplateIDPutBadRequest() *SMSTemplatesByTemplateIDPutBadRequest {
	return &SMSTemplatesByTemplateIDPutBadRequest{}
}

/*SMSTemplatesByTemplateIDPutBadRequest handles this case with default header values.

BAD_REQUEST
*/
type SMSTemplatesByTemplateIDPutBadRequest struct {
	Payload string
}

func (o *SMSTemplatesByTemplateIDPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /sms/templates/{template_id}][%d] smsTemplatesByTemplateIdPutBadRequest  %+v", 400, o.Payload)
}

func (o *SMSTemplatesByTemplateIDPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSTemplatesByTemplateIDPutUnauthorized creates a SMSTemplatesByTemplateIDPutUnauthorized with default headers values
func NewSMSTemplatesByTemplateIDPutUnauthorized() *SMSTemplatesByTemplateIDPutUnauthorized {
	return &SMSTemplatesByTemplateIDPutUnauthorized{}
}

/*SMSTemplatesByTemplateIDPutUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type SMSTemplatesByTemplateIDPutUnauthorized struct {
	Payload string
}

func (o *SMSTemplatesByTemplateIDPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /sms/templates/{template_id}][%d] smsTemplatesByTemplateIdPutUnauthorized  %+v", 401, o.Payload)
}

func (o *SMSTemplatesByTemplateIDPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSTemplatesByTemplateIDPutForbidden creates a SMSTemplatesByTemplateIDPutForbidden with default headers values
func NewSMSTemplatesByTemplateIDPutForbidden() *SMSTemplatesByTemplateIDPutForbidden {
	return &SMSTemplatesByTemplateIDPutForbidden{}
}

/*SMSTemplatesByTemplateIDPutForbidden handles this case with default header values.

FORBIDDEN
*/
type SMSTemplatesByTemplateIDPutForbidden struct {
	Payload string
}

func (o *SMSTemplatesByTemplateIDPutForbidden) Error() string {
	return fmt.Sprintf("[PUT /sms/templates/{template_id}][%d] smsTemplatesByTemplateIdPutForbidden  %+v", 403, o.Payload)
}

func (o *SMSTemplatesByTemplateIDPutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSTemplatesByTemplateIDPutNotFound creates a SMSTemplatesByTemplateIDPutNotFound with default headers values
func NewSMSTemplatesByTemplateIDPutNotFound() *SMSTemplatesByTemplateIDPutNotFound {
	return &SMSTemplatesByTemplateIDPutNotFound{}
}

/*SMSTemplatesByTemplateIDPutNotFound handles this case with default header values.

NOT_FOUND
*/
type SMSTemplatesByTemplateIDPutNotFound struct {
	Payload string
}

func (o *SMSTemplatesByTemplateIDPutNotFound) Error() string {
	return fmt.Sprintf("[PUT /sms/templates/{template_id}][%d] smsTemplatesByTemplateIdPutNotFound  %+v", 404, o.Payload)
}

func (o *SMSTemplatesByTemplateIDPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSTemplatesByTemplateIDPutMethodNotAllowed creates a SMSTemplatesByTemplateIDPutMethodNotAllowed with default headers values
func NewSMSTemplatesByTemplateIDPutMethodNotAllowed() *SMSTemplatesByTemplateIDPutMethodNotAllowed {
	return &SMSTemplatesByTemplateIDPutMethodNotAllowed{}
}

/*SMSTemplatesByTemplateIDPutMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type SMSTemplatesByTemplateIDPutMethodNotAllowed struct {
	Payload string
}

func (o *SMSTemplatesByTemplateIDPutMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /sms/templates/{template_id}][%d] smsTemplatesByTemplateIdPutMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *SMSTemplatesByTemplateIDPutMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSTemplatesByTemplateIDPutTooManyRequests creates a SMSTemplatesByTemplateIDPutTooManyRequests with default headers values
func NewSMSTemplatesByTemplateIDPutTooManyRequests() *SMSTemplatesByTemplateIDPutTooManyRequests {
	return &SMSTemplatesByTemplateIDPutTooManyRequests{}
}

/*SMSTemplatesByTemplateIDPutTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type SMSTemplatesByTemplateIDPutTooManyRequests struct {
	Payload string
}

func (o *SMSTemplatesByTemplateIDPutTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /sms/templates/{template_id}][%d] smsTemplatesByTemplateIdPutTooManyRequests  %+v", 429, o.Payload)
}

func (o *SMSTemplatesByTemplateIDPutTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSTemplatesByTemplateIDPutDefault creates a SMSTemplatesByTemplateIDPutDefault with default headers values
func NewSMSTemplatesByTemplateIDPutDefault(code int) *SMSTemplatesByTemplateIDPutDefault {
	return &SMSTemplatesByTemplateIDPutDefault{
		_statusCode: code,
	}
}

/*SMSTemplatesByTemplateIDPutDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type SMSTemplatesByTemplateIDPutDefault struct {
	_statusCode int

	Payload string
}

// Code gets the status code for the Sms templates by template Id put default response
func (o *SMSTemplatesByTemplateIDPutDefault) Code() int {
	return o._statusCode
}

func (o *SMSTemplatesByTemplateIDPutDefault) Error() string {
	return fmt.Sprintf("[PUT /sms/templates/{template_id}][%d] SmsTemplatesByTemplateIdPut default  %+v", o._statusCode, o.Payload)
}

func (o *SMSTemplatesByTemplateIDPutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}