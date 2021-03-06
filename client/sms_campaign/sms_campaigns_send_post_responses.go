// Code generated by go-swagger; DO NOT EDIT.

package sms_campaign

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// SMSCampaignsSendPostReader is a Reader for the SMSCampaignsSendPost structure.
type SMSCampaignsSendPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SMSCampaignsSendPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSMSCampaignsSendPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSMSCampaignsSendPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSMSCampaignsSendPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSMSCampaignsSendPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSMSCampaignsSendPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewSMSCampaignsSendPostMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewSMSCampaignsSendPostTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSMSCampaignsSendPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSMSCampaignsSendPostOK creates a SMSCampaignsSendPostOK with default headers values
func NewSMSCampaignsSendPostOK() *SMSCampaignsSendPostOK {
	return &SMSCampaignsSendPostOK{}
}

/*SMSCampaignsSendPostOK handles this case with default header values.

SUCCESS
*/
type SMSCampaignsSendPostOK struct {
	Payload interface{}
}

func (o *SMSCampaignsSendPostOK) Error() string {
	return fmt.Sprintf("[POST /sms-campaigns/send][%d] smsCampaignsSendPostOK  %+v", 200, o.Payload)
}

func (o *SMSCampaignsSendPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSCampaignsSendPostBadRequest creates a SMSCampaignsSendPostBadRequest with default headers values
func NewSMSCampaignsSendPostBadRequest() *SMSCampaignsSendPostBadRequest {
	return &SMSCampaignsSendPostBadRequest{}
}

/*SMSCampaignsSendPostBadRequest handles this case with default header values.

BAD_REQUEST
*/
type SMSCampaignsSendPostBadRequest struct {
	Payload interface{}
}

func (o *SMSCampaignsSendPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /sms-campaigns/send][%d] smsCampaignsSendPostBadRequest  %+v", 400, o.Payload)
}

func (o *SMSCampaignsSendPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSCampaignsSendPostUnauthorized creates a SMSCampaignsSendPostUnauthorized with default headers values
func NewSMSCampaignsSendPostUnauthorized() *SMSCampaignsSendPostUnauthorized {
	return &SMSCampaignsSendPostUnauthorized{}
}

/*SMSCampaignsSendPostUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type SMSCampaignsSendPostUnauthorized struct {
	Payload interface{}
}

func (o *SMSCampaignsSendPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /sms-campaigns/send][%d] smsCampaignsSendPostUnauthorized  %+v", 401, o.Payload)
}

func (o *SMSCampaignsSendPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSCampaignsSendPostForbidden creates a SMSCampaignsSendPostForbidden with default headers values
func NewSMSCampaignsSendPostForbidden() *SMSCampaignsSendPostForbidden {
	return &SMSCampaignsSendPostForbidden{}
}

/*SMSCampaignsSendPostForbidden handles this case with default header values.

FORBIDDEN
*/
type SMSCampaignsSendPostForbidden struct {
	Payload interface{}
}

func (o *SMSCampaignsSendPostForbidden) Error() string {
	return fmt.Sprintf("[POST /sms-campaigns/send][%d] smsCampaignsSendPostForbidden  %+v", 403, o.Payload)
}

func (o *SMSCampaignsSendPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSCampaignsSendPostNotFound creates a SMSCampaignsSendPostNotFound with default headers values
func NewSMSCampaignsSendPostNotFound() *SMSCampaignsSendPostNotFound {
	return &SMSCampaignsSendPostNotFound{}
}

/*SMSCampaignsSendPostNotFound handles this case with default header values.

NOT_FOUND
*/
type SMSCampaignsSendPostNotFound struct {
	Payload interface{}
}

func (o *SMSCampaignsSendPostNotFound) Error() string {
	return fmt.Sprintf("[POST /sms-campaigns/send][%d] smsCampaignsSendPostNotFound  %+v", 404, o.Payload)
}

func (o *SMSCampaignsSendPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSCampaignsSendPostMethodNotAllowed creates a SMSCampaignsSendPostMethodNotAllowed with default headers values
func NewSMSCampaignsSendPostMethodNotAllowed() *SMSCampaignsSendPostMethodNotAllowed {
	return &SMSCampaignsSendPostMethodNotAllowed{}
}

/*SMSCampaignsSendPostMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type SMSCampaignsSendPostMethodNotAllowed struct {
	Payload interface{}
}

func (o *SMSCampaignsSendPostMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /sms-campaigns/send][%d] smsCampaignsSendPostMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *SMSCampaignsSendPostMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSCampaignsSendPostTooManyRequests creates a SMSCampaignsSendPostTooManyRequests with default headers values
func NewSMSCampaignsSendPostTooManyRequests() *SMSCampaignsSendPostTooManyRequests {
	return &SMSCampaignsSendPostTooManyRequests{}
}

/*SMSCampaignsSendPostTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type SMSCampaignsSendPostTooManyRequests struct {
	Payload interface{}
}

func (o *SMSCampaignsSendPostTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /sms-campaigns/send][%d] smsCampaignsSendPostTooManyRequests  %+v", 429, o.Payload)
}

func (o *SMSCampaignsSendPostTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSCampaignsSendPostDefault creates a SMSCampaignsSendPostDefault with default headers values
func NewSMSCampaignsSendPostDefault(code int) *SMSCampaignsSendPostDefault {
	return &SMSCampaignsSendPostDefault{
		_statusCode: code,
	}
}

/*SMSCampaignsSendPostDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type SMSCampaignsSendPostDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the Sms campaigns send post default response
func (o *SMSCampaignsSendPostDefault) Code() int {
	return o._statusCode
}

func (o *SMSCampaignsSendPostDefault) Error() string {
	return fmt.Sprintf("[POST /sms-campaigns/send][%d] SmsCampaignsSendPost default  %+v", o._statusCode, o.Payload)
}

func (o *SMSCampaignsSendPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
