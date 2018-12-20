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

// EmailCampaignHistoryGetReader is a Reader for the EmailCampaignHistoryGet structure.
type EmailCampaignHistoryGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmailCampaignHistoryGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEmailCampaignHistoryGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewEmailCampaignHistoryGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewEmailCampaignHistoryGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewEmailCampaignHistoryGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewEmailCampaignHistoryGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewEmailCampaignHistoryGetMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewEmailCampaignHistoryGetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewEmailCampaignHistoryGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEmailCampaignHistoryGetOK creates a EmailCampaignHistoryGetOK with default headers values
func NewEmailCampaignHistoryGetOK() *EmailCampaignHistoryGetOK {
	return &EmailCampaignHistoryGetOK{}
}

/*EmailCampaignHistoryGetOK handles this case with default header values.

SUCCESS
*/
type EmailCampaignHistoryGetOK struct {
	Payload string
}

func (o *EmailCampaignHistoryGetOK) Error() string {
	return fmt.Sprintf("[GET /email-campaigns/{email_campaign_id}/history][%d] emailCampaignHistoryGetOK  %+v", 200, o.Payload)
}

func (o *EmailCampaignHistoryGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailCampaignHistoryGetBadRequest creates a EmailCampaignHistoryGetBadRequest with default headers values
func NewEmailCampaignHistoryGetBadRequest() *EmailCampaignHistoryGetBadRequest {
	return &EmailCampaignHistoryGetBadRequest{}
}

/*EmailCampaignHistoryGetBadRequest handles this case with default header values.

BAD_REQUEST
*/
type EmailCampaignHistoryGetBadRequest struct {
	Payload string
}

func (o *EmailCampaignHistoryGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /email-campaigns/{email_campaign_id}/history][%d] emailCampaignHistoryGetBadRequest  %+v", 400, o.Payload)
}

func (o *EmailCampaignHistoryGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailCampaignHistoryGetUnauthorized creates a EmailCampaignHistoryGetUnauthorized with default headers values
func NewEmailCampaignHistoryGetUnauthorized() *EmailCampaignHistoryGetUnauthorized {
	return &EmailCampaignHistoryGetUnauthorized{}
}

/*EmailCampaignHistoryGetUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type EmailCampaignHistoryGetUnauthorized struct {
	Payload string
}

func (o *EmailCampaignHistoryGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /email-campaigns/{email_campaign_id}/history][%d] emailCampaignHistoryGetUnauthorized  %+v", 401, o.Payload)
}

func (o *EmailCampaignHistoryGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailCampaignHistoryGetForbidden creates a EmailCampaignHistoryGetForbidden with default headers values
func NewEmailCampaignHistoryGetForbidden() *EmailCampaignHistoryGetForbidden {
	return &EmailCampaignHistoryGetForbidden{}
}

/*EmailCampaignHistoryGetForbidden handles this case with default header values.

FORBIDDEN
*/
type EmailCampaignHistoryGetForbidden struct {
	Payload string
}

func (o *EmailCampaignHistoryGetForbidden) Error() string {
	return fmt.Sprintf("[GET /email-campaigns/{email_campaign_id}/history][%d] emailCampaignHistoryGetForbidden  %+v", 403, o.Payload)
}

func (o *EmailCampaignHistoryGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailCampaignHistoryGetNotFound creates a EmailCampaignHistoryGetNotFound with default headers values
func NewEmailCampaignHistoryGetNotFound() *EmailCampaignHistoryGetNotFound {
	return &EmailCampaignHistoryGetNotFound{}
}

/*EmailCampaignHistoryGetNotFound handles this case with default header values.

NOT_FOUND
*/
type EmailCampaignHistoryGetNotFound struct {
	Payload string
}

func (o *EmailCampaignHistoryGetNotFound) Error() string {
	return fmt.Sprintf("[GET /email-campaigns/{email_campaign_id}/history][%d] emailCampaignHistoryGetNotFound  %+v", 404, o.Payload)
}

func (o *EmailCampaignHistoryGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailCampaignHistoryGetMethodNotAllowed creates a EmailCampaignHistoryGetMethodNotAllowed with default headers values
func NewEmailCampaignHistoryGetMethodNotAllowed() *EmailCampaignHistoryGetMethodNotAllowed {
	return &EmailCampaignHistoryGetMethodNotAllowed{}
}

/*EmailCampaignHistoryGetMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type EmailCampaignHistoryGetMethodNotAllowed struct {
	Payload string
}

func (o *EmailCampaignHistoryGetMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /email-campaigns/{email_campaign_id}/history][%d] emailCampaignHistoryGetMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *EmailCampaignHistoryGetMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailCampaignHistoryGetTooManyRequests creates a EmailCampaignHistoryGetTooManyRequests with default headers values
func NewEmailCampaignHistoryGetTooManyRequests() *EmailCampaignHistoryGetTooManyRequests {
	return &EmailCampaignHistoryGetTooManyRequests{}
}

/*EmailCampaignHistoryGetTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type EmailCampaignHistoryGetTooManyRequests struct {
	Payload string
}

func (o *EmailCampaignHistoryGetTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /email-campaigns/{email_campaign_id}/history][%d] emailCampaignHistoryGetTooManyRequests  %+v", 429, o.Payload)
}

func (o *EmailCampaignHistoryGetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailCampaignHistoryGetDefault creates a EmailCampaignHistoryGetDefault with default headers values
func NewEmailCampaignHistoryGetDefault(code int) *EmailCampaignHistoryGetDefault {
	return &EmailCampaignHistoryGetDefault{
		_statusCode: code,
	}
}

/*EmailCampaignHistoryGetDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type EmailCampaignHistoryGetDefault struct {
	_statusCode int

	Payload string
}

// Code gets the status code for the email campaign history get default response
func (o *EmailCampaignHistoryGetDefault) Code() int {
	return o._statusCode
}

func (o *EmailCampaignHistoryGetDefault) Error() string {
	return fmt.Sprintf("[GET /email-campaigns/{email_campaign_id}/history][%d] EmailCampaignHistoryGet default  %+v", o._statusCode, o.Payload)
}

func (o *EmailCampaignHistoryGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
