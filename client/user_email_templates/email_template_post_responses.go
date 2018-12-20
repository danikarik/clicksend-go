// Code generated by go-swagger; DO NOT EDIT.

package user_email_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// EmailTemplatePostReader is a Reader for the EmailTemplatePost structure.
type EmailTemplatePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmailTemplatePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEmailTemplatePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewEmailTemplatePostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewEmailTemplatePostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewEmailTemplatePostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewEmailTemplatePostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewEmailTemplatePostMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewEmailTemplatePostTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewEmailTemplatePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEmailTemplatePostOK creates a EmailTemplatePostOK with default headers values
func NewEmailTemplatePostOK() *EmailTemplatePostOK {
	return &EmailTemplatePostOK{}
}

/*EmailTemplatePostOK handles this case with default header values.

SUCCESS
*/
type EmailTemplatePostOK struct {
	Payload string
}

func (o *EmailTemplatePostOK) Error() string {
	return fmt.Sprintf("[POST /email/templates][%d] emailTemplatePostOK  %+v", 200, o.Payload)
}

func (o *EmailTemplatePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailTemplatePostBadRequest creates a EmailTemplatePostBadRequest with default headers values
func NewEmailTemplatePostBadRequest() *EmailTemplatePostBadRequest {
	return &EmailTemplatePostBadRequest{}
}

/*EmailTemplatePostBadRequest handles this case with default header values.

BAD_REQUEST
*/
type EmailTemplatePostBadRequest struct {
	Payload string
}

func (o *EmailTemplatePostBadRequest) Error() string {
	return fmt.Sprintf("[POST /email/templates][%d] emailTemplatePostBadRequest  %+v", 400, o.Payload)
}

func (o *EmailTemplatePostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailTemplatePostUnauthorized creates a EmailTemplatePostUnauthorized with default headers values
func NewEmailTemplatePostUnauthorized() *EmailTemplatePostUnauthorized {
	return &EmailTemplatePostUnauthorized{}
}

/*EmailTemplatePostUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type EmailTemplatePostUnauthorized struct {
	Payload string
}

func (o *EmailTemplatePostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /email/templates][%d] emailTemplatePostUnauthorized  %+v", 401, o.Payload)
}

func (o *EmailTemplatePostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailTemplatePostForbidden creates a EmailTemplatePostForbidden with default headers values
func NewEmailTemplatePostForbidden() *EmailTemplatePostForbidden {
	return &EmailTemplatePostForbidden{}
}

/*EmailTemplatePostForbidden handles this case with default header values.

FORBIDDEN
*/
type EmailTemplatePostForbidden struct {
	Payload string
}

func (o *EmailTemplatePostForbidden) Error() string {
	return fmt.Sprintf("[POST /email/templates][%d] emailTemplatePostForbidden  %+v", 403, o.Payload)
}

func (o *EmailTemplatePostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailTemplatePostNotFound creates a EmailTemplatePostNotFound with default headers values
func NewEmailTemplatePostNotFound() *EmailTemplatePostNotFound {
	return &EmailTemplatePostNotFound{}
}

/*EmailTemplatePostNotFound handles this case with default header values.

NOT_FOUND
*/
type EmailTemplatePostNotFound struct {
	Payload string
}

func (o *EmailTemplatePostNotFound) Error() string {
	return fmt.Sprintf("[POST /email/templates][%d] emailTemplatePostNotFound  %+v", 404, o.Payload)
}

func (o *EmailTemplatePostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailTemplatePostMethodNotAllowed creates a EmailTemplatePostMethodNotAllowed with default headers values
func NewEmailTemplatePostMethodNotAllowed() *EmailTemplatePostMethodNotAllowed {
	return &EmailTemplatePostMethodNotAllowed{}
}

/*EmailTemplatePostMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type EmailTemplatePostMethodNotAllowed struct {
	Payload string
}

func (o *EmailTemplatePostMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /email/templates][%d] emailTemplatePostMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *EmailTemplatePostMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailTemplatePostTooManyRequests creates a EmailTemplatePostTooManyRequests with default headers values
func NewEmailTemplatePostTooManyRequests() *EmailTemplatePostTooManyRequests {
	return &EmailTemplatePostTooManyRequests{}
}

/*EmailTemplatePostTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type EmailTemplatePostTooManyRequests struct {
	Payload string
}

func (o *EmailTemplatePostTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /email/templates][%d] emailTemplatePostTooManyRequests  %+v", 429, o.Payload)
}

func (o *EmailTemplatePostTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailTemplatePostDefault creates a EmailTemplatePostDefault with default headers values
func NewEmailTemplatePostDefault(code int) *EmailTemplatePostDefault {
	return &EmailTemplatePostDefault{
		_statusCode: code,
	}
}

/*EmailTemplatePostDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type EmailTemplatePostDefault struct {
	_statusCode int

	Payload string
}

// Code gets the status code for the email template post default response
func (o *EmailTemplatePostDefault) Code() int {
	return o._statusCode
}

func (o *EmailTemplatePostDefault) Error() string {
	return fmt.Sprintf("[POST /email/templates][%d] EmailTemplatePost default  %+v", o._statusCode, o.Payload)
}

func (o *EmailTemplatePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}