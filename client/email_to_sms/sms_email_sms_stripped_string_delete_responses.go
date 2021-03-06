// Code generated by go-swagger; DO NOT EDIT.

package email_to_sms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// SMSEmailSMSStrippedStringDeleteReader is a Reader for the SMSEmailSMSStrippedStringDelete structure.
type SMSEmailSMSStrippedStringDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SMSEmailSMSStrippedStringDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSMSEmailSMSStrippedStringDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSMSEmailSMSStrippedStringDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSMSEmailSMSStrippedStringDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSMSEmailSMSStrippedStringDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSMSEmailSMSStrippedStringDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewSMSEmailSMSStrippedStringDeleteMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewSMSEmailSMSStrippedStringDeleteTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSMSEmailSMSStrippedStringDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSMSEmailSMSStrippedStringDeleteOK creates a SMSEmailSMSStrippedStringDeleteOK with default headers values
func NewSMSEmailSMSStrippedStringDeleteOK() *SMSEmailSMSStrippedStringDeleteOK {
	return &SMSEmailSMSStrippedStringDeleteOK{}
}

/*SMSEmailSMSStrippedStringDeleteOK handles this case with default header values.

SUCCESS
*/
type SMSEmailSMSStrippedStringDeleteOK struct {
	Payload interface{}
}

func (o *SMSEmailSMSStrippedStringDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /sms/email-sms-stripped-strings/{rule_id}][%d] smsEmailSmsStrippedStringDeleteOK  %+v", 200, o.Payload)
}

func (o *SMSEmailSMSStrippedStringDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSEmailSMSStrippedStringDeleteBadRequest creates a SMSEmailSMSStrippedStringDeleteBadRequest with default headers values
func NewSMSEmailSMSStrippedStringDeleteBadRequest() *SMSEmailSMSStrippedStringDeleteBadRequest {
	return &SMSEmailSMSStrippedStringDeleteBadRequest{}
}

/*SMSEmailSMSStrippedStringDeleteBadRequest handles this case with default header values.

BAD_REQUEST
*/
type SMSEmailSMSStrippedStringDeleteBadRequest struct {
	Payload interface{}
}

func (o *SMSEmailSMSStrippedStringDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /sms/email-sms-stripped-strings/{rule_id}][%d] smsEmailSmsStrippedStringDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *SMSEmailSMSStrippedStringDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSEmailSMSStrippedStringDeleteUnauthorized creates a SMSEmailSMSStrippedStringDeleteUnauthorized with default headers values
func NewSMSEmailSMSStrippedStringDeleteUnauthorized() *SMSEmailSMSStrippedStringDeleteUnauthorized {
	return &SMSEmailSMSStrippedStringDeleteUnauthorized{}
}

/*SMSEmailSMSStrippedStringDeleteUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type SMSEmailSMSStrippedStringDeleteUnauthorized struct {
	Payload interface{}
}

func (o *SMSEmailSMSStrippedStringDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /sms/email-sms-stripped-strings/{rule_id}][%d] smsEmailSmsStrippedStringDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *SMSEmailSMSStrippedStringDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSEmailSMSStrippedStringDeleteForbidden creates a SMSEmailSMSStrippedStringDeleteForbidden with default headers values
func NewSMSEmailSMSStrippedStringDeleteForbidden() *SMSEmailSMSStrippedStringDeleteForbidden {
	return &SMSEmailSMSStrippedStringDeleteForbidden{}
}

/*SMSEmailSMSStrippedStringDeleteForbidden handles this case with default header values.

FORBIDDEN
*/
type SMSEmailSMSStrippedStringDeleteForbidden struct {
	Payload interface{}
}

func (o *SMSEmailSMSStrippedStringDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /sms/email-sms-stripped-strings/{rule_id}][%d] smsEmailSmsStrippedStringDeleteForbidden  %+v", 403, o.Payload)
}

func (o *SMSEmailSMSStrippedStringDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSEmailSMSStrippedStringDeleteNotFound creates a SMSEmailSMSStrippedStringDeleteNotFound with default headers values
func NewSMSEmailSMSStrippedStringDeleteNotFound() *SMSEmailSMSStrippedStringDeleteNotFound {
	return &SMSEmailSMSStrippedStringDeleteNotFound{}
}

/*SMSEmailSMSStrippedStringDeleteNotFound handles this case with default header values.

NOT_FOUND
*/
type SMSEmailSMSStrippedStringDeleteNotFound struct {
	Payload interface{}
}

func (o *SMSEmailSMSStrippedStringDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /sms/email-sms-stripped-strings/{rule_id}][%d] smsEmailSmsStrippedStringDeleteNotFound  %+v", 404, o.Payload)
}

func (o *SMSEmailSMSStrippedStringDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSEmailSMSStrippedStringDeleteMethodNotAllowed creates a SMSEmailSMSStrippedStringDeleteMethodNotAllowed with default headers values
func NewSMSEmailSMSStrippedStringDeleteMethodNotAllowed() *SMSEmailSMSStrippedStringDeleteMethodNotAllowed {
	return &SMSEmailSMSStrippedStringDeleteMethodNotAllowed{}
}

/*SMSEmailSMSStrippedStringDeleteMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type SMSEmailSMSStrippedStringDeleteMethodNotAllowed struct {
	Payload interface{}
}

func (o *SMSEmailSMSStrippedStringDeleteMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /sms/email-sms-stripped-strings/{rule_id}][%d] smsEmailSmsStrippedStringDeleteMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *SMSEmailSMSStrippedStringDeleteMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSEmailSMSStrippedStringDeleteTooManyRequests creates a SMSEmailSMSStrippedStringDeleteTooManyRequests with default headers values
func NewSMSEmailSMSStrippedStringDeleteTooManyRequests() *SMSEmailSMSStrippedStringDeleteTooManyRequests {
	return &SMSEmailSMSStrippedStringDeleteTooManyRequests{}
}

/*SMSEmailSMSStrippedStringDeleteTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type SMSEmailSMSStrippedStringDeleteTooManyRequests struct {
	Payload interface{}
}

func (o *SMSEmailSMSStrippedStringDeleteTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /sms/email-sms-stripped-strings/{rule_id}][%d] smsEmailSmsStrippedStringDeleteTooManyRequests  %+v", 429, o.Payload)
}

func (o *SMSEmailSMSStrippedStringDeleteTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSEmailSMSStrippedStringDeleteDefault creates a SMSEmailSMSStrippedStringDeleteDefault with default headers values
func NewSMSEmailSMSStrippedStringDeleteDefault(code int) *SMSEmailSMSStrippedStringDeleteDefault {
	return &SMSEmailSMSStrippedStringDeleteDefault{
		_statusCode: code,
	}
}

/*SMSEmailSMSStrippedStringDeleteDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type SMSEmailSMSStrippedStringDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the Sms email Sms stripped string delete default response
func (o *SMSEmailSMSStrippedStringDeleteDefault) Code() int {
	return o._statusCode
}

func (o *SMSEmailSMSStrippedStringDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /sms/email-sms-stripped-strings/{rule_id}][%d] SmsEmailSmsStrippedStringDelete default  %+v", o._statusCode, o.Payload)
}

func (o *SMSEmailSMSStrippedStringDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
