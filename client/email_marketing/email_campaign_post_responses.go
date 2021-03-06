// Code generated by go-swagger; DO NOT EDIT.

package email_marketing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// EmailCampaignPostReader is a Reader for the EmailCampaignPost structure.
type EmailCampaignPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmailCampaignPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEmailCampaignPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewEmailCampaignPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewEmailCampaignPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewEmailCampaignPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewEmailCampaignPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewEmailCampaignPostMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewEmailCampaignPostTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewEmailCampaignPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEmailCampaignPostOK creates a EmailCampaignPostOK with default headers values
func NewEmailCampaignPostOK() *EmailCampaignPostOK {
	return &EmailCampaignPostOK{}
}

/*EmailCampaignPostOK handles this case with default header values.

SUCCESS
*/
type EmailCampaignPostOK struct {
	Payload interface{}
}

func (o *EmailCampaignPostOK) Error() string {
	return fmt.Sprintf("[POST /email-campaigns/send][%d] emailCampaignPostOK  %+v", 200, o.Payload)
}

func (o *EmailCampaignPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailCampaignPostBadRequest creates a EmailCampaignPostBadRequest with default headers values
func NewEmailCampaignPostBadRequest() *EmailCampaignPostBadRequest {
	return &EmailCampaignPostBadRequest{}
}

/*EmailCampaignPostBadRequest handles this case with default header values.

BAD_REQUEST
*/
type EmailCampaignPostBadRequest struct {
	Payload interface{}
}

func (o *EmailCampaignPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /email-campaigns/send][%d] emailCampaignPostBadRequest  %+v", 400, o.Payload)
}

func (o *EmailCampaignPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailCampaignPostUnauthorized creates a EmailCampaignPostUnauthorized with default headers values
func NewEmailCampaignPostUnauthorized() *EmailCampaignPostUnauthorized {
	return &EmailCampaignPostUnauthorized{}
}

/*EmailCampaignPostUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type EmailCampaignPostUnauthorized struct {
	Payload interface{}
}

func (o *EmailCampaignPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /email-campaigns/send][%d] emailCampaignPostUnauthorized  %+v", 401, o.Payload)
}

func (o *EmailCampaignPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailCampaignPostForbidden creates a EmailCampaignPostForbidden with default headers values
func NewEmailCampaignPostForbidden() *EmailCampaignPostForbidden {
	return &EmailCampaignPostForbidden{}
}

/*EmailCampaignPostForbidden handles this case with default header values.

FORBIDDEN
*/
type EmailCampaignPostForbidden struct {
	Payload interface{}
}

func (o *EmailCampaignPostForbidden) Error() string {
	return fmt.Sprintf("[POST /email-campaigns/send][%d] emailCampaignPostForbidden  %+v", 403, o.Payload)
}

func (o *EmailCampaignPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailCampaignPostNotFound creates a EmailCampaignPostNotFound with default headers values
func NewEmailCampaignPostNotFound() *EmailCampaignPostNotFound {
	return &EmailCampaignPostNotFound{}
}

/*EmailCampaignPostNotFound handles this case with default header values.

NOT_FOUND
*/
type EmailCampaignPostNotFound struct {
	Payload interface{}
}

func (o *EmailCampaignPostNotFound) Error() string {
	return fmt.Sprintf("[POST /email-campaigns/send][%d] emailCampaignPostNotFound  %+v", 404, o.Payload)
}

func (o *EmailCampaignPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailCampaignPostMethodNotAllowed creates a EmailCampaignPostMethodNotAllowed with default headers values
func NewEmailCampaignPostMethodNotAllowed() *EmailCampaignPostMethodNotAllowed {
	return &EmailCampaignPostMethodNotAllowed{}
}

/*EmailCampaignPostMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type EmailCampaignPostMethodNotAllowed struct {
	Payload interface{}
}

func (o *EmailCampaignPostMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /email-campaigns/send][%d] emailCampaignPostMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *EmailCampaignPostMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailCampaignPostTooManyRequests creates a EmailCampaignPostTooManyRequests with default headers values
func NewEmailCampaignPostTooManyRequests() *EmailCampaignPostTooManyRequests {
	return &EmailCampaignPostTooManyRequests{}
}

/*EmailCampaignPostTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type EmailCampaignPostTooManyRequests struct {
	Payload interface{}
}

func (o *EmailCampaignPostTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /email-campaigns/send][%d] emailCampaignPostTooManyRequests  %+v", 429, o.Payload)
}

func (o *EmailCampaignPostTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailCampaignPostDefault creates a EmailCampaignPostDefault with default headers values
func NewEmailCampaignPostDefault(code int) *EmailCampaignPostDefault {
	return &EmailCampaignPostDefault{
		_statusCode: code,
	}
}

/*EmailCampaignPostDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type EmailCampaignPostDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the email campaign post default response
func (o *EmailCampaignPostDefault) Code() int {
	return o._statusCode
}

func (o *EmailCampaignPostDefault) Error() string {
	return fmt.Sprintf("[POST /email-campaigns/send][%d] EmailCampaignPost default  %+v", o._statusCode, o.Payload)
}

func (o *EmailCampaignPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
