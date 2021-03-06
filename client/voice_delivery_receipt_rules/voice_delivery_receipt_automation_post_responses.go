// Code generated by go-swagger; DO NOT EDIT.

package voice_delivery_receipt_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// VoiceDeliveryReceiptAutomationPostReader is a Reader for the VoiceDeliveryReceiptAutomationPost structure.
type VoiceDeliveryReceiptAutomationPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VoiceDeliveryReceiptAutomationPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewVoiceDeliveryReceiptAutomationPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewVoiceDeliveryReceiptAutomationPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewVoiceDeliveryReceiptAutomationPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewVoiceDeliveryReceiptAutomationPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewVoiceDeliveryReceiptAutomationPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewVoiceDeliveryReceiptAutomationPostMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewVoiceDeliveryReceiptAutomationPostTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewVoiceDeliveryReceiptAutomationPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVoiceDeliveryReceiptAutomationPostOK creates a VoiceDeliveryReceiptAutomationPostOK with default headers values
func NewVoiceDeliveryReceiptAutomationPostOK() *VoiceDeliveryReceiptAutomationPostOK {
	return &VoiceDeliveryReceiptAutomationPostOK{}
}

/*VoiceDeliveryReceiptAutomationPostOK handles this case with default header values.

SUCCESS
*/
type VoiceDeliveryReceiptAutomationPostOK struct {
	Payload interface{}
}

func (o *VoiceDeliveryReceiptAutomationPostOK) Error() string {
	return fmt.Sprintf("[POST /automations/voice/receipts][%d] voiceDeliveryReceiptAutomationPostOK  %+v", 200, o.Payload)
}

func (o *VoiceDeliveryReceiptAutomationPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceDeliveryReceiptAutomationPostBadRequest creates a VoiceDeliveryReceiptAutomationPostBadRequest with default headers values
func NewVoiceDeliveryReceiptAutomationPostBadRequest() *VoiceDeliveryReceiptAutomationPostBadRequest {
	return &VoiceDeliveryReceiptAutomationPostBadRequest{}
}

/*VoiceDeliveryReceiptAutomationPostBadRequest handles this case with default header values.

BAD_REQUEST
*/
type VoiceDeliveryReceiptAutomationPostBadRequest struct {
	Payload interface{}
}

func (o *VoiceDeliveryReceiptAutomationPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /automations/voice/receipts][%d] voiceDeliveryReceiptAutomationPostBadRequest  %+v", 400, o.Payload)
}

func (o *VoiceDeliveryReceiptAutomationPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceDeliveryReceiptAutomationPostUnauthorized creates a VoiceDeliveryReceiptAutomationPostUnauthorized with default headers values
func NewVoiceDeliveryReceiptAutomationPostUnauthorized() *VoiceDeliveryReceiptAutomationPostUnauthorized {
	return &VoiceDeliveryReceiptAutomationPostUnauthorized{}
}

/*VoiceDeliveryReceiptAutomationPostUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type VoiceDeliveryReceiptAutomationPostUnauthorized struct {
	Payload interface{}
}

func (o *VoiceDeliveryReceiptAutomationPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /automations/voice/receipts][%d] voiceDeliveryReceiptAutomationPostUnauthorized  %+v", 401, o.Payload)
}

func (o *VoiceDeliveryReceiptAutomationPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceDeliveryReceiptAutomationPostForbidden creates a VoiceDeliveryReceiptAutomationPostForbidden with default headers values
func NewVoiceDeliveryReceiptAutomationPostForbidden() *VoiceDeliveryReceiptAutomationPostForbidden {
	return &VoiceDeliveryReceiptAutomationPostForbidden{}
}

/*VoiceDeliveryReceiptAutomationPostForbidden handles this case with default header values.

FORBIDDEN
*/
type VoiceDeliveryReceiptAutomationPostForbidden struct {
	Payload interface{}
}

func (o *VoiceDeliveryReceiptAutomationPostForbidden) Error() string {
	return fmt.Sprintf("[POST /automations/voice/receipts][%d] voiceDeliveryReceiptAutomationPostForbidden  %+v", 403, o.Payload)
}

func (o *VoiceDeliveryReceiptAutomationPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceDeliveryReceiptAutomationPostNotFound creates a VoiceDeliveryReceiptAutomationPostNotFound with default headers values
func NewVoiceDeliveryReceiptAutomationPostNotFound() *VoiceDeliveryReceiptAutomationPostNotFound {
	return &VoiceDeliveryReceiptAutomationPostNotFound{}
}

/*VoiceDeliveryReceiptAutomationPostNotFound handles this case with default header values.

NOT_FOUND
*/
type VoiceDeliveryReceiptAutomationPostNotFound struct {
	Payload interface{}
}

func (o *VoiceDeliveryReceiptAutomationPostNotFound) Error() string {
	return fmt.Sprintf("[POST /automations/voice/receipts][%d] voiceDeliveryReceiptAutomationPostNotFound  %+v", 404, o.Payload)
}

func (o *VoiceDeliveryReceiptAutomationPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceDeliveryReceiptAutomationPostMethodNotAllowed creates a VoiceDeliveryReceiptAutomationPostMethodNotAllowed with default headers values
func NewVoiceDeliveryReceiptAutomationPostMethodNotAllowed() *VoiceDeliveryReceiptAutomationPostMethodNotAllowed {
	return &VoiceDeliveryReceiptAutomationPostMethodNotAllowed{}
}

/*VoiceDeliveryReceiptAutomationPostMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type VoiceDeliveryReceiptAutomationPostMethodNotAllowed struct {
	Payload interface{}
}

func (o *VoiceDeliveryReceiptAutomationPostMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /automations/voice/receipts][%d] voiceDeliveryReceiptAutomationPostMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *VoiceDeliveryReceiptAutomationPostMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceDeliveryReceiptAutomationPostTooManyRequests creates a VoiceDeliveryReceiptAutomationPostTooManyRequests with default headers values
func NewVoiceDeliveryReceiptAutomationPostTooManyRequests() *VoiceDeliveryReceiptAutomationPostTooManyRequests {
	return &VoiceDeliveryReceiptAutomationPostTooManyRequests{}
}

/*VoiceDeliveryReceiptAutomationPostTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type VoiceDeliveryReceiptAutomationPostTooManyRequests struct {
	Payload interface{}
}

func (o *VoiceDeliveryReceiptAutomationPostTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /automations/voice/receipts][%d] voiceDeliveryReceiptAutomationPostTooManyRequests  %+v", 429, o.Payload)
}

func (o *VoiceDeliveryReceiptAutomationPostTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceDeliveryReceiptAutomationPostDefault creates a VoiceDeliveryReceiptAutomationPostDefault with default headers values
func NewVoiceDeliveryReceiptAutomationPostDefault(code int) *VoiceDeliveryReceiptAutomationPostDefault {
	return &VoiceDeliveryReceiptAutomationPostDefault{
		_statusCode: code,
	}
}

/*VoiceDeliveryReceiptAutomationPostDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type VoiceDeliveryReceiptAutomationPostDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the voice delivery receipt automation post default response
func (o *VoiceDeliveryReceiptAutomationPostDefault) Code() int {
	return o._statusCode
}

func (o *VoiceDeliveryReceiptAutomationPostDefault) Error() string {
	return fmt.Sprintf("[POST /automations/voice/receipts][%d] VoiceDeliveryReceiptAutomationPost default  %+v", o._statusCode, o.Payload)
}

func (o *VoiceDeliveryReceiptAutomationPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
