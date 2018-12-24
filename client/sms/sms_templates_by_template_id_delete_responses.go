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

// SMSTemplatesByTemplateIDDeleteReader is a Reader for the SMSTemplatesByTemplateIDDelete structure.
type SMSTemplatesByTemplateIDDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SMSTemplatesByTemplateIDDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSMSTemplatesByTemplateIDDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSMSTemplatesByTemplateIDDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSMSTemplatesByTemplateIDDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSMSTemplatesByTemplateIDDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSMSTemplatesByTemplateIDDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewSMSTemplatesByTemplateIDDeleteMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewSMSTemplatesByTemplateIDDeleteTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSMSTemplatesByTemplateIDDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSMSTemplatesByTemplateIDDeleteOK creates a SMSTemplatesByTemplateIDDeleteOK with default headers values
func NewSMSTemplatesByTemplateIDDeleteOK() *SMSTemplatesByTemplateIDDeleteOK {
	return &SMSTemplatesByTemplateIDDeleteOK{}
}

/*SMSTemplatesByTemplateIDDeleteOK handles this case with default header values.

SUCCESS
*/
type SMSTemplatesByTemplateIDDeleteOK struct {
	Payload interface{}
}

func (o *SMSTemplatesByTemplateIDDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /sms/templates/{template_id}][%d] smsTemplatesByTemplateIdDeleteOK  %+v", 200, o.Payload)
}

func (o *SMSTemplatesByTemplateIDDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSTemplatesByTemplateIDDeleteBadRequest creates a SMSTemplatesByTemplateIDDeleteBadRequest with default headers values
func NewSMSTemplatesByTemplateIDDeleteBadRequest() *SMSTemplatesByTemplateIDDeleteBadRequest {
	return &SMSTemplatesByTemplateIDDeleteBadRequest{}
}

/*SMSTemplatesByTemplateIDDeleteBadRequest handles this case with default header values.

BAD_REQUEST
*/
type SMSTemplatesByTemplateIDDeleteBadRequest struct {
	Payload interface{}
}

func (o *SMSTemplatesByTemplateIDDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /sms/templates/{template_id}][%d] smsTemplatesByTemplateIdDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *SMSTemplatesByTemplateIDDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSTemplatesByTemplateIDDeleteUnauthorized creates a SMSTemplatesByTemplateIDDeleteUnauthorized with default headers values
func NewSMSTemplatesByTemplateIDDeleteUnauthorized() *SMSTemplatesByTemplateIDDeleteUnauthorized {
	return &SMSTemplatesByTemplateIDDeleteUnauthorized{}
}

/*SMSTemplatesByTemplateIDDeleteUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type SMSTemplatesByTemplateIDDeleteUnauthorized struct {
	Payload interface{}
}

func (o *SMSTemplatesByTemplateIDDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /sms/templates/{template_id}][%d] smsTemplatesByTemplateIdDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *SMSTemplatesByTemplateIDDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSTemplatesByTemplateIDDeleteForbidden creates a SMSTemplatesByTemplateIDDeleteForbidden with default headers values
func NewSMSTemplatesByTemplateIDDeleteForbidden() *SMSTemplatesByTemplateIDDeleteForbidden {
	return &SMSTemplatesByTemplateIDDeleteForbidden{}
}

/*SMSTemplatesByTemplateIDDeleteForbidden handles this case with default header values.

FORBIDDEN
*/
type SMSTemplatesByTemplateIDDeleteForbidden struct {
	Payload interface{}
}

func (o *SMSTemplatesByTemplateIDDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /sms/templates/{template_id}][%d] smsTemplatesByTemplateIdDeleteForbidden  %+v", 403, o.Payload)
}

func (o *SMSTemplatesByTemplateIDDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSTemplatesByTemplateIDDeleteNotFound creates a SMSTemplatesByTemplateIDDeleteNotFound with default headers values
func NewSMSTemplatesByTemplateIDDeleteNotFound() *SMSTemplatesByTemplateIDDeleteNotFound {
	return &SMSTemplatesByTemplateIDDeleteNotFound{}
}

/*SMSTemplatesByTemplateIDDeleteNotFound handles this case with default header values.

NOT_FOUND
*/
type SMSTemplatesByTemplateIDDeleteNotFound struct {
	Payload interface{}
}

func (o *SMSTemplatesByTemplateIDDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /sms/templates/{template_id}][%d] smsTemplatesByTemplateIdDeleteNotFound  %+v", 404, o.Payload)
}

func (o *SMSTemplatesByTemplateIDDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSTemplatesByTemplateIDDeleteMethodNotAllowed creates a SMSTemplatesByTemplateIDDeleteMethodNotAllowed with default headers values
func NewSMSTemplatesByTemplateIDDeleteMethodNotAllowed() *SMSTemplatesByTemplateIDDeleteMethodNotAllowed {
	return &SMSTemplatesByTemplateIDDeleteMethodNotAllowed{}
}

/*SMSTemplatesByTemplateIDDeleteMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type SMSTemplatesByTemplateIDDeleteMethodNotAllowed struct {
	Payload interface{}
}

func (o *SMSTemplatesByTemplateIDDeleteMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /sms/templates/{template_id}][%d] smsTemplatesByTemplateIdDeleteMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *SMSTemplatesByTemplateIDDeleteMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSTemplatesByTemplateIDDeleteTooManyRequests creates a SMSTemplatesByTemplateIDDeleteTooManyRequests with default headers values
func NewSMSTemplatesByTemplateIDDeleteTooManyRequests() *SMSTemplatesByTemplateIDDeleteTooManyRequests {
	return &SMSTemplatesByTemplateIDDeleteTooManyRequests{}
}

/*SMSTemplatesByTemplateIDDeleteTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type SMSTemplatesByTemplateIDDeleteTooManyRequests struct {
	Payload interface{}
}

func (o *SMSTemplatesByTemplateIDDeleteTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /sms/templates/{template_id}][%d] smsTemplatesByTemplateIdDeleteTooManyRequests  %+v", 429, o.Payload)
}

func (o *SMSTemplatesByTemplateIDDeleteTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSTemplatesByTemplateIDDeleteDefault creates a SMSTemplatesByTemplateIDDeleteDefault with default headers values
func NewSMSTemplatesByTemplateIDDeleteDefault(code int) *SMSTemplatesByTemplateIDDeleteDefault {
	return &SMSTemplatesByTemplateIDDeleteDefault{
		_statusCode: code,
	}
}

/*SMSTemplatesByTemplateIDDeleteDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type SMSTemplatesByTemplateIDDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the Sms templates by template Id delete default response
func (o *SMSTemplatesByTemplateIDDeleteDefault) Code() int {
	return o._statusCode
}

func (o *SMSTemplatesByTemplateIDDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /sms/templates/{template_id}][%d] SmsTemplatesByTemplateIdDelete default  %+v", o._statusCode, o.Payload)
}

func (o *SMSTemplatesByTemplateIDDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
