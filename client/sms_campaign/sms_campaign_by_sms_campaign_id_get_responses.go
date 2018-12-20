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

// SMSCampaignBySMSCampaignIDGetReader is a Reader for the SMSCampaignBySMSCampaignIDGet structure.
type SMSCampaignBySMSCampaignIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SMSCampaignBySMSCampaignIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSMSCampaignBySMSCampaignIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSMSCampaignBySMSCampaignIDGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSMSCampaignBySMSCampaignIDGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSMSCampaignBySMSCampaignIDGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSMSCampaignBySMSCampaignIDGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewSMSCampaignBySMSCampaignIDGetMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewSMSCampaignBySMSCampaignIDGetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSMSCampaignBySMSCampaignIDGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSMSCampaignBySMSCampaignIDGetOK creates a SMSCampaignBySMSCampaignIDGetOK with default headers values
func NewSMSCampaignBySMSCampaignIDGetOK() *SMSCampaignBySMSCampaignIDGetOK {
	return &SMSCampaignBySMSCampaignIDGetOK{}
}

/*SMSCampaignBySMSCampaignIDGetOK handles this case with default header values.

SUCCESS
*/
type SMSCampaignBySMSCampaignIDGetOK struct {
	Payload string
}

func (o *SMSCampaignBySMSCampaignIDGetOK) Error() string {
	return fmt.Sprintf("[GET /sms-campaigns/{sms_campaign_id}][%d] smsCampaignBySmsCampaignIdGetOK  %+v", 200, o.Payload)
}

func (o *SMSCampaignBySMSCampaignIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSCampaignBySMSCampaignIDGetBadRequest creates a SMSCampaignBySMSCampaignIDGetBadRequest with default headers values
func NewSMSCampaignBySMSCampaignIDGetBadRequest() *SMSCampaignBySMSCampaignIDGetBadRequest {
	return &SMSCampaignBySMSCampaignIDGetBadRequest{}
}

/*SMSCampaignBySMSCampaignIDGetBadRequest handles this case with default header values.

BAD_REQUEST
*/
type SMSCampaignBySMSCampaignIDGetBadRequest struct {
	Payload string
}

func (o *SMSCampaignBySMSCampaignIDGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /sms-campaigns/{sms_campaign_id}][%d] smsCampaignBySmsCampaignIdGetBadRequest  %+v", 400, o.Payload)
}

func (o *SMSCampaignBySMSCampaignIDGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSCampaignBySMSCampaignIDGetUnauthorized creates a SMSCampaignBySMSCampaignIDGetUnauthorized with default headers values
func NewSMSCampaignBySMSCampaignIDGetUnauthorized() *SMSCampaignBySMSCampaignIDGetUnauthorized {
	return &SMSCampaignBySMSCampaignIDGetUnauthorized{}
}

/*SMSCampaignBySMSCampaignIDGetUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type SMSCampaignBySMSCampaignIDGetUnauthorized struct {
	Payload string
}

func (o *SMSCampaignBySMSCampaignIDGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /sms-campaigns/{sms_campaign_id}][%d] smsCampaignBySmsCampaignIdGetUnauthorized  %+v", 401, o.Payload)
}

func (o *SMSCampaignBySMSCampaignIDGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSCampaignBySMSCampaignIDGetForbidden creates a SMSCampaignBySMSCampaignIDGetForbidden with default headers values
func NewSMSCampaignBySMSCampaignIDGetForbidden() *SMSCampaignBySMSCampaignIDGetForbidden {
	return &SMSCampaignBySMSCampaignIDGetForbidden{}
}

/*SMSCampaignBySMSCampaignIDGetForbidden handles this case with default header values.

FORBIDDEN
*/
type SMSCampaignBySMSCampaignIDGetForbidden struct {
	Payload string
}

func (o *SMSCampaignBySMSCampaignIDGetForbidden) Error() string {
	return fmt.Sprintf("[GET /sms-campaigns/{sms_campaign_id}][%d] smsCampaignBySmsCampaignIdGetForbidden  %+v", 403, o.Payload)
}

func (o *SMSCampaignBySMSCampaignIDGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSCampaignBySMSCampaignIDGetNotFound creates a SMSCampaignBySMSCampaignIDGetNotFound with default headers values
func NewSMSCampaignBySMSCampaignIDGetNotFound() *SMSCampaignBySMSCampaignIDGetNotFound {
	return &SMSCampaignBySMSCampaignIDGetNotFound{}
}

/*SMSCampaignBySMSCampaignIDGetNotFound handles this case with default header values.

NOT_FOUND
*/
type SMSCampaignBySMSCampaignIDGetNotFound struct {
	Payload string
}

func (o *SMSCampaignBySMSCampaignIDGetNotFound) Error() string {
	return fmt.Sprintf("[GET /sms-campaigns/{sms_campaign_id}][%d] smsCampaignBySmsCampaignIdGetNotFound  %+v", 404, o.Payload)
}

func (o *SMSCampaignBySMSCampaignIDGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSCampaignBySMSCampaignIDGetMethodNotAllowed creates a SMSCampaignBySMSCampaignIDGetMethodNotAllowed with default headers values
func NewSMSCampaignBySMSCampaignIDGetMethodNotAllowed() *SMSCampaignBySMSCampaignIDGetMethodNotAllowed {
	return &SMSCampaignBySMSCampaignIDGetMethodNotAllowed{}
}

/*SMSCampaignBySMSCampaignIDGetMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type SMSCampaignBySMSCampaignIDGetMethodNotAllowed struct {
	Payload string
}

func (o *SMSCampaignBySMSCampaignIDGetMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /sms-campaigns/{sms_campaign_id}][%d] smsCampaignBySmsCampaignIdGetMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *SMSCampaignBySMSCampaignIDGetMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSCampaignBySMSCampaignIDGetTooManyRequests creates a SMSCampaignBySMSCampaignIDGetTooManyRequests with default headers values
func NewSMSCampaignBySMSCampaignIDGetTooManyRequests() *SMSCampaignBySMSCampaignIDGetTooManyRequests {
	return &SMSCampaignBySMSCampaignIDGetTooManyRequests{}
}

/*SMSCampaignBySMSCampaignIDGetTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type SMSCampaignBySMSCampaignIDGetTooManyRequests struct {
	Payload string
}

func (o *SMSCampaignBySMSCampaignIDGetTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /sms-campaigns/{sms_campaign_id}][%d] smsCampaignBySmsCampaignIdGetTooManyRequests  %+v", 429, o.Payload)
}

func (o *SMSCampaignBySMSCampaignIDGetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSMSCampaignBySMSCampaignIDGetDefault creates a SMSCampaignBySMSCampaignIDGetDefault with default headers values
func NewSMSCampaignBySMSCampaignIDGetDefault(code int) *SMSCampaignBySMSCampaignIDGetDefault {
	return &SMSCampaignBySMSCampaignIDGetDefault{
		_statusCode: code,
	}
}

/*SMSCampaignBySMSCampaignIDGetDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type SMSCampaignBySMSCampaignIDGetDefault struct {
	_statusCode int

	Payload string
}

// Code gets the status code for the Sms campaign by Sms campaign Id get default response
func (o *SMSCampaignBySMSCampaignIDGetDefault) Code() int {
	return o._statusCode
}

func (o *SMSCampaignBySMSCampaignIDGetDefault) Error() string {
	return fmt.Sprintf("[GET /sms-campaigns/{sms_campaign_id}][%d] SmsCampaignBySmsCampaignIdGet default  %+v", o._statusCode, o.Payload)
}

func (o *SMSCampaignBySMSCampaignIDGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}