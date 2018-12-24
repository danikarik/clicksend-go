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

// MMSCampaignsPricePostReader is a Reader for the MMSCampaignsPricePost structure.
type MMSCampaignsPricePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MMSCampaignsPricePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewMMSCampaignsPricePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewMMSCampaignsPricePostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewMMSCampaignsPricePostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewMMSCampaignsPricePostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewMMSCampaignsPricePostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewMMSCampaignsPricePostMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewMMSCampaignsPricePostTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewMMSCampaignsPricePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMMSCampaignsPricePostOK creates a MMSCampaignsPricePostOK with default headers values
func NewMMSCampaignsPricePostOK() *MMSCampaignsPricePostOK {
	return &MMSCampaignsPricePostOK{}
}

/*MMSCampaignsPricePostOK handles this case with default header values.

SUCCESS
*/
type MMSCampaignsPricePostOK struct {
	Payload interface{}
}

func (o *MMSCampaignsPricePostOK) Error() string {
	return fmt.Sprintf("[POST /mms-campaigns/price][%d] mmsCampaignsPricePostOK  %+v", 200, o.Payload)
}

func (o *MMSCampaignsPricePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMMSCampaignsPricePostBadRequest creates a MMSCampaignsPricePostBadRequest with default headers values
func NewMMSCampaignsPricePostBadRequest() *MMSCampaignsPricePostBadRequest {
	return &MMSCampaignsPricePostBadRequest{}
}

/*MMSCampaignsPricePostBadRequest handles this case with default header values.

BAD_REQUEST
*/
type MMSCampaignsPricePostBadRequest struct {
	Payload interface{}
}

func (o *MMSCampaignsPricePostBadRequest) Error() string {
	return fmt.Sprintf("[POST /mms-campaigns/price][%d] mmsCampaignsPricePostBadRequest  %+v", 400, o.Payload)
}

func (o *MMSCampaignsPricePostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMMSCampaignsPricePostUnauthorized creates a MMSCampaignsPricePostUnauthorized with default headers values
func NewMMSCampaignsPricePostUnauthorized() *MMSCampaignsPricePostUnauthorized {
	return &MMSCampaignsPricePostUnauthorized{}
}

/*MMSCampaignsPricePostUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type MMSCampaignsPricePostUnauthorized struct {
	Payload interface{}
}

func (o *MMSCampaignsPricePostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /mms-campaigns/price][%d] mmsCampaignsPricePostUnauthorized  %+v", 401, o.Payload)
}

func (o *MMSCampaignsPricePostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMMSCampaignsPricePostForbidden creates a MMSCampaignsPricePostForbidden with default headers values
func NewMMSCampaignsPricePostForbidden() *MMSCampaignsPricePostForbidden {
	return &MMSCampaignsPricePostForbidden{}
}

/*MMSCampaignsPricePostForbidden handles this case with default header values.

FORBIDDEN
*/
type MMSCampaignsPricePostForbidden struct {
	Payload interface{}
}

func (o *MMSCampaignsPricePostForbidden) Error() string {
	return fmt.Sprintf("[POST /mms-campaigns/price][%d] mmsCampaignsPricePostForbidden  %+v", 403, o.Payload)
}

func (o *MMSCampaignsPricePostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMMSCampaignsPricePostNotFound creates a MMSCampaignsPricePostNotFound with default headers values
func NewMMSCampaignsPricePostNotFound() *MMSCampaignsPricePostNotFound {
	return &MMSCampaignsPricePostNotFound{}
}

/*MMSCampaignsPricePostNotFound handles this case with default header values.

NOT_FOUND
*/
type MMSCampaignsPricePostNotFound struct {
	Payload interface{}
}

func (o *MMSCampaignsPricePostNotFound) Error() string {
	return fmt.Sprintf("[POST /mms-campaigns/price][%d] mmsCampaignsPricePostNotFound  %+v", 404, o.Payload)
}

func (o *MMSCampaignsPricePostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMMSCampaignsPricePostMethodNotAllowed creates a MMSCampaignsPricePostMethodNotAllowed with default headers values
func NewMMSCampaignsPricePostMethodNotAllowed() *MMSCampaignsPricePostMethodNotAllowed {
	return &MMSCampaignsPricePostMethodNotAllowed{}
}

/*MMSCampaignsPricePostMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type MMSCampaignsPricePostMethodNotAllowed struct {
	Payload interface{}
}

func (o *MMSCampaignsPricePostMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /mms-campaigns/price][%d] mmsCampaignsPricePostMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *MMSCampaignsPricePostMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMMSCampaignsPricePostTooManyRequests creates a MMSCampaignsPricePostTooManyRequests with default headers values
func NewMMSCampaignsPricePostTooManyRequests() *MMSCampaignsPricePostTooManyRequests {
	return &MMSCampaignsPricePostTooManyRequests{}
}

/*MMSCampaignsPricePostTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type MMSCampaignsPricePostTooManyRequests struct {
	Payload interface{}
}

func (o *MMSCampaignsPricePostTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /mms-campaigns/price][%d] mmsCampaignsPricePostTooManyRequests  %+v", 429, o.Payload)
}

func (o *MMSCampaignsPricePostTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMMSCampaignsPricePostDefault creates a MMSCampaignsPricePostDefault with default headers values
func NewMMSCampaignsPricePostDefault(code int) *MMSCampaignsPricePostDefault {
	return &MMSCampaignsPricePostDefault{
		_statusCode: code,
	}
}

/*MMSCampaignsPricePostDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type MMSCampaignsPricePostDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the Mms campaigns price post default response
func (o *MMSCampaignsPricePostDefault) Code() int {
	return o._statusCode
}

func (o *MMSCampaignsPricePostDefault) Error() string {
	return fmt.Sprintf("[POST /mms-campaigns/price][%d] MmsCampaignsPricePost default  %+v", o._statusCode, o.Payload)
}

func (o *MMSCampaignsPricePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
