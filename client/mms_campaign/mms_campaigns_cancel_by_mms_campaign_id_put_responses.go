// Code generated by go-swagger; DO NOT EDIT.

package mms_campaign

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// MMSCampaignsCancelByMMSCampaignIDPutReader is a Reader for the MMSCampaignsCancelByMMSCampaignIDPut structure.
type MMSCampaignsCancelByMMSCampaignIDPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MMSCampaignsCancelByMMSCampaignIDPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewMMSCampaignsCancelByMMSCampaignIDPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewMMSCampaignsCancelByMMSCampaignIDPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewMMSCampaignsCancelByMMSCampaignIDPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewMMSCampaignsCancelByMMSCampaignIDPutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewMMSCampaignsCancelByMMSCampaignIDPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewMMSCampaignsCancelByMMSCampaignIDPutMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewMMSCampaignsCancelByMMSCampaignIDPutTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewMMSCampaignsCancelByMMSCampaignIDPutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMMSCampaignsCancelByMMSCampaignIDPutOK creates a MMSCampaignsCancelByMMSCampaignIDPutOK with default headers values
func NewMMSCampaignsCancelByMMSCampaignIDPutOK() *MMSCampaignsCancelByMMSCampaignIDPutOK {
	return &MMSCampaignsCancelByMMSCampaignIDPutOK{}
}

/*MMSCampaignsCancelByMMSCampaignIDPutOK handles this case with default header values.

SUCCESS
*/
type MMSCampaignsCancelByMMSCampaignIDPutOK struct {
	Payload interface{}
}

func (o *MMSCampaignsCancelByMMSCampaignIDPutOK) Error() string {
	return fmt.Sprintf("[PUT /mms-campaigns/{mms_campaign_id}/cancel][%d] mmsCampaignsCancelByMmsCampaignIdPutOK  %+v", 200, o.Payload)
}

func (o *MMSCampaignsCancelByMMSCampaignIDPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMMSCampaignsCancelByMMSCampaignIDPutBadRequest creates a MMSCampaignsCancelByMMSCampaignIDPutBadRequest with default headers values
func NewMMSCampaignsCancelByMMSCampaignIDPutBadRequest() *MMSCampaignsCancelByMMSCampaignIDPutBadRequest {
	return &MMSCampaignsCancelByMMSCampaignIDPutBadRequest{}
}

/*MMSCampaignsCancelByMMSCampaignIDPutBadRequest handles this case with default header values.

BAD_REQUEST
*/
type MMSCampaignsCancelByMMSCampaignIDPutBadRequest struct {
	Payload interface{}
}

func (o *MMSCampaignsCancelByMMSCampaignIDPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /mms-campaigns/{mms_campaign_id}/cancel][%d] mmsCampaignsCancelByMmsCampaignIdPutBadRequest  %+v", 400, o.Payload)
}

