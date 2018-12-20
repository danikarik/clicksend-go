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

// CancelEmailCampaignPutReader is a Reader for the CancelEmailCampaignPut structure.
type CancelEmailCampaignPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelEmailCampaignPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCancelEmailCampaignPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCancelEmailCampaignPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCancelEmailCampaignPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCancelEmailCampaignPutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCancelEmailCampaignPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewCancelEmailCampaignPutMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewCancelEmailCampaignPutTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCancelEmailCampaignPutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCancelEmailCampaignPutOK creates a CancelEmailCampaignPutOK with default headers values
func NewCancelEmailCampaignPutOK() *CancelEmailCampaignPutOK {
	return &CancelEmailCampaignPutOK{}
}

/*CancelEmailCampaignPutOK handles this case with default header values.

SUCCESS
*/
type CancelEmailCampaignPutOK struct {
	Payload string
}

func (o *CancelEmailCampaignPutOK) Error() string {
	return fmt.Sprintf("[PUT /email-campaigns/{email_campaign_id}/cancel][%d] cancelEmailCampaignPutOK  %+v", 200, o.Payload)
}

func (o *CancelEmailCampaignPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelEmailCampaignPutBadRequest creates a CancelEmailCampaignPutBadRequest with default headers values
func NewCancelEmailCampaignPutBadRequest() *CancelEmailCampaignPutBadRequest {
	return &CancelEmailCampaignPutBadRequest{}
}

/*CancelEmailCampaignPutBadRequest handles this case with default header values.

BAD_REQUEST
*/
type CancelEmailCampaignPutBadRequest struct {
	Payload string
}

func (o *CancelEmailCampaignPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /email-campaigns/{email_campaign_id}/cancel][%d] cancelEmailCampaignPutBadRequest  %+v", 400, o.Payload)
}

func (o *CancelEmailCampaignPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelEmailCampaignPutUnauthorized creates a CancelEmailCampaignPutUnauthorized with default headers values
func NewCancelEmailCampaignPutUnauthorized() *CancelEmailCampaignPutUnauthorized {
	return &CancelEmailCampaignPutUnauthorized{}
}

/*CancelEmailCampaignPutUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type CancelEmailCampaignPutUnauthorized struct {
	Payload string
}

func (o *CancelEmailCampaignPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /email-campaigns/{email_campaign_id}/cancel][%d] cancelEmailCampaignPutUnauthorized  %+v", 401, o.Payload)
}

func (o *CancelEmailCampaignPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelEmailCampaignPutForbidden creates a CancelEmailCampaignPutForbidden with default headers values
func NewCancelEmailCampaignPutForbidden() *CancelEmailCampaignPutForbidden {
	return &CancelEmailCampaignPutForbidden{}
}

/*CancelEmailCampaignPutForbidden handles this case with default header values.

FORBIDDEN
*/
type CancelEmailCampaignPutForbidden struct {
	Payload string
}

func (o *CancelEmailCampaignPutForbidden) Error() string {
	return fmt.Sprintf("[PUT /email-campaigns/{email_campaign_id}/cancel][%d] cancelEmailCampaignPutForbidden  %+v", 403, o.Payload)
}

func (o *CancelEmailCampaignPutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelEmailCampaignPutNotFound creates a CancelEmailCampaignPutNotFound with default headers values
func NewCancelEmailCampaignPutNotFound() *CancelEmailCampaignPutNotFound {
	return &CancelEmailCampaignPutNotFound{}
}

/*CancelEmailCampaignPutNotFound handles this case with default header values.

NOT_FOUND
*/
type CancelEmailCampaignPutNotFound struct {
	Payload string
}

func (o *CancelEmailCampaignPutNotFound) Error() string {
	return fmt.Sprintf("[PUT /email-campaigns/{email_campaign_id}/cancel][%d] cancelEmailCampaignPutNotFound  %+v", 404, o.Payload)
}

func (o *CancelEmailCampaignPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelEmailCampaignPutMethodNotAllowed creates a CancelEmailCampaignPutMethodNotAllowed with default headers values
func NewCancelEmailCampaignPutMethodNotAllowed() *CancelEmailCampaignPutMethodNotAllowed {
	return &CancelEmailCampaignPutMethodNotAllowed{}
}

/*CancelEmailCampaignPutMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type CancelEmailCampaignPutMethodNotAllowed struct {
	Payload string
}

func (o *CancelEmailCampaignPutMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /email-campaigns/{email_campaign_id}/cancel][%d] cancelEmailCampaignPutMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *CancelEmailCampaignPutMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelEmailCampaignPutTooManyRequests creates a CancelEmailCampaignPutTooManyRequests with default headers values
func NewCancelEmailCampaignPutTooManyRequests() *CancelEmailCampaignPutTooManyRequests {
	return &CancelEmailCampaignPutTooManyRequests{}
}

/*CancelEmailCampaignPutTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type CancelEmailCampaignPutTooManyRequests struct {
	Payload string
}

func (o *CancelEmailCampaignPutTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /email-campaigns/{email_campaign_id}/cancel][%d] cancelEmailCampaignPutTooManyRequests  %+v", 429, o.Payload)
}

func (o *CancelEmailCampaignPutTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelEmailCampaignPutDefault creates a CancelEmailCampaignPutDefault with default headers values
func NewCancelEmailCampaignPutDefault(code int) *CancelEmailCampaignPutDefault {
	return &CancelEmailCampaignPutDefault{
		_statusCode: code,
	}
}

/*CancelEmailCampaignPutDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type CancelEmailCampaignPutDefault struct {
	_statusCode int

	Payload string
}

// Code gets the status code for the cancel email campaign put default response
func (o *CancelEmailCampaignPutDefault) Code() int {
	return o._statusCode
}

func (o *CancelEmailCampaignPutDefault) Error() string {
	return fmt.Sprintf("[PUT /email-campaigns/{email_campaign_id}/cancel][%d] CancelEmailCampaignPut default  %+v", o._statusCode, o.Payload)
}

func (o *CancelEmailCampaignPutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}