func (o *MMSCampaignsCancelByMMSCampaignIDPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMMSCampaignsCancelByMMSCampaignIDPutUnauthorized creates a MMSCampaignsCancelByMMSCampaignIDPutUnauthorized with default headers values
func NewMMSCampaignsCancelByMMSCampaignIDPutUnauthorized() *MMSCampaignsCancelByMMSCampaignIDPutUnauthorized {
	return &MMSCampaignsCancelByMMSCampaignIDPutUnauthorized{}
}

/*MMSCampaignsCancelByMMSCampaignIDPutUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type MMSCampaignsCancelByMMSCampaignIDPutUnauthorized struct {
	Payload interface{}
}

func (o *MMSCampaignsCancelByMMSCampaignIDPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /mms-campaigns/{mms_campaign_id}/cancel][%d] mmsCampaignsCancelByMmsCampaignIdPutUnauthorized  %+v", 401, o.Payload)
}

func (o *MMSCampaignsCancelByMMSCampaignIDPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMMSCampaignsCancelByMMSCampaignIDPutForbidden creates a MMSCampaignsCancelByMMSCampaignIDPutForbidden with default headers values
func NewMMSCampaignsCancelByMMSCampaignIDPutForbidden() *MMSCampaignsCancelByMMSCampaignIDPutForbidden {
	return &MMSCampaignsCancelByMMSCampaignIDPutForbidden{}
}

/*MMSCampaignsCancelByMMSCampaignIDPutForbidden handles this case with default header values.

FORBIDDEN
*/
type MMSCampaignsCancelByMMSCampaignIDPutForbidden struct {
	Payload interface{}
}

func (o *MMSCampaignsCancelByMMSCampaignIDPutForbidden) Error() string {
	return fmt.Sprintf("[PUT /mms-campaigns/{mms_campaign_id}/cancel][%d] mmsCampaignsCancelByMmsCampaignIdPutForbidden  %+v", 403, o.Payload)
}

func (o *MMSCampaignsCancelByMMSCampaignIDPutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMMSCampaignsCancelByMMSCampaignIDPutNotFound creates a MMSCampaignsCancelByMMSCampaignIDPutNotFound with default headers values
func NewMMSCampaignsCancelByMMSCampaignIDPutNotFound() *MMSCampaignsCancelByMMSCampaignIDPutNotFound {
	return &MMSCampaignsCancelByMMSCampaignIDPutNotFound{}
}

/*MMSCampaignsCancelByMMSCampaignIDPutNotFound handles this case with default header values.

NOT_FOUND
*/
type MMSCampaignsCancelByMMSCampaignIDPutNotFound struct {
	Payload interface{}
}

func (o *MMSCampaignsCancelByMMSCampaignIDPutNotFound) Error() string {
	return fmt.Sprintf("[PUT /mms-campaigns/{mms_campaign_id}/cancel][%d] mmsCampaignsCancelByMmsCampaignIdPutNotFound  %+v", 404, o.Payload)
}

func (o *MMSCampaignsCancelByMMSCampaignIDPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMMSCampaignsCancelByMMSCampaignIDPutMethodNotAllowed creates a MMSCampaignsCancelByMMSCampaignIDPutMethodNotAllowed with default headers values
func NewMMSCampaignsCancelByMMSCampaignIDPutMethodNotAllowed() *MMSCampaignsCancelByMMSCampaignIDPutMethodNotAllowed {
	return &MMSCampaignsCancelByMMSCampaignIDPutMethodNotAllowed{}
}

/*MMSCampaignsCancelByMMSCampaignIDPutMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type MMSCampaignsCancelByMMSCampaignIDPutMethodNotAllowed struct {
	Payload interface{}
}

func (o *MMSCampaignsCancelByMMSCampaignIDPutMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /mms-campaigns/{mms_campaign_id}/cancel][%d] mmsCampaignsCancelByMmsCampaignIdPutMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *MMSCampaignsCancelByMMSCampaignIDPutMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMMSCampaignsCancelByMMSCampaignIDPutTooManyRequests creates a MMSCampaignsCancelByMMSCampaignIDPutTooManyRequests with default headers values
func NewMMSCampaignsCancelByMMSCampaignIDPutTooManyRequests() *MMSCampaignsCancelByMMSCampaignIDPutTooManyRequests {
	return &MMSCampaignsCancelByMMSCampaignIDPutTooManyRequests{}
}

/*MMSCampaignsCancelByMMSCampaignIDPutTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type MMSCampaignsCancelByMMSCampaignIDPutTooManyRequests struct {
	Payload interface{}
}

func (o *MMSCampaignsCancelByMMSCampaignIDPutTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /mms-campaigns/{mms_campaign_id}/cancel][%d] mmsCampaignsCancelByMmsCampaignIdPutTooManyRequests  %+v", 429, o.Payload)
}

func (o *MMSCampaignsCancelByMMSCampaignIDPutTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMMSCampaignsCancelByMMSCampaignIDPutDefault creates a MMSCampaignsCancelByMMSCampaignIDPutDefault with default headers values
func NewMMSCampaignsCancelByMMSCampaignIDPutDefault(code int) *MMSCampaignsCancelByMMSCampaignIDPutDefault {
	return &MMSCampaignsCancelByMMSCampaignIDPutDefault{
		_statusCode: code,
	}
}

/*MMSCampaignsCancelByMMSCampaignIDPutDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type MMSCampaignsCancelByMMSCampaignIDPutDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the Mms campaigns cancel by Mms campaign Id put default response
func (o *MMSCampaignsCancelByMMSCampaignIDPutDefault) Code() int {
	return o._statusCode
}

func (o *MMSCampaignsCancelByMMSCampaignIDPutDefault) Error() string {
	return fmt.Sprintf("[PUT /mms-campaigns/{mms_campaign_id}/cancel][%d] MmsCampaignsCancelByMmsCampaignIdPut default  %+v", o._statusCode, o.Payload)
}

func (o *MMSCampaignsCancelByMMSCampaignIDPutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